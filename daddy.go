package main

import (
  fb "github.com/huandu/facebook"
  "fmt"
  "log"
  "os"
)

const tokenName = "DADDY_GO_FB_TOKEN"

type Photo struct {
  Source string `facebook:",required"`
}


func main() {
  token := os.Getenv(tokenName)
  if len(token) == 0 {
    log.Fatal(tokenName + " required")
  }

  res, _ := fb.Get("/me/photos", fb.Params{
    "access_token": token,
  })

  // fmt.Println(res)
  var firstPhoto Photo
  res.DecodeField("data.0", &firstPhoto)
  fmt.Println(firstPhoto)
}
