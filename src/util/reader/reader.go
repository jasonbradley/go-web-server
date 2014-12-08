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

func Load() {
    data, err := ioutil.ReadFile("/var/apps/gowebserver/config.json")

    if err != nil{
        fmt.Print("Error:",err)
    }

    err1 := json.Unmarshal(data, &ServerSettings)

    if err1 != nil {
        fmt.Println("error: ", err1)
    }
}

