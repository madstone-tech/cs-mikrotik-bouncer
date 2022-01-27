package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var addr string = os.Getenv("ADDR")
var username string = os.Getenv("USERNAME")
var password string = os.Getenv("PASSWORD")
var resId string

type Blacklist struct {
	Address string `json:"address"`
	List    string `json:"list"`
	Timeout string `json:"timeout"`
}

func addAddress(ip string, duration string, name string) {

	re := regexp.MustCompile("[a-z].")
	txt := duration
	split := re.Split(txt, -1)
	set := []string{}
	for i := range split {
		set = append(set, split[i])
	}

	rosDuration := fmt.Sprintf(set[0] + ":" + set[1] + ":" + set[2])
	fmt.Println(rosDuration)

	addAddr := fmt.Sprintf("https://%s/rest/ip/firewall/address-list/add", addr)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	data := []byte(fmt.Sprintf(`{"address":"%s","list":"%s","timeout":"%s"}`, ip, name, rosDuration))

	req, err := http.NewRequest(http.MethodPost, addAddr, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(username, password)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Adding Status: %d\n", res.StatusCode)
	fmt.Printf("Adding Body: %s\n", string(resBody))
}

func getAddress(ip string, name string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	getAddr := fmt.Sprintf("https://%s/rest/ip/firewall/address-list", addr)

	req, err := http.NewRequest(http.MethodGet, getAddr, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(username, password)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Getting Status: %d\n", res.StatusCode)
	fmt.Printf("Getting Body: %s\n", string(resBody))

	var resArray []map[string]interface{}

	_ = json.Unmarshal(resBody, &resArray)

	var noCidr string = ip[:strings.IndexByte(ip, '/')]

	for i := range resArray {
		fmt.Println("array fired")
		if resArray[i]["address"] == noCidr {
			fmt.Println("Found", resArray[i]["address"], resArray[i][".id"])
			resId = fmt.Sprintf("%v", resArray[i][".id"])
		}
	}
}

func delAddress(ip string, duration string, name string) {

	getAddress(ip, name)

	delAddr := fmt.Sprintf("https://%s/rest/ip/firewall/address-list/%s", addr, resId)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(http.MethodDelete, delAddr, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(username, password)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleting Status: %d\n", res.StatusCode)
	fmt.Printf("Deleting Body: %s\n", string(resBody))
}

func main() {

	flag.Parse()

	if flag.Arg(0) == "add" {
		addAddress(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	} else {
		delAddress(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	}
}
