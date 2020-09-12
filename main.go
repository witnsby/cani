package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"strings"
	"time"
)

const (
	UrlForCheck = "http://ifconfig.io/"
)

type listData map[string]string

// func putDefaultParameters is a list with data which should be requested
func (l listData) putDefaultParameters() {
	myList := [2]string{"country_code", "ip"}

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

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get(UrlForCheck + s)
	if err != nil {
		log.Printf("%v", err)
		return ""
	}

	bites, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%v", err)
		return ""
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return strings.TrimRight(string(bites), "\r\n")
}

func (l listData) generateMyJson() {
	data, err := json.Marshal(l)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	createConfigFile(data)
}

// func getCurrentUserHomeDir get information about current user homeDir
func getCurrentUserHomeDir() string {
	user, err := user.Current()
	if err != nil {
		log.Printf("%v", err)
	}
	return user.HomeDir
}

// func createConfigFile creates json file
func createConfigFile(data []byte) {
	homeDir := getCurrentUserHomeDir()

	if err := ioutil.WriteFile(homeDir+"/"+"cani.json", data, 0644); err != nil {
		log.Printf("%v", err)
	}
}

func main() {
	listData := make(listData)

	for {
		listData.putDefaultParameters()
		listData.putCertainParameters()
		listData.generateMyJson()
		time.Sleep(2 * time.Second)
	}
}
