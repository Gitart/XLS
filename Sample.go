/*
   Templates
*/

package main

import (
    "encoding/json"
    "fmt"

    r "github.com/dancannon/gorethink"
)

var session *r.Session

// Struct tags are used to map struct fields to fields in the database
type Person struct {
    Id    string `gorethink:"id,omitempty"`
    Name  string `gorethink:"name"`
    Place string `gorethink:"place"`
}

func init() {
    var err error
    session, err = r.Connect(r.ConnectOpts{
        Address:  "localhost:28015",
        Database: "test",
    })
    if err != nil {
        fmt.Println(err)
        return
    }
}

func main() {
    // Create a table
    createTable()

    // Insert a record
    id := insertRecord()

    if id != "" {
        // Update a record
        updateRecord(id)
    }

    // Fetch one
    fetchOneRecord()

    // Record count
    recordCount()

    // Fetch all
    fetchAllRecords()

    if id != "" {
        // Delete a record
        deleteRecord(id)
    }
}

func createTable() {
    result, err := r.DB("test").TableCreate("people").RunWrite(session)
    if err != nil {
        fmt.Println(err)
    }

    printStr("*** Create table result: ***")
    printObj(result)
    printStr("\n")
}

func insertRecord() string {
    var data = map[string]interface{}{
        "Name":  "David Davidson",
        "Place": "Somewhere",
    }

    result, err := r.Table("people").Insert(data).RunWrite(session)
    if err != nil {
        fmt.Println(err)
        return ""
    }

    printStr("*** Insert result: ***")
    printObj(result)
    printStr("\n")

    return result.GeneratedKeys[0]
}

func updateRecord(id string) {
    var data = map[string]interface{}{
        "Name":  "Steve Stevenson",
        "Place": "Anywhere",
    }

    result, err := r.Table("people").Get(id).Update(data).RunWrite(session)
    if err != nil {
        fmt.Println(err)
        return
    }

    printStr("*** Update result: ***")
    printObj(result)
    printStr("\n")
}

func fetchOneRecord() {
    cursor, err := r.Table("people").Run(session)
    if err != nil {
        fmt.Println(err)
        return
    }

    var person interface{}
    cursor.One(&person)
    cursor.Close()

    printStr("*** Fetch one record: ***")
    printObj(person)
    printStr("\n")
}

func recordCount() {
    cursor, err := r.Table("people").Count().Run(session)
    if err != nil {
        fmt.Println(err)
        return
    }

    var cnt int
    cursor.One(&cnt)
    cursor.Close()

    printStr("*** Count: ***")
    printObj(cnt)
    printStr("\n")
}

func fetchAllRecords() {
    rows, err := r.Table("people").Run(session)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Read records into persons slice
    var persons []Person
    err2 := rows.All(&persons)
    if err2 != nil {
        fmt.Println(err2)
        return
    }

    printStr("*** Fetch all rows: ***")
    for _, p := range persons {
        printObj(p)
    }
    printStr("\n")
}

func deleteRecord(id string) {
    result, err := r.Table("people").Get(id).Delete().Run(session)
    if err != nil {
        fmt.Println(err)
        return
    }

    printStr("*** Delete result: ***")
    printObj(result)
    printStr("\n")
}

func printStr(v string) {
    fmt.Println(v)
}

func printObj(v interface{}) {
    vBytes, _ := json.Marshal(v)
    fmt.Println(string(vBytes))
}


















chages


package main

import (
	"code.google.com/p/gcfg"
	"fmt"
	r "github.com/dancannon/gorethink"
	irc "github.com/fluffle/goirc/client"
	"log"
	"strings"
)

type Config struct {
	IRC struct{ Host, Channel, Nickname string }
	DB  struct{ Host string }
}

type Issue struct {
	Description, Type string
}

type Server struct {
	Name, Status string
}

func main() {
	quit := make(chan bool, 1)

	var config Config
	if err := gcfg.ReadFileInto(&config, "config.gcfg"); err != nil {
		log.Fatal("Couldn't read configuration")
	}

	db, err := r.Connect(r.ConnectOpts{Address: config.DB.Host})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	ircConf := irc.NewConfig(config.IRC.Nickname)
	ircConf.Server = config.IRC.Host
	bot := irc.Client(ircConf)

	bot.HandleFunc("connected", func(conn *irc.Conn, line *irc.Line) {
		log.Println("Connected to IRC server", config.IRC.Host)
		conn.Join(config.IRC.Channel)
	})

	bot.HandleFunc("privmsg", func(conn *irc.Conn, line *irc.Line) {
		log.Println("Received:", line.Nick, line.Text())
		if strings.HasPrefix(line.Text(), config.IRC.Nickname) {
			command := strings.Split(line.Text(), " ")[1]
			switch command {
			case "quit":
				log.Println("Received command to quit")
				quit <- true
			}
		}
	})

	log.Println("Connecting to IRC server", config.IRC.Host)
	if err := bot.Connect(); err != nil {
		log.Fatal("IRC connection failed:", err)
	}

	issues, _ := r.Db("rethinkdb").Table("current_issues").Filter(
		r.Row.Field("critical").Eq(true)).Changes().Field("new_val").Run(db)

	go func() {
		var issue Issue
		for issues.Next(&issue) {
			if issue.Type != "" {
				text := strings.Split(issue.Description, "\n")[0]
				message := fmt.Sprintf("(%s) %s ...", issue.Type, text)
				bot.Privmsg(config.IRC.Channel, message)
			}
		}
	}()

	<-quit
}

















