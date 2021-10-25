# Welcome
**pickleDB is a lightweight and simple key-value store.** 
It is inspired by written in Python - [pickleDB](https://github.com/patx/pickledb).

```go
package main

import (
    cucumber "github.com/gopherzz/cucumberdb"
)

func main() {
    db := cucumber.New()
    db.Load("database.jdb")
    
    db.Set("key", "value") // -> true
    db.Get("key") // -> value
    
    db.Dump()
}
```
### New
```go
func New() *Db
    / Return new cucumberdb object
```

### Append
```go
func (db *Db) Append(key, more string) bool
    / Add more to a key's value
```

### DelDb
```go
func (db *Db) DelDb() bool
    / Delete everything from the database
```

### Dump
```go
func (db *Db) Dump() error
    / Save the database from memory to a file specified in Load or Init
```

### Get
```go
fuc (db *Db) Get(key string) interface{}
    / Get the value of a key
```

### GetAll
```go
func (db *Db) GetAll() []string
    / Return a list of all keys in database
```

### Init
```go
func (db *Db) Init(name string) error
    / Create new database file with name, and open it
```

### LAdd
```go
func (db *Db) LAdd(name string, value interface{}) bool
    / Add a value to a list
```

### LCreate
```go
func (db *Db) LCreate(name string) bool
    / Create a list
```

### LGet
```go
func (db *Db) LGet(name string, pos int) interface{}
    / Return one value in a list
```
### LGetAll
```go
func (db *Db) LGetAll(name string) []interface{}
    / Return all values in a list
```

### LPop
```go
func (db *Db) LPop(name string, pos int) interface{}
    / Remove one value in a list and return it
```

### LRem
```go
func (db *Db) LRem(name string) bool
    / Remove a list and all of its values
```

### Load
```go
func (db *Db) Load(database string) error
    / Load a database from a file
```

### Rem
```go
func (db *Db) Rem(key string) bool
    / Delete a key
```

### Set
```go
func (db *Db) Set(key string, value interface{}) bool
    / Set the value of a keyn
```
