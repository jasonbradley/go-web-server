package reader

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
)

type Config struct {
    Webroot string
    Port string
}

var ServerSettings Config

func SetServerSettings(data []byte) {
    err1 := json.Unmarshal(data, &ServerSettings)

    if err1 != nil {
        fmt.Println("error: ", err1)
    }
}

func Load() {
    data, err := ioutil.ReadFile("/var/apps/gowebserver/config.json")

    if err != nil {
        fmt.Printf("Error: %s", err)
    }

    SetServerSettings(data)
}
