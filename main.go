package main

/*
#cgo LDFLAGS: -shared

#include "postgres.h"
#include "fmgr.h"
#include "utils/elog.h"
#include "utils/builtins.h"
#include "utils/date.h"
#include "utils/timestamp.h"
#include "datatype/timestamp.h"

#ifdef PG_MODULE_MAGIC
PG_MODULE_MAGIC;
#endif

PG_FUNCTION_INFO_V1(Ulid);
PG_FUNCTION_INFO_V1(UlidToTime);
PG_FUNCTION_INFO_V1(UlidToLocalTime);

Datum void_datum(){
    PG_RETURN_VOID();
}

void elog_notice(char* string) {
    elog(NOTICE, string, "");
}

void elog_error(char* string) {
    elog(ERROR, string, "");
}

Datum get_arg(PG_FUNCTION_ARGS, uint i) {
	return PG_GETARG_DATUM(i);
}

Datum cstring_to_datum(char *val) {
    return CStringGetDatum(cstring_to_text(val));
}

char* datum_to_cstring(Datum val) {
    return DatumGetCString(text_to_cstring((struct varlena *)val));
}

Datum timestamp_to_datum(Timestamp val) {
	return TimestampGetDatum(val);
}

Datum timestamptz_to_datum(TimestampTz val) {
	return TimestampTzGetDatum(val);
}

*/
import "C"
import (
	"log"
	"time"
	"unsafe"
)

func main() {}

type Datum C.Datum

func toDatum(val string) Datum {
	return (Datum)(C.cstring_to_datum(C.CString(val)))
}

func tsToDatum(val time.Time) Datum {
	return (Datum)(C.timestamp_to_datum(C.Timestamp((val.Unix() - 946684800) * int64(C.USECS_PER_SEC))))
}

func tstzToDatum(val time.Time) Datum {
	return (Datum)(C.timestamptz_to_datum(C.TimestampTz((val.Unix() - 946684800) * int64(C.USECS_PER_SEC))))
}

func getArgText(fcinfo *C.FunctionCallInfoBaseData, n int) string {
	return C.GoString(C.datum_to_cstring(C.get_arg((*C.struct_FunctionCallInfoBaseData)(unsafe.Pointer(fcinfo)), C.uint(n))))
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
