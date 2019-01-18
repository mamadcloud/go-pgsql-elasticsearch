package main

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"time"
	"io/ioutil"
	"bytes"

	services "DeltaUpdate/pkg/services"
)

func main() {
	for {
		runUpdate()
		time.Sleep(time.Millisecond * time.Duration(5000))
	}
}

func runUpdate() {
	deltaUpdates, _ := services.FindDeltaUpdates()
	for _, deltaUpdate := range deltaUpdates {
		uri := "http://elasticsearch:9200/ads/_doc/" + strconv.Itoa(deltaUpdate.ItemID)
		client := &http.Client{}
		client.Timeout = time.Second * 15

		if deltaUpdate.Action != "delete" {
			ads, _ := services.FindAd(deltaUpdate.ItemID)
			bodyReq, err := json.Marshal(ads)
			
			body := bytes.NewBuffer(bodyReq)
			req, err := http.NewRequest(http.MethodPut, uri, body)
			if err != nil {
				fmt.Println("http.NewRequest() failed with '%s'\n", err)
				return
			}
			
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("client.Do() failed with '%s'\n", err)
			}
			
			defer resp.Body.Close()
			bodyRes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("ioutil.ReadAll() failed with '%s'\n", bodyRes)
				return
			}
			fmt.Println("ioutil.ReadAll() '%s'\n", string(bodyRes))
		} else {
			req, err := http.NewRequest(http.MethodDelete, uri, nil)
			if err != nil {
				fmt.Println("http.NewRequest() failed with '%s'\n", err)
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("client.Do() failed with '%s'\n", err)
				return
			}
			
			defer resp.Body.Close()
			bodyRes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("ioutil.ReadAll() failed with '%s'\n", bodyRes)
				return
			}
			fmt.Println("ioutil.ReadAll() '%s'\n", string(bodyRes))
		}

		services.DeleteDeltaUpdate(deltaUpdate.ID);
	}
}