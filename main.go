package main

/*
#include "postgres.h"
#include "fmgr.h"
#include "utils/elog.h"
#include "utils/builtins.h"

#ifdef PG_MODULE_MAGIC
PG_MODULE_MAGIC;
#endif

PG_FUNCTION_INFO_V1(Ulid);

void elog_notice(char* string) {
    elog(NOTICE, string, "");
}

void elog_error(char* string) {
    elog(ERROR, string, "");
}

Datum cstring_to_datum(char *val) {
    return CStringGetDatum(cstring_to_text(val));
}
*/
import "C"
import "log"

func main() {}

type Datum C.Datum

func toDatum(val string) Datum {
	return (Datum)(C.cstring_to_datum(C.CString(val)))
}

type elogLevel int

const (
	noticeLevel elogLevel = iota
	errorLevel
)

type elog struct {
	Level elogLevel
}

func (e *elog) Write(p []byte) (n int, err error) {
	switch e.Level {
	case noticeLevel:
		C.elog_notice(C.CString(string(p)))
	case errorLevel:
		C.elog_error(C.CString(string(p)))
	}
	return len(p), nil
}

func NewErrorLogger(prefix string, flag int) *log.Logger {
	return log.New(&elog{Level: errorLevel}, prefix, flag)
}