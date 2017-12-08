package main

import "log"

func GetDataService(params map[string]string) (IDataService, error) {
	var service IDataService = nil

	log.Printf("Get class for [%s]\n", params["dataService"])

	switch params["dataService"] {
	case "testservice":
		service = TestService{}
		break
	case "a9s_mongodb":
		service = DataServiceMongoDB{}
		break
	default:
		log.Printf("No class found for [%s]!\n", params["dataService"])
	}

	return service, nil
}
