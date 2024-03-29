package scyna

import (
	"crypto/tls"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

type SetupHandler func()

var setups []SetupHandler
var testMode bool = false

func RegisterSetup(starter SetupHandler) {
	setups = append(setups, starter)
}

func startAll() {
	for _, starter := range setups {
		starter()
	}
}

const REQUEST_TIMEOUT = 10

var Nats *nats.Conn
var JetStream nats.JetStreamContext
var Session *session
var DB *db
var ID generator
var Settings settings

var httpClient *http.Client
var module string

func Module() string {
	return module
}

func Release() {
	releaseLog()
	Session.release()
	Nats.Close()
	DB.Close()
}

func Start() {
	startDomainEvent()
	startEventStreams()
	startEndpoints()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func HttpClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 5,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	}
	return httpClient
}
