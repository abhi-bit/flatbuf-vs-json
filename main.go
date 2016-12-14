package main

import (
	"fmt"

	"github.com/abhi-bit/flatbuf-vs-json/dcpevent"
	"github.com/google/flatbuffers/go"
)

func MakeMutation(b *flatbuffers.Builder, key, val []byte,
	cas, ttl uint64) []byte {

	// reuse already allocated buffer
	b.Reset()

	keyPosition := b.CreateByteString(key)
	valPosition := b.CreateByteString(val)

	dcpevent.MutationStart(b)

	dcpevent.MutationAddKey(b, keyPosition)
	dcpevent.MutationAddKey(b, valPosition)

	dcpevent.MutationAddCas(b, cas)
	dcpevent.MutationAddTtl(b, ttl)

	mutationPosition := dcpevent.MutationEnd(b)

	b.Finish(mutationPosition)

	return b.Bytes[b.Head():]
}

func ReadMutation(buf []byte) (key, val []byte, cas, ttl uint64) {
	// init mutation reader from supplied buffer
	mutation := dcpevent.GetRootAsMutation(buf, 0)

	key = mutation.Key()
	val = mutation.Value()
	cas = mutation.Cas()
	ttl = mutation.Ttl()

	return
}

func main() {
	builder := flatbuffers.NewBuilder(0)
	buf := MakeMutation(builder,
		[]byte("pymc0"), []byte("{\"city\":\"BLR\"}"), 1234567890, 0)

	key, val, cas, ttl := ReadMutation(buf)
	fmt.Printf("key: %s val: %s cas: %d ttl: %d\n", key, val, cas, ttl)
}
