package in_memory_db

type InMemoryDB struct {
	InMemoryDBInterface
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{}
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
}

func (db *InMemoryDB) Get(key, field string) *string {
	// Custom implementation here
	return nil
}

func (db *InMemoryDB) Delete(key, field string) bool {
	// Custom implementation here
	return true
}
