MODULE_big = pg_ulid
OBJS       = pg_ulid.so

EXTENSION = pg_ulid
DATA = pg_ulid--1.0.sql

EXTRA_CLEAN = pg_ulid.h

PG_CONFIG = pg_config
PGXS := $(shell $(PG_CONFIG) --pgxs)
include $(PGXS)

LDFLAGS :=
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
    LDFLAGS = -Wl,-undefined,dynamic_lookup
else
    LDFLAGS = -Wl,--unresolved-symbols=ignore-all
endif

export CGO_CFLAGS = -I$(shell $(PG_CONFIG) --includedir-server)
export CGO_LDFLAGS = $(LDFLAGS)

pg_ulid.so: main.go pg_ulid.go
	go build -buildmode=c-shared -o $@ $^
