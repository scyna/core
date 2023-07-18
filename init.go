package scyna

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sigv4-auth-cassandra-gocql-driver-plugin/sigv4"
	"github.com/gocql/gocql"
	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2"
	scyna_const "github.com/scyna/core/const"
	scyna_proto "github.com/scyna/core/proto/generated"
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
		Connection, err = nats.Connect(strings.Join(nats_, ","), nats.UserInfo(c.NatsUsername, c.NatsPassword))
	} else {
		Connection, err = nats.Connect(strings.Join(nats_, ","))
	}

	if err != nil {
		log.Fatal("Can not connect to NATS:", nats_)
	}

	/*init jetstream*/
	JetStream, err = Connection.JetStream()
	if err != nil {
		panic("Init: " + err.Error())
	}

	/*init db*/
	hosts := strings.Split(c.DBHost, ",")
	if c.IsAWSKeyspaces {
		initKeyspaces(c.DBHost, c.DBUsername, c.DBPassword, c.DBLocation, c.DBPemFile)
	} else {
		initScylla(hosts, c.DBUsername, c.DBPassword, c.DBLocation)
	}

	Settings.init()

	/*registration*/
	RegisterSignal(scyna_const.SETTING_UPDATE_CHANNEL+module, updateSettingHandler, SIGNAL_SCOPE_SESSION)
	RegisterSignal(scyna_const.SETTING_REMOVE_CHANNEL+module, removeSettingHandler, SIGNAL_SCOPE_SESSION)
}

func initScylla(host []string, username string, password string, location string) {
	cluster := gocql.NewCluster(host...)
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: username, Password: password}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(location)
	cluster.ConnectTimeout = time.Second * 10
	cluster.DisableInitialHostLookup = true
	cluster.Consistency = gocql.Quorum

	//TODO: Config connect with TLS/SSL

	log.Printf("Connect to db: %s\n", host)

	var err error
	DB, err = gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		panic(fmt.Sprintf("Can not create session: Host = %s, Error = %s ", host, err.Error()))
	}
}

func initKeyspaces(host string, accessKey string, secretKey string, region string, pemFile string) {
	cluster := gocql.NewCluster(host)
	var auth sigv4.AwsAuthenticator = sigv4.NewAwsAuthenticator()
	auth.Region = region
	auth.AccessKeyId = accessKey
	auth.SecretAccessKey = secretKey

	cluster.Authenticator = auth

	cluster.SslOpts = &gocql.SslOptions{
		CaPath: pemFile,
	}
	cluster.Consistency = gocql.LocalQuorum
	cluster.DisableInitialHostLookup = true

	var err error
	DB, err = gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		panic(fmt.Sprintf("Can not create session: Host = %s, Error = %s ", host, err.Error()))
	}
}
