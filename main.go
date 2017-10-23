package main

import (
  "github.com/jawher/mow.cli"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Response struct {
  Message string
}

func main() {
  app := cli.App("dummyc", "dummy client")
  server := app.StringArg("SERVER", "http://127.0.0.1:8080", "Address of server to connect to")
  resp, err := http.Get(*server + "/ping")
  if err != nil {
    panic(err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  fmt.Println(ParseResponse(body))
}

func ParseResponse(body []byte) string {
  data := &Response{}
  err := json.Unmarshal(body, data)
  if err != nil {
    panic(err)
  }
  return data.Message
}
