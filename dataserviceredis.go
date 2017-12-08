package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type DataServiceRedis struct{}

func getRedisClient(credentials *Credentials) (*redis.Client, error) {
	log.Printf(fmt.Sprintf("Create redis client for %v:%v\n", (*credentials)["host"].(string), (*credentials)["port"].(float64)))
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", (*credentials)["host"].(string), (*credentials)["port"].(float64)),
		Password: (*credentials)["password"].(string),
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()
	log.Printf("pong: %v ; err = %v\n", pong, err)

	return client, err
}

func (d DataServiceRedis) GetStatus(credentials *Credentials) (int, error) {
	client, err := getRedisClient(credentials)
	if err != nil {
		return 2, err
	}

	testKey := "testkey"
	t := time.Now()
	testValue := t.Format("20060102150405")
	err = client.Set(testKey, testValue, 0).Err()
	if err != nil {
		return 2, err
	}

	value, err := client.Get(testKey).Result()
	if err != nil {
		return 2, err
	}

	found := value == testValue

	client.Del(testKey)

	if found {
		return 0, nil
	}

	return 1, nil
}
func (d DataServiceRedis) Insert(credentials *Credentials, value string) error {
	client, err := getRedisClient(credentials)
	if err != nil {
		return err
	}

	testKey := "testkey"
	err = client.Set(testKey, value, 0).Err()
	if err != nil {
		return err
	}

	return fmt.Errorf("not implemented")
}

func (d DataServiceRedis) Exists(credentials *Credentials, value string) (bool, error) {
	client, err := getRedisClient(credentials)
	if err != nil {
		return false, err
	}

	testKey := "testkey"
	returnValue, err := client.Get(testKey).Result()

	if err != nil {
		return false, err
	}

	return returnValue == value, fmt.Errorf("not implemented")
}

func (d DataServiceRedis) Delete(credentials *Credentials, value string) error {
	client, err := getRedisClient(credentials)
	if err != nil {
		return err
	}

	testKey := "testkey"
	client.Del(testKey)

	return nil
}

func CreateRedis() IDataService {
	return DataServiceRedis{}
}
