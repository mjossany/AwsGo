package api

import (
	"lambda-func/database"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore string) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}
