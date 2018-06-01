package main

import (
	"fmt"
        "net/http"
	"io/ioutil"
	"encoding/json"
)

func main() {
    type Message struct {
        Repositories []string `json:"repositories"`
    }
    resp, err := http.Get("http://localhost:5000/v2/_catalog")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    var decoded Message
    err = json.Unmarshal(body, &decoded)
    if err != nil {
        panic(err)
    }
    for _,image := range decoded.Repositories {
        fmt.Println(image)
    }

}
