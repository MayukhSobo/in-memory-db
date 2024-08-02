package in_memory_db

type InMemoryDBInterface interface {
	Set(key, field, value string)
	Get(key, field string) *string
	Delete(key, field string) bool

	Scan(key string) []string
	ScanByPrefix(key, prefix string) []string

	SetAt(key, field, value string, timestamp int)
	SetAtWithTtl(key, field, value string, timestamp, ttl int)
	DeleteAt(key, field string, timestamp int) bool
	GetAt(key, field string, timestamp int) *string
	ScanAt(key string, timestamp int) []string
	ScanByPrefixAt(key, prefix string, timestamp int) []string
}
