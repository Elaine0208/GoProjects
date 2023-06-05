package main

import (
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    response, err := http.Get("http://localhost:8080/hello")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(string(responseData))
}
