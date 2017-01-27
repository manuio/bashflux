/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"encoding/json"

	"github.com/mainflux/mainflux-core/models"
	"github.com/hokaccha/go-prettyjson"
)

// CreateDevice - creates new device and generates device UUID
func CreateDevice(msg string) string {
	var err error

	url := UrlHTTP + "/devices"
	rsp, err := netClient.Post(url, "application/json", nil)
	if err != nil {
		return err.Error()
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, err := prettyjson.Format([]byte(body))
	if err != nil {
		return err.Error()
	}

	return string(b)
}

// GetDevices - gets all devices
func GetDevices() string {
	url := UrlHTTP + "/devices"
	rsp, err := netClient.Get(url)
	if err != nil {
		return err.Error()
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, e := prettyjson.Format([]byte(body))
	if e != nil {
		return err.Error()
	}

	return string(b)
}

// GetDevice - gets device by ID
func GetDevice(id string) string {
	url := UrlHTTP + "/devices/" + id
	rsp, err := netClient.Get(url)
	if err != nil {
		return err.Error()
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, e := prettyjson.Format([]byte(body))
	if e != nil {
		return err.Error()
	}

	return string(b)
}

// UpdateDevice - updates device by ID
func UpdateDevice(id string, msg string) string {
	var err error

	url := UrlHTTP + "/devices/" + id
	sr := strings.NewReader(msg)
	req, err := http.NewRequest("PUT", url, sr)
	if err != nil {
		return err.Error()
	}

	req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(msg)))

	rsp, err := netClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, err := prettyjson.Format([]byte(body))
	if err != nil {
		return err.Error()
	}

	return string(b)
}

// DeleteDevice - removes device
func DeleteDevice(id string) string {
	var err error

	url := UrlHTTP + "/devices/" + id
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err.Error()
	}
	rsp, err := netClient.Do(req)
	if err != nil {
		return err.Error()
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, err := prettyjson.Format([]byte(body))
	if err != nil {
		return err.Error()
	}

	return string(b)
}

// DeleteAllDevices - removes all devices
func DeleteAllDevices() string {
	var err error

	url := UrlHTTP + "/devices"
	rsp, err := netClient.Get(url)
	if err != nil {
		return err.Error()
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	var devices []models.Device
	json.Unmarshal(body, &devices)
	s := ""
	for i := 0; i < len(devices); i++ {
		s = s + DeleteDevice(devices[i].ID)
	}

	return s
}

// CreateDevice - creates new device and generates device UUID
func PlugDevice(id string, channels string) string {
	var err error

	url := UrlHTTP + "/devices/" + id + "/plug"
	sr := strings.NewReader(channels)
	rsp, err := netClient.Post(url, "application/json", sr)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, err := prettyjson.Format([]byte(body))
	if err != nil {
		return err.Error()
	}

	return string(b)
}
