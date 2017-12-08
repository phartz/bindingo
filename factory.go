package main

import (
	"fmt"
	"log"
)

var RegisteredServices map[string]func() IDataService

func GetDataService(params map[string]string) (IDataService, error) {
	//var service IDataService = nil

	log.Printf("Get class for [%s]\n", params["dataService"])

	createFunc := RegisteredServices[params["dataService"]]
	if createFunc == nil {
		return nil, fmt.Errorf("No class found for [%s]!\n", params["dataService"])
	}
	return createFunc(), nil
	/*
		switch params["dataService"] {
		case "testservice":
			service = TestService{}
			break
		case "a9s_mongodb":
			service = DataServiceMongoDB{}
			break
		case "a9s_redis":
			service = DataServiceRedis{}
			break
		default:
			log.Printf("No class found for [%s]!\n", params["dataService"])
		}

		return service, nil*/
}
