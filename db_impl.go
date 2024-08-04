package in_memory_db

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type InMemoryDB struct {
	Store map[string]map[string]Value
	InMemoryDBInterface
}

type Value struct {
	value  string
	expiry int
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		Store: make(map[string]map[string]Value),
	}
}

func isKeyPresent(m map[string]map[string]Value, key string) bool {
	_, exist := m[key]
	return exist
}

func isFieldPresent(m map[string]map[string]Value, key, field string) bool {
	if isKeyPresent(m, key) {
		_, exist := m[key][field]
		return exist
	}
	return false
}

/*
	--- Level 1 ---
	1. Set(key, field, value string) - Should insert a
	field-value pair to the record associated with key.
	If the field in the record already exists, replace
	the existing value with the specified value.
	If record doesn't exist, create a new one.

	2. Get(key, field string) *string - Should return
	the value contained within field of record associated
	with key. If record or field doesn't exist, should return nil

	3. Delete(key, field string) bool - Should remove the field
	from the record associated with key. Returns true if the field
	was successfully deleted, and false if the key or the field do
	not exist in the database
*/

func (db *InMemoryDB) Set(key, field, value string) {
	// Custom implementation here
	if !isKeyPresent(db.Store, key) {
		db.Store[key] = make(map[string]Value)
	}
	db.Store[key][field] = Value{value: value}
}

func (db *InMemoryDB) Get(key, field string) *string {
	// Custom implementation here
	if !isFieldPresent(db.Store, key, field) {
		return nil
	}
	val, exist := db.Store[key][field]
	if !exist {
		return nil
	}
	return &val.value
}

func (db *InMemoryDB) Delete(key, field string) bool {
	// Custom implementation here
	if !isFieldPresent(db.Store, key, field) {
		return false
	}
	delete(db.Store[key], field)
	return true
}

/*
	--- Level 2 ---
	1. Scan(key string) []string - Should return a list
	of strings representing the fields of a record associated
	with the key. The returned list should be in the following
	format
		["<field1>(<value1>)" , "<field2>(<value2>)", ...]
	where the fields are lexicographically sorted. If specified
	record doesn't exist, return empty list.

	2. ScanByPrefix(key, prefix string) []string - Should return a list
	of strings representing some fields of a records associated
	with the key. Specifically, only fields that starts with the prefix
	should be included. The returned list should be the same format as
	the Scan operation with the fields sorted in lexicographical order.
*/

func (db *InMemoryDB) Scan(key string) []string {
	if !isKeyPresent(db.Store, key) {
		return []string{}
	}
	var res = []string{}
	for field, val := range db.Store[key] {
		res = append(res, fmt.Sprintf("%s(%s)", field, val.value))
	}
	sort.Strings(res)
	return res
}

func (db *InMemoryDB) ScanByPrefix(key, prefix string) []string {
	if !isKeyPresent(db.Store, key) {
		return []string{}
	}
	var res = []string{}
	for field, val := range db.Store[key] {
		if strings.HasPrefix(field, prefix) {
			res = append(res, fmt.Sprintf("%s(%s)", field, val.value))
		}
	}
	sort.Strings(res)
	return res
}

/*
	--- Level 3 ---
	1. SetAt(key, field, value string, timestamp int) []string -
	Should insert a field-value pair or update the value of the
	field in the record associated with key

	2. SetAtWithTtl(key, field, value string, timestamp, ttl int) []string -
	Should insert a field-value pair or update the value of the
	field in the record associated with key. Also sets its Time-to-Live
	starting at timestamp to be ttl. The ttl is the amount of time that this
	field-value pair should exist in the database, meaning it will be avaialble
	during the interval: [timestamp, timestamp + ttl]

	3. DeleteAt(key, field string, timestamp int) bool
	The same as Delete, but with timestamp of the operation
	specified. Should return true if the field existed and was
	successfully deleted and false if the key didn't exist.

	4. GetAt(key, field string, timestamp int) *string
	The same as Get, but with timestamp of the operation specified

	5. ScanAt(key string, timestamp int) []string
	The same Scan but with the timestamp of the operation specified

	6. ScanPrefixAt(key, prefix string, timestamp int) []string
	The same as ScanPrefix but with the timestamp of the operation specified.
*/

func (db *InMemoryDB) SetAt(key, field, value string, timestamp int) {
	if !isKeyPresent(db.Store, key) {
		db.Store[key] = make(map[string]Value)
	}
	db.Store[key][field] = Value{value: value, expiry: math.MaxInt}
}

func (db *InMemoryDB) SetAtWithTtl(key, field, value string, timestamp, ttl int) {
	if !isKeyPresent(db.Store, key) {
		db.Store[key] = make(map[string]Value)
	}
	db.Store[key][field] = Value{value: value, expiry: timestamp + ttl}
}

func (db *InMemoryDB) DeleteAt(key, field string, timestamp int) bool {
	// Custom implementation here
	if !isFieldPresent(db.Store, key, field) {
		return false
	}
	val := db.Store[key][field]
	delete(db.Store[key], field)
	return timestamp <= val.expiry
}

func (db *InMemoryDB) GetAt(key, field string, timestamp int) *string {
	if !isFieldPresent(db.Store, key, field) {
		return nil
	}
	val := db.Store[key][field]
	if timestamp <= val.expiry {
		return &val.value
	}
	return nil
}

func (db *InMemoryDB) ScanAt(key string, timestamp int) []string {
	if !isKeyPresent(db.Store, key) {
		return []string{}
	}
	var res = []string{}
	for field, val := range db.Store[key] {
		if timestamp <= val.expiry {
			res = append(res, fmt.Sprintf("%s(%s)", field, val.value))
		}
	}
	sort.Strings(res)
	return res
}

func (db *InMemoryDB) ScanPrefixAt(key, prefix string, timestamp int) []string {
	if !isKeyPresent(db.Store, key) {
		return []string{}
	}
	var res = []string{}
	for field, val := range db.Store[key] {
		if strings.HasPrefix(field, prefix) {
			if timestamp <= val.expiry {
				res = append(res, fmt.Sprintf("%s(%s)", field, val.value))
			}
		}
	}
	sort.Strings(res)
	return res
}
