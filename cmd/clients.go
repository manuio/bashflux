/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package cmd

import (
	"net/http"
	"strings"
	"encoding/json"
	"io/ioutil"
)

var endPoint = "/clients"

// CreateClient - creates new client and generates client UUID
func CreateClient(msg string, token string) string {
	url := UrlHTTP + endPoint
	req, err := http.NewRequest("POST", url,  strings.NewReader(msg))
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHttpResp(resp, err)

	return s
}

// GetClients - gets all clients
func GetClients(token string) string {
	url := UrlHTTP + endPoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHttpResp(resp, err)

	return s
}

// GetClient - gets client by ID
func GetClient(id string, token string) string {
	url := UrlHTTP + endPoint + "/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHttpResp(resp, err)

	return s
}

// UpdateClient - updates client by ID
func UpdateClient(id string, msg string, token string) string {
	url := UrlHTTP + endPoint + "/" + id
	req, err := http.NewRequest("PUT", url, strings.NewReader(msg))
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHttpResp(resp, err)

	return s
}

// DeleteClient - removes client
func DeleteClient(id string, token string) string {
	url := UrlHTTP + endPoint + "/" + id
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHttpResp(resp, err)

	return s
}

// DeleteAllClients - removes all clients
func DeleteAllClients(token string) string {
	url := UrlHTTP + endPoint
	resp, _ := netClient.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)

	var clients []struct{}
	json.Unmarshal([]byte(body), &clients)
	println(clients)
	s := ""
	for i := 0; i < len(clients); i++ {
		println(".")
		//s = s + DeleteClient(clients[nil, token) + "\n\n"
	}

	return s
}
