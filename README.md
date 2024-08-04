## Level 1
1. `Set(key, field, value string)` - Should insert a
field-value pair to the record associated with key.
If the field in the record already exists, replace
the existing value with the specified value.
If record doesn't exist, create a new one.

2. `Get(key, field string) *string` - Should return
the value contained within field of record associated
with key. If record or field doesn't exist, should return nil

3. `Delete(key, field string) bool` - Should remove the field
from the record associated with key. Returns true if the field
was successfully deleted, and false if the key or the field do
not exist in the database

## Level 2

1. `Scan(key string) []string` - Should return a list
of strings representing the fields of a record associated
with the key. The returned list should be in the following
format
  `["<field1>(<value1>)" , "<field2>(<value2>)", ...]`
where the fields are lexicographically sorted. If specified
record doesn't exist, return empty list.

2. `ScanByPrefix(key, prefix string) []string` - Should return a list
of strings representing some fields of a records associated
with the key. Specifically, only fields that starts with the prefix
should be included. The returned list should be the same format as
the Scan operation with the fields sorted in lexicographical order.

## Level 3
1. `SetAt(key, field, value string, timestamp int) []string` -
Should insert a field-value pair or update the value of the
field in the record associated with key

2. `SetAtWithTtl(key, field, value string, timestamp, ttl int) []string` -
Should insert a field-value pair or update the value of the
field in the record associated with key. Also sets its Time-to-Live
starting at timestamp to be ttl. The ttl is the amount of time that this
field-value pair should exist in the database, meaning it will be avaialble
during the interval: [timestamp, timestamp + ttl]

3. `DeleteAt(key, field string, timestamp int) bool`
The same as Delete, but with timestamp of the operation
specified. Should return true if the field existed and was
successfully deleted and false if the key didn't exist.

4. `GetAt(key, field string, timestamp int) *string`
The same as Get, but with timestamp of the operation specified

5. `ScanAt(key string, timestamp int) []string`
The same Scan but with the timestamp of the operation specified

6. `ScanPrefixAt(key, prefix string, timestamp int) []string`
The same as ScanPrefix but with the timestamp of the operation specified.
