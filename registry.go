package main

import (
	"fmt"
	"strings"
        "net/http"
	"io/ioutil"
	"encoding/json"
)

func read_http_endpoint(query string) []byte {
    resp, err := http.Get(query)
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


func main() {
    type Message struct {
        Repositories []string `json:"repositories"`
    }
    type Message2  struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
    }
    http_json := read_http_endpoint("http://localhost:5000/v2/_catalog")
    var decoded Message
    err := json.Unmarshal(http_json, &decoded)
    if err != nil {
        panic(err)
    }
    for _,image := range decoded.Repositories {
        http_tags := strings.Join([]string{"http://localhost:5000/v2/",image,"/tags/list"},"")
	http_json2 := read_http_endpoint(http_tags)
        var decoded2 Message2
	 err := json.Unmarshal(http_json2, &decoded2)
        if err != nil {
            panic(err)
        }
//        fmt.Println(decoded2.Name)
	for _,tag := range decoded2.Tags {
	    fmt.Printf("%s:%s\n",decoded2.Name,tag)
	}
    }

}
