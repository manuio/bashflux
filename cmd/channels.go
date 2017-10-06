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
	"strconv"
	"strings"
	"encoding/json"
	"io/ioutil"

	"github.com/mainflux/mainflux-core/models"
)

var endPointC = "/channels"

// CreateChannel - creates new channel and generates UUID
func CreateChannel(msg string, token string) string {
	url := UrlHTTP + endPointC
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

// GetChannels - gets all channels
func GetChannels(limit int, token string) string {
	url := UrlHTTP + "/channels?climit=" + strconv.Itoa(limit)
	req, err := http.NewRequest("GET", url,  nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHttpResp(resp, err)

	return s
}

// GetChannel - gets channel by ID
func GetChannel(id string, token string) string {
	url := UrlHTTP + "/channels/" + id
	req, err := http.NewRequest("GET", url,  nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	s := PrettyHttpResp(resp, err)

	return s
}

// UpdateChannel - publishes SenML message on the channel
func UpdateChannel(id string, msg string, token string) string {
	url := UrlHTTP + "/channels/" + id
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

// DeleteChannel - removes channel
func DeleteChannel(id string, token string) string {
	url := UrlHTTP + "/channels/" + id
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

// DeleteAllChannels - removes all channels
func DeleteAllChannels(token string) string {
	url := UrlHTTP + "/channels"
	req, err := http.NewRequest("GET", url,  nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/senml+json")

	resp, err := netClient.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	var channels []models.Channel
	json.Unmarshal([]byte(body), &channels)
	s := ""
	for i := 0; i < len(channels); i++ {
		s = s + DeleteChannel(channels[i].ID, token) + "\n\n"
	}

	return s
}
