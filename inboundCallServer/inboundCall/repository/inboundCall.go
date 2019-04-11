package repository

import (
	"strings"
	. "github.com/0x19/goesl"
	"github.com/Sirupsen/logrus"
)

func NewESLConnectionHandle(s *OutboundServer,logger *logrus.Logger) {

	for {
		conn := <-s.Conns
		logger.Info("New incomming connection: %v", conn)

		if err := conn.Connect(); err != nil {
			logger.WithError(err).Fatal("Got error while accepting connection: %s", err)
			break
		}

		go func() {
			for {
				msg, err := conn.ReadMessage()
				if err != nil {

					// If it contains EOF, we really dont care...
					if !strings.Contains(err.Error(), "EOF") {
						logger.WithError(err).Fatal("Error while reading Freeswitch message: %s", err)
					}
					break
				}
				logger.Info("FreeSWITCH Message %s", msg)
			}
		}()
	}

}
