package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
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
	// empty collection
	coll := session.DB("Companies").C("Companies")
	coll.RemoveAll(nil)
}
