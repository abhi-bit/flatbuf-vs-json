package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/flatbuffers/go"
)

func BenchmarkFlatbufRead(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	key := []byte("pymc0")
	val := []byte("{\"city\":\"BLR\"}")
	buf := MakeMutation(builder, key, val, 1234567890, 0)
	b.SetBytes(int64(len(buf)))
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		rKey, rVal, _, _ := ReadMutation(buf)

		bytes.Equal(rKey, key)
		bytes.Equal(rVal, val)
	}
}

func BenchmarkFlatbufWrite(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		builder.Reset()
		buf := MakeMutation(builder,
			[]byte("pymc0"), []byte("{\"city\":\"BLR\"}"), 1234567890, 0)
		if i == 0 {
			b.SetBytes(int64(len(buf)))
		}
	}
}

type doc struct {
	City string `json:"city"`
}

func BenchmarkJSONRead(b *testing.B) {
	var d doc
	buf := []byte("{\"city\":\"BLR\"}")
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(buf, &d)
		if err != nil {
			fmt.Errorf("err: %s\n", err.Error())
		}
		if i == 0 {
			b.SetBytes(int64(len(buf)))
		}
	}
}

func BenchmarkJSONWrite(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf, err := json.Marshal([]byte("{\"city\":\"BLR\"}"))
		if err != nil {
			fmt.Errorf("err: %s\n", err.Error())
		}
		if i == 0 {
			b.SetBytes(int64(len(buf)))
		}
	}
}
