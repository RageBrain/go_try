package main

type Event struct {
  Name     string  `xml:"Name,attr"`
  Id       string  `xml:"ID,attr"`
  Category string  `xml:"CategoryID,attr"`
  Matches  []Match `xml:"Match"`
}
