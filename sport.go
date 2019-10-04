package main

type XmlSports struct {
  Sports []Sport `xml:"Sport"`
}

type Sport struct {
  Events []Event `xml:"Event"`
}
