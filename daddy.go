package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  sjson "github.com/bitly/go-simplejson"
)

const tokenName = "DADDY_GO_FB_TOKEN"


func main() {
  token := os.Getenv(tokenName)
  if len(token) == 0 {
    log.Fatal(tokenName + " required")
  }

  res, err := http.Get("https://graph.facebook.com/me/photos?access_token=" + token)
  if err != nil {
    log.Fatal(err)
  }
  body, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }

  // fmt.Println(res)
  fbJson, err := sjson.NewJson(body)
  if err != nil {
    log.Fatal(err)
  }
  // fmt.Println(fbJson)
  firstPhoto := fbJson.Get("data").GetIndex(0)
  fmt.Println(firstPhoto)
}
