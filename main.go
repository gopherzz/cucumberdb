package cucumberdb

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"reflect"
)

const (
	EXTENSION = ".jdb"
)

type Db struct {
	data     map[string]interface{}
	isOpen   bool
	fileName string
}

func New() *Db {
	return &Db{
		data:   make(map[string]interface{}),
		isOpen: false,
	}
}

func (db *Db) Init(name string) error {
	name = name + EXTENSION
	if err := ioutil.WriteFile(name, []byte("{}"), 0777); err != nil {
		return err
	}
	db.isOpen = true
	db.fileName = name
	return nil
}

func (db *Db) Load(database string) error {
	file, err := ioutil.ReadFile(database)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &db.data)
	if err != nil {
		return err
	}
	db.isOpen = true
	db.fileName = database
	return nil
}

func (db *Db) Dump() error {
	if !db.isOpen {
		return errors.New("Error: Database not opened!")
	}
	eData, err := json.Marshal(db.data)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(db.fileName, eData, fs.ModePerm); err != nil {
		return err
	}
	return nil
}

func (db *Db) Set(key string, value interface{}) bool {
	if db.isOpen {
		db.data[key] = value
		return true
	}
	return false
}

func (db *Db) Get(key string) interface{} {
	if db.isOpen {
		return db.data[key]
	}
	return nil
}

func (db *Db) GetAll() []string {
	if !db.isOpen {
		return nil
	}
	keys := []string{}
	for k, _ := range db.data {
		keys = append(keys, k)
	}
	return keys
}

func (db *Db) Rem(key string) bool {
	if !db.isOpen {
		return false
	}
	if _, ok := db.data[key]; ok {
		delete(db.data, key)
		return true
	}
	return false
}

func (db *Db) Append(key, more string) bool {
	if !db.isOpen {
		return false
	}
	switch reflect.ValueOf(db.data[key]).Kind() {
	case reflect.String:
		db.data[key] = db.data[key].(string) + more
		return true
	}
	return false
}

func (db *Db) LCreate(name string) bool {
	if !db.isOpen {
		return false
	}
	db.Set(name, []interface{}{})
	return true
}

func (db *Db) LAdd(name string, value interface{}) bool {
	if !db.isOpen {
		return false
	}
	switch reflect.ValueOf(db.data[name]).Kind() {
	case reflect.Slice:
		db.data[name] = append(db.data[name].([]interface{}), value)
		return true
	}
	return false
}

func (db *Db) LGet(name string, pos int) interface{} {
	if !db.isOpen {
		return nil
	}
	switch reflect.ValueOf(db.data[name]).Kind() {
	case reflect.Slice:
		return db.data[name].([]interface{})[pos]
	}
	return nil
}

func (db *Db) LPop(name string, pos int) interface{} {
	if !db.isOpen {
		return nil
	}
	switch reflect.ValueOf(db.data[name]).Kind() {
	case reflect.Slice:
		val := db.data[name].([]interface{})[pos]
		db.data[name] = append(db.data[name].([]interface{})[:pos], db.data[name].([]interface{})[pos+1:]...)
		return val
	}
	return nil
}

func (db *Db) LGetAll(name string) []interface{} {
	if !db.isOpen {
		return nil
	}
	res := []interface{}{}
	switch reflect.ValueOf(db.data[name]).Kind() {
	case reflect.Slice:
		for _, v := range db.data[name].([]interface{}) {
			res = append(res, v)
		}
	}
	return res
}

func (db *Db) LRem(name string) bool {
	if !db.isOpen {
		return false
	}
	if _, ok := db.data[name]; !ok {
		return false
	}
	delete(db.data, name)
	return true
}

func (db *Db) DelDb() bool {
	if db.isOpen {
		db.data = make(map[string]interface{})
		return true
	}
	return false
}

type User struct {
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}
