package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/joho/godotenv"
)

var addr string = os.Getenv("ADDR")
var username string = os.Getenv("USERNAME")
var password string = os.Getenv("PASSWORD")

type Blacklist struct {
	Address string `json:"address"`
	List    string `json:"list"`
	Timeout string `json:"timeout"`
}

func addAddress(ip string, duration string, name string) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addAddr := fmt.Sprintf("https://%s/rest/ip/firewall/address-list/add", addr)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	data := []byte(fmt.Sprintf(`{"address":"%s","list":"%s","timeout":"%s"}`, ip, name, duration))

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

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s\n", string(resBody))
}

func delAddress(ip string, duration string, name string) {

	//getAddr := fmt.Sprintf("https://%s/rest/ip/firewall/address-list", addr)
	//delAddr := fmt.Sprintf("https://%s/rest/ip/firewall/address-list/remove", addr)

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }

	// client := &http.Client{Transport: tr}

	fmt.Println("deleting ip", ip, duration, name)

	// get, err := http.Get(getAddr, "application/json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer get.Body.Close()

	// fmt.Println(get.Body)

	// req, err := http.Post(delAddr, "application/json", bytes.NewBuffer(body))
	// defer req.Body.Close()

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Parse()
	fmt.Println("ADDR:", addr)

	//fmt.Println(flag.Args())

	// blacklist := Blacklist{
	// 	Address: flag.Args()[1],
	// 	Timeout: flag.Args()[2],
	// 	List:    flag.Args()[3],
	// }

	if flag.Arg(0) == "add" {
		fmt.Println("Add", flag.Arg(1))
		addAddress(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	} else {
		fmt.Println("deleting ip", flag.Arg(1))
		delAddress(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	}
}
