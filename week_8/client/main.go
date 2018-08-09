package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	getUser()
	getAdmin()
}

func getUser() {
	url := "http://127.0.0.1:8888/user/get"

	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	fmt.Printf("user data: %v\n", result)

	defer resp.Body.Close()
}

func getAdmin() {
	url := "http://127.0.0.1:8888/admin/get"

	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	fmt.Printf("admin data: %v\n", result)

	defer resp.Body.Close()
}
