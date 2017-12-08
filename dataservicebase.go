package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
)

func GetCredentials(params *map[strin]string) (*map[string]interface{}, error) {
	log.Println("SetCredentials!")
	id := params["id"]
	credentialsParam := params["credentials"]

	if len(credentialsParam) > 0 {
		credentials := make(map[string]interface{})
		err := json.Unmarshal([]byte(credentialsParam), credentials)
		if err != nil {
			return nil, err
		}
		return &credentials, nil
	}

	log.Println("No default credentials given.")
	// In any case of an exception recover to not stop the app
	defer fmt.Errorf("Could not parse VCAP_SERVICES.")

	// read the env variables
	appEnv, err := cfenv.Current()
	if err != nil {
		return nil, err
	}

	// convert string to i
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	var service cfenv.Service

	// get first service and than the index
	for _, v := range appEnv.Services {
		service = v[i]
		break
	}

	credentials := make(map[string]interface{})
	for k, v := range service.Credentials {
		credentials[k] = v
	}

	log.Println("print credentials")
	for k, v := range credentials {
		log.Printf("%s----%s", k, v)
	}

	return &credentials, nil
}
