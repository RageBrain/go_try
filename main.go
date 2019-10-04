package main

import (
    "fmt"
    "log"
)

func main() {
  db := databaseConnection().db

  col := readCollection(db, "tournaments").col

  match := readCollection(db, "matches").col

  events := parseXml()

  for _, v := range events {
    if tournamentExist(v.Id) {
      fmt.Println("Event with id %v already exist %v", v.Id, v.Name)
    } else {
      tournament := Tournament{
          Title:    v.Name,
          Category: v.Category,
          EventId:  v.Id,
      }
      _, err := col.CreateDocument(nil, tournament)
      if err != nil {
          log.Fatalf("Failed to create document: %v", err)
      }
      for _,m := range v.Matches {
        _, err := match.CreateDocument(nil, m)
        if err != nil {
            log.Fatalf("Failed to create document: %v", err)
        }
      }
    }
  }
}
