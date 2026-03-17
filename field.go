package infonyz

// Field represents a structured log field with a key, type, and value.
type Field struct {
	// Key is the name of the field.
	Key string

	// Val holds the value of the field, which can be of any type.
	Val any
}

// F creates a new Field with the given key and value.
func F(k string, v any) *Field {
	return &Field{
		Key: k,
		Val: v,
	}
}

// String creates a string field with the given key and value.
func String(key string, val string) *Field { return F(key, val) }

// Int64 creates an int64 field with the given key and value.
func Int64(key string, val int64) *Field { return F(key, val) }

// Int creates an int field with the given key and value.
func Int(key string, val int) *Field { return F(key, val) }

// Float64 creates a float64 field with the given key and value.
func Float64(key string, val float64) *Field { return F(key, val) }

// Float32 creates a float32 field with the given key and value.
func Float32(key string, val float32) *Field { return F(key, val) }

// Bool creates a boolean field with the given key and value.
func Bool(key string, val bool) *Field { return F(key, val) }

// Bytes creates a byte slice field with the given key and value.
func Bytes(key string, val []byte) *Field { return F(key, val) }
