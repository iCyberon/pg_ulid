# pg_ulid
Universally Unique Lexicographically Sortable Identifier (ULID) for PostgreSQL

Experimental PostgreSQL extension for generating ULIDs built by go ([cgo](https://golang.org/cmd/cgo/)). 

## Install
You need to have go and compiler (clang/gcc) installed on your system.

```sh
$ git clone https://github.com/icyberon/pg_ulid
$ cd pg_ulid
$ make
$ sudo make install
```

This will build and install the extension.

```sql
CREATE EXTENSION pg_ulid;
```

## Usage

Now you can start generating ULIDs from PostgreSQL by calling `ulid`. Function returns one column `TEXT`
```sql
SELECT ulid(); -- 01C1P15MBWYBWDG2WYRG08VCFR
SELECT pg_typeof(ulid()); -- text
```

## Todo

- [ ] Add binary support
- [ ] Add tests
