package main

/*
#include "postgres.h"
#include "fmgr.h"
*/
import "C"
import (
	"crypto/rand"
	"log"

	"github.com/oklog/ulid"
)

//export Ulid
func Ulid() Datum {
	logger := NewErrorLogger("", log.Ltime|log.Lshortfile)
	id, err := ulid.New(ulid.Now(), rand.Reader)
	if err != nil {
		logger.Fatalf("Error: %s", err)
	}
	return toDatum(id.String())
}