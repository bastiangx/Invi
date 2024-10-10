BigFish is a super simple website written in Go that focuses on a super simple
RSVP feature for our parties and events. It's a great way to get a headcount for
the people who want to let us know they're coming or not.

---

## Features

- Has a simple Form to put your name and email + a submit button
- Show's the number of people who have RSVP'd already
- The event name, title, time and location are shown and new events can be added 
easily (a parser grabs the event details from a db file, if not empty, adds it to the html)
- The RSVP's are stored in a db file and can be easily accessed and parsed
- each RSVP gets a unique ID for enumeration and deletion.

---

js + simple css frontend
golang (gin-gonic) Backend
sqlite3 Database
Netlify for deployment

---

## events

Two types of events here, A generic event that comes from `rhul-scraper` and a custom event that can be added by the authors (houseparties, etc)
- the generic event is parsed from a `db` file and displayed on the website (event title, time, location, description)
- custom events are just added to the database and displayed on the website if needed.

## Database

using SQLite3 for the database, the database is stored in the `db` folder and is
called `rsvp.db`. The database has a single table called `rsvp` with the following
columns:

```
id INTEGER PRIMARY KEY
name TEXT
email TEXT
event TEXT
```

---

## Backend

Go for the serverless functions (Netlify Functions) to handle form submissions and interact with the database.

