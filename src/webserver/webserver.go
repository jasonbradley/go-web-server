package main

import (
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "util/reader"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadFile(reader.ServerSettings.Webroot + "/" + r.URL.Path[1:])

    if err != nil {
        io.WriteString(w, "File not found.")
    } else {
        io.WriteString(w, string(data))
    }
}

func main() {

    //Load the server settings
    reader.Load()

    fmt.Printf("\nStarted server on port %s\n", reader.ServerSettings.Port)

    http.HandleFunc("/", viewHandler)

    http.ListenAndServe(":" + reader.ServerSettings.Port, nil)
}
