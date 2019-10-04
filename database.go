package main

import (
  "log"
  driver "github.com/arangodb/go-driver"
  arango_http "github.com/arangodb/go-driver/http"
)

type Database struct {
  db driver.Database
}

type Collection struct {
  col driver.Collection
}

func databaseConnection() Database {
  conn, err := arango_http.NewConnection(arango_http.ConnectionConfig{
      Endpoints: []string{"http://localhost:8529"},
  })
  if err != nil {
      log.Fatalf("Failed to create HTTP connection: %v", err)
  }

  c, err := driver.NewClient(driver.ClientConfig{
      Connection: conn,
      Authentication: driver.BasicAuthentication("root", "Aa55281281"),
  })

  data, err := c.Database(nil, "fantasy_cards")
  if err != nil {
      log.Fatalf("Failed to create database: %v", err)
  }

  db := Database {
    db: data,
  }

  return db
}

func readCollection(db driver.Database, collection string) Collection {

  col, err := db.Collection(nil, collection)
  if driver.IsNotFound(err) {
    col, err = db.CreateCollection(nil, collection, nil)
  } else if err != nil {
		log.Fatalf("Failed to create collection: %v", err)
  }

  table := Collection {
    col: col,
  }

  return table
}
