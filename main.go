package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const UrlForCheck = "http://ifconfig.io/"

type listData map[string]string

// func putDefaultParameters is a list with data which should be requested
func (l listData) putDefaultParameters() {
	myList := [2]string{"country_code","ip"}

	for _, elem := range myList {
		l[elem] = "nil"
	}
}

// func putCertainParameters creates map with actual information
func (l listData) putCertainParameters() {
	for key := range l {
		l[key] = requestData(key)
	}
}

// func requestData should return actual information like IP, Country
func requestData(s string) string {
	resp, err := http.Get(UrlForCheck + s)
	if err != nil {
		log.Fatal(err)
	}

	bites, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return string(bites)
}

// func setMyEnvVariables should create local environment variables in the OS
func (l listData) setMyEnvVariables() {
	for key, value := range l {
		err := os.Setenv(key,value)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// func getMyEnvVariables is a testing func
func (l listData) getMyEnvVariables() {
	for key := range l {
		fmt.Println(os.Getenv(key))
	}
}


func main() {

	listData := make(listData)

	for {
		listData.putDefaultParameters()
		listData.putCertainParameters()
		listData.setMyEnvVariables()
		time.Sleep(5 * time.Second)
		//listData.getMyEnvVariables()
	}
}