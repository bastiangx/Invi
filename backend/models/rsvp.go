package models

import (
  "database/sql"
  "log"
)

type Rsvp struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Email string `json:"email"`
  EventId int `json:"event_id"`
}

// func GetDB() returns a pointer to a sql.DB object
func GetDB() *sql.DB {
  db, err := sql.Open("sqlite3", "./db/rsvp.db")
  if err != nil {
    log.Fatal(err)
  }
  return db
}
