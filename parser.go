package main

import (
  "net/http"
  "io/ioutil"
  "encoding/xml"
)
func parseXml() []Event {
  resp, _ := http.Get("http://sportsbookfeed.com/sportxml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  resp.Body.Close()

  var bet XmlSports
  xml.Unmarshal(bytes, &bet)

  return bet.Sports[0].Events
}
