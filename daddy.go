package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  simplejson "github.com/bitly/go-simplejson"
)

const tokenName = "DADDY_GO_FB_TOKEN"


func main() {
  fbJson := getPhotos()
  // fmt.Println(fbJson)
  firstPhoto := fbJson.Get("data").GetIndex(0)
  fmt.Println(firstPhoto)
}

func getPhotos() *simplejson.Json {
  token := os.Getenv(tokenName)
  if len(token) == 0 {
    log.Fatal(tokenName + " required")
  }

  body := request("https://graph.facebook.com/me/photos?access_token=" + token)
  fbJson, err := simplejson.NewJson(body)
  if err != nil {
    log.Fatal(err)
  }

  return fbJson
}

func request(url string) []byte {
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  body, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  // fmt.Println(res)
  return body
}
