package domain

import "log-management/utils"

type Microservice struct {
	Base

	Name        string `bun:",notnull"`
	Description string `bun:",notnull"`
}

func NewMicroservice(name string, description string) *Microservice {
	m := &Microservice{
		Name:        name,
		Description: description,
	}
	m.ID = utils.NewID("microservice")

	return m
}