package domain

import "log-management/utils"

type Log struct {
	Base

	MicroserviceID string `bun:",notnull"`
	LogLevel       string `bun:",notnull"`
	Message        string `bun:",notnull"`

	Microservice *Microservice `bun:"rel:belongs-to,join:microservice_id=id"`
}

func NewLog(microserviceID string, logLevel string, message string) *Log {
	l := &Log{
		MicroserviceID: microserviceID,
		LogLevel:       logLevel,
		Message:        message,
	}
	l.ID = utils.NewID("log")

	return l
}
