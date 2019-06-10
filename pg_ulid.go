package main

/*
#include "postgres.h"
#include "fmgr.h"
*/
import "C"
import (
	"crypto/rand"
	"github.com/oklog/ulid"
	"log"
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

//export UlidToTime
func UlidToTime(fcinfo *C.FunctionCallInfoData) Datum {
	logger := NewErrorLogger("", log.Ltime|log.Lshortfile)
	id, err := ulid.Parse(getArgText(fcinfo, 0))
	if err != nil {
		logger.Fatalf("Error: %s", err)
	}
	return tsToDatum(ulid.Time(id.Time()))
}

//export UlidToLocalTime
func UlidToLocalTime(fcinfo *C.FunctionCallInfoData) Datum {
	logger := NewErrorLogger("", log.Ltime|log.Lshortfile)
	id, err := ulid.Parse(getArgText(fcinfo, 0))
	if err != nil {
		logger.Fatalf("Error: %s", err)
	}
	return tstzToDatum(ulid.Time(id.Time()).UTC())
}
