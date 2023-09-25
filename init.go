package scyna

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/nats-io/nats.go"
	scyna_const "github.com/scyna/core/const"
	"github.com/scyna/core/internal/base"
	scyna_proto "github.com/scyna/core/proto"
	"google.golang.org/protobuf/proto"
)

type RemoteConfig struct {
	ManagerUrl string
	Name       string
	Secret     string
}

func RemoteInit(config RemoteConfig) {

	request := scyna_proto.CreateSessionRequest{
		Module: config.Name,
		Secret: config.Secret,
	}

	data, err := proto.Marshal(&request)
	if err != nil {
		log.Fatal("Bad authentication request")
	}

	req, err := http.NewRequest("POST", config.ManagerUrl+scyna_const.SESSION_CREATE_URL, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error in create http request:", err)
	}

	res, err := HttpClient().Do(req)
	if err != nil {
		log.Fatal("Error in send http request:", err)
	}

	if res.StatusCode != 200 {
		log.Fatal("Error in autheticate")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Can not read response body:", err)
	}

	var response scyna_proto.CreateSessionResponse
	if err := proto.Unmarshal(resBody, &response); err != nil {
		log.Fatal("Authenticate error")
	}

	Session = NewSession(response.SessionID)
	DirectInit(config.Name, response.Config)
}

func DirectInit(name string, c *scyna_proto.Configuration) {
	module = name
	var err error
	var nats_ []string
	for _, n := range strings.Split(c.NatsUrl, ",") {
		fmt.Printf("Nats configuration: nats://%s:4222\n", n)
		nats_ = append(nats_, fmt.Sprintf("nats://%s:4222", n))
	}

	if c.NatsUsername != "" && c.NatsPassword != "" {
		Nats, err = nats.Connect(strings.Join(nats_, ","), nats.UserInfo(c.NatsUsername, c.NatsPassword))
	} else {
		Nats, err = nats.Connect(strings.Join(nats_, ","))
	}

	if err != nil {
		log.Fatal("Can not connect to NATS:", nats_)
	}

	/*init jetstream*/
	JetStream, err = Nats.JetStream()
	if err != nil {
		panic("Init: " + err.Error())
	}

	hosts := strings.Split(c.DBHost, ",")
	DB = base.NewDB(hosts, c.DBUsername, c.DBPassword, c.DBLocation)

	Settings.init()

	/*registration*/
	RegisterSignal(scyna_const.SETTING_UPDATE_CHANNEL+module, updateSettingHandler, SIGNAL_SCOPE_SESSION)
	RegisterSignal(scyna_const.SETTING_REMOVE_CHANNEL+module, removeSettingHandler, SIGNAL_SCOPE_SESSION)
}
