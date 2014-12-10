package reader_test

import (
    "testing"
    "util/reader"
    "io/ioutil"
)

func TestSetServerSettings(t *testing.T) {
    data, err := ioutil.ReadFile("./config.json")

    if err != nil {
        t.Error(err)
    }

    reader.SetServerSettings(data)

    if reader.ServerSettings.Webroot != "webroot_test" {
        t.Error("Failed to set webroot")
    }

    if reader.ServerSettings.Port != "port_test" {
        t.Error("Failed to set port")
    }

}
