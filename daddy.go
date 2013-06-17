package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
)

const tokenName = "DADDY_GO_FB_TOKEN"


type FacebookFrom struct {
  Name string "name"
}

type Photo struct {
  Source string "source"
  From FacebookFrom "from"
}

type FacebookJson struct {
  Data []Photo "tag"
}


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
  var fbJson FacebookJson
  error := json.Unmarshal(body, &fbJson)
  if error != nil {
    log.Fatal(error)
  }
  // fmt.Println(fbJson)
  firstPhoto := fbJson.Data[0]
  fmt.Println(firstPhoto)
}
