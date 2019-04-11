package main

import (
	. "github.com/0x19/goesl"
	"strconv"
	"inboundCallServer/config"
	"inboundCallServer/logger"
	inboundHandler "inboundCallServer/inboundCall/repository"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			Error("Recovered in f", r)
		}
	}()


	//Setting up the config
	config := config.GetConfig()

	//Setting up the Logger
	log := logger.NewLogger(config.Log.LogFile, config.Log.LogLevel)

	//Setting up the listen address
	outboundSocketPort := strconv.Itoa(config.EslConfig.Port)
	outboundSocketAddr := config.EslConfig.Host + ":" + outboundSocketPort

	log.Info("outbound server address is : %s", outboundSocketAddr)

	// TODO address should be coming from config file only
	if s, err := NewOutboundServer(outboundSocketAddr); err != nil {
		log.WithError(err).Fatal("Freeswitch outbound server is not able to start: %s", err)
	} else {
		go inboundHandler.NewESLConnectionHandle(s,log)
		s.Start()
	}
}

