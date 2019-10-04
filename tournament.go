package main

import (
    "log"

    driver "github.com/arangodb/go-driver"
)

type Tournament struct {
  Title    string `json:"title"`
  Category string `json:"category"`
  EventId  string `json:"event_id"`
}

func listOfTournaments() []string {
  var list []string
  db := databaseConnection().db
	query := "FOR t IN tournaments RETURN t.EventId"
  cursor, err := db.Query(nil, query, nil)
	if err != nil {
		log.Fatalf("Failed to open document: %v", err)
	}
  defer cursor.Close()

  fmt.Println("dasdasdasdasdasdasd+================================")
  fmt.Println(cursor)

  for {
		var tournament Tournament
		_, err := cursor.ReadDocument(nil, &tournament)
		if driver.IsNoMoreDocuments(err) {
      return list
    } else if err != nil {
        log.Fatalf("Something went wrong: %v", err)
    }
    list = append(list, tournament.EventId)
	}
}

func tournamentExist(eventId string) bool {
  list := listOfTournaments()
  for _, n := range list {
    if eventId == n {
      return true
    }
  }
  return false
}
