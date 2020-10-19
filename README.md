# Flat-struct
The main purpose of this package is to flatten event structure fields into map[string]string for its fitting into 2 ClickHouse columns: 
 * keys (Array(LowCardinality(String))
 * vals Array(String)

This package allows you to make map[string]string from any structure.

Limitation of solution:
 1. every structure fields must have json tag
 2. structure field value cannot be type of Array, Map, Slice, Interface

## Benchmark
Reflection solution is ~2 times slower than hand coded when make **map[string]string from structure** 

````
BenchmarkWithReflection-4      	 1023620	 1105 ns/op	   512 B/op	   12 allocs/op
BenchmarkWithReflection-4      	  971816	 1114 ns/op	   512 B/op	   12 allocs/op

BenchmarkWithoutReflection-4   	 2381493	  508 ns/op	   560 B/op	   11 allocs/op
BenchmarkWithoutReflection-4   	 2300116	  533 ns/op	   560 B/op	   11 allocs/op
````

## Example
```go
package main

import (
	"fmt"

	"flatstruct/flatstruct"
)

func main() {
	sessionID := "sessionId"
	dataStructure := struct {
		ID        int     `json:"id"`
		Name      string  `json:"name"`
		LastName  string  `json:"last_name"`
		SessionID *string `json:"session_id"`
	}{
		1,
		"name",
		"lastName",
		&sessionID,
	}

	flatStruct, err := flatstruct.StructToFlatMap(dataStructure)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(flatStruct)
	// output: map[id:1 last_name:lastName name:name session_id:sessionId]
}
```
