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
