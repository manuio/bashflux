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
)

// Status - server health check
func Status() string {
	url := UrlHTTP + "/status"
	rsp, err := netClient.Get(url)
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	b, e := prettyJSON(body)
	if e != nil {
		return "ERROR JSON"
	}

	return string(b)
}