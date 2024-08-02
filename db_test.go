package in_memory_db

import (
	"reflect"
	"testing"
)

func TestLevel1(t *testing.T) {
	db := NewInMemoryDB()

	db.Set("user1", "name", "Alice")
	db.Set("user1", "age", "30")
	if got := db.Get("user1", "name"); !reflect.DeepEqual(*got, "Alice") {
		t.Errorf("Expected 'Alice', got %v", got)
	}
	if got := db.Get("user1", "age"); !reflect.DeepEqual(*got, "30") {
		t.Errorf("Expected '30', got %v", got)
	}

	db.Set("user1", "age", "31")
	if got := db.Get("user1", "age"); !reflect.DeepEqual(*got, "31") {
		t.Errorf("Expected '31', got %v", got)
	}

	if got := db.Get("user1", "address"); got != nil {
		t.Errorf("Expected nil, got %v", got)
	}

	if got := db.Delete("user1", "age"); !got {
		t.Errorf("Expected true, got %v", got)
	}
	if got := db.Get("user1", "age"); got != nil {
		t.Errorf("Expected nil, got %v", got)
	}

	if got := db.Delete("user1", "age"); got {
		t.Errorf("Expected false, got %v", got)
	}

	if got := db.Delete("user2", "name"); got {
		t.Errorf("Expected false, got %v", got)
	}
}

func TestLevel2(t *testing.T) {
	db := NewInMemoryDB()

	db.Set("user1", "name", "Alice")
	db.Set("user1", "age", "30")
	db.Set("user1", "address", "Wonderland")

	expected := []string{"address(Wonderland)", "age(30)", "name(Alice)"}
	if got := db.Scan("user1"); !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	expected = []string{"address(Wonderland)"}
	if got := db.ScanByPrefix("user1", "add"); !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	if got := db.Scan("user2"); !reflect.DeepEqual(got, []string{}) {
		t.Errorf("Expected empty slice, got %v", got)
	}

	if got := db.ScanByPrefix("user2", "name"); !reflect.DeepEqual(got, []string{}) {
		t.Errorf("Expected empty slice, got %v", got)
	}

	if got := db.ScanByPrefix("user1", "xyz"); !reflect.DeepEqual(got, []string{}) {
		t.Errorf("Expected empty slice, got %v", got)
	}
}

func TestLevel3(t *testing.T) {
	db := NewInMemoryDB()

	db.SetAt("user1", "name", "Alice", 1)
	db.SetAt("user1", "age", "30", 2)
	db.SetAt("user1", "address", "Wonderland", 3)
	db.SetAtWithTtl("user1", "tempField", "tempValue", 4, 2)

	if got := db.GetAt("user1", "name", 1); !reflect.DeepEqual(*got, "Alice") {
		t.Errorf("Expected 'Alice', got %v", got)
	}
	if got := db.GetAt("user1", "age", 2); !reflect.DeepEqual(*got, "30") {
		t.Errorf("Expected '30', got %v", got)
	}

	if got := db.GetAt("user1", "tempField", 7); got != nil {
		t.Errorf("Expected nil, got %v", got)
	}

	if got := db.DeleteAt("user1", "age", 2); !got {
		t.Errorf("Expected true, got %v", got)
	}
	if got := db.GetAt("user1", "age", 3); got != nil {
		t.Errorf("Expected nil, got %v", got)
	}

	if got := db.DeleteAt("user1", "tempField", 8); got {
		t.Errorf("Expected false, got %v", got)
	}

	expected := []string{"address(Wonderland)", "name(Alice)"}
	if got := db.ScanAt("user1", 2); !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	expected = []string{"address(Wonderland)"}
	if got := db.ScanByPrefixAt("user1", "add", 3); !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
