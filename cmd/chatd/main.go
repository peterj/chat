package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ardanlabs/kit/cfg"
	"github.com/ardanlabs/kit/tcp"
)

const (
	configKey = "CHAT"
)

func init() {
	os.Setenv("CHAT_HOST", ":6000")

	log.SetOutput(os.Stdout)
	// Logging is everything - if you don't do it now (at the beginning),
	// you won't do it later
	// sets the logging flags e.g.: 2017/09/25 08:25:29.471778 main.go:29: Configuration
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds)
}

func main() {
	if err := cfg.Init(cfg.EnvProvider{Namespace: configKey}); err != nil {
		fmt.Println("Error initializing config system", err)
		os.Exit(1)
	}

	log.Println("Configuration\n", cfg.Log())

	host := cfg.MustString("HOST")

	cfg := tcp.Config{
		NetType: "tcp4",
		Addr:    host,

		ConnHandler: connHandler{},
		ReqHandler:  reqHandler{},
		RespHandler: respHandler{},

		OptEvent: tcp.OptEvent{
			Event: Event,
		},
	}

	// Create a new TCP value.
	t, err := tcp.New("Sample", cfg)
	if err != nil {
		log.Printf("main : %s", err)
		return
	}

	// Start accepting client data.
	if err := t.Start(); err != nil {
		log.Printf("main : %s", err)
		return
	}

	// Defer the stop on shutdown.
	defer t.Stop()

	log.Printf("main: Waiting for data on: %s", t.Addr())

	// Listen for an interrupt signal from the OS.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	// Use telnet to test the server.
	// telnet localhost 6000

}

var evtTypes = []string{
	"unknown",
	"Accept",
	"Join",
	"Read",
	"Remove",
	"Drop",
	"Groom",
}

var typeTypes = []string{
	"unknown",
	"Error",
	"Info",
	"Trigger",
}

func Event(evt, typ int, ipAddress string, format string, a ...interface{}) {
	log.Printf("***> EVENT : IP [ %s ] : EVT[%s] TYP[%s]: %s", ipAddress, evtTypes[evt], typeTypes[typ], fmt.Sprintf(format, a...))
}
