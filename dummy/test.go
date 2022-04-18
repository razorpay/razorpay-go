package main

import (
  "fmt"
 // "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://api.razorpay.com/v1/virtual_accounts/va_JGrZBme0FqgsVh/allowed_payers/ba_JHBm6ejKquIUjS"
  method := "DELETE"

//   payload := strings.NewReader(`{
//   "types": [
//     "vpa"
//   ],
//   "vpa": {
//     "descriptor": "gausikumar"
//   }
// }`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "Basic cnpwX3Rlc3RfazZ1TDg5N1ZQQnoyMHE6RW5MczIxTTQ3QmxsUjNYOFBTRnRqdGJk")
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}