package pkg

import (
	"app/config"
)

type metaData struct {
	BACKEND_VERSION  string      `json:"backend_version" bson:"backend_version"`
	FRONTEND_VERSION string      `json:"frontend_version" bson:"frontend_version"`
	APP_VERSION      string      `json:"app_version" bson:"app_version"`
	MODE             string      `json:"mode" bson:"mode"`
	SUCCESS          bool        `json:"success" bson:"success"`
	RESULT           interface{} `json:"result" bson:"result"`
	ERRORS           interface{} `json:"errors" bson:"errors"`
}

func AddMetaData(data interface{}, status int) *metaData {
	metadata := metaData{
		BACKEND_VERSION:  config.AppConfig.BACKEND_VERSION,
		FRONTEND_VERSION: config.AppConfig.FRONTEND_VERSION,
		APP_VERSION:      config.AppConfig.APP_VERSION,
		MODE:             config.AppConfig.MODE,
	}

	if status >= 400 {
		metadata.SUCCESS = false
		metadata.ERRORS = data
	} else {
		metadata.SUCCESS = true
		metadata.RESULT = data
	}

	return &metadata
}
