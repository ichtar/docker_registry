package main

import (
	"fmt"
	"strings"
        "net/http"
	"io/ioutil"
	"encoding/json"
	"flag"
	"strconv"
)

func init() {
    flag.StringVar(&address, "u", "http://localhost", "registry url with schme (http/https)")
    flag.IntVar(&port, "p" ,5000, "registry port")
    flag.StringVar(&user, "U" ,"dummy", "username for registry authentication")
    flag.StringVar(&pass, "P" ,"dummy", "passwd for registry authentication")
	// Tie the command-line flag to the intervalFlag variable and
	// set a usage message.
}

func read_http_endpoint(query string) []byte {
    var headerName string = "Authorization"
    var headerValue string = "silly"
    client := &http.Client{}
    req, err := http.NewRequest("GET", query, nil)
    if err != nil {
        panic(err)
    }
    req.SetBasicAuth(user, pass)
    req.Header.Set(headerName,headerValue)

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    return body
}

    var address string 
    var port int
    var user,pass  string

func main() {
    flag.Parse()
    url := strings.Join([]string{address,":",strconv.Itoa(port),"/v2/"},"")
    fmt.Println(url)
   //   url = "http://localhost:5000/v2/"
    catalog := "_catalog"
    tags := "/tags/list"

    type Message struct {
        Repositories []string `json:"repositories"`
    }
    type Message2  struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
    }
    http_catalog := strings.Join([]string{url,catalog},"")
    http_json := read_http_endpoint(http_catalog)
    var decoded Message
    err := json.Unmarshal(http_json, &decoded)
    if err != nil {
        panic(err)
    }
    for _,image := range decoded.Repositories {
        fmt.Println(image)
        http_tags := strings.Join([]string{url,image,tags},"")
	http_json2 := read_http_endpoint(http_tags)
        var decoded2 Message2
	 err := json.Unmarshal(http_json2, &decoded2)
        if err != nil {
            panic(err)
        }
	for _,tag := range decoded2.Tags {
	    fmt.Printf("%s:%s\n",decoded2.Name,tag)
	}
    }

}
