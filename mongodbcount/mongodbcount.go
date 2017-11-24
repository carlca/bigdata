package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/mgo.v2"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("%+v", errors.WithStack(err))
		os.Exit(1) // or anything else ...
	}
}

func main() {
	// establish MongoDB session
	session, err := mgo.Dial("127.0.0.1")
	checkError(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	// get collection count
	coll := session.DB("Companies").C("Companies")
	count, err := coll.Count()
	checkError(err)
	p := message.NewPrinter(language.English)
	p.Printf("\nRecord count for Companies: %d", count)
}
