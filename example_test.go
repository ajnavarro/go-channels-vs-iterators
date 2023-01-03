package main

import (
	"context"
	"io"
	"testing"
)

func BenchmarkChannels(b *testing.B) {
	ch := GetValuesChan(b.N)

	b.ResetTimer()

	last := 0
	for v := range ch {
		last = v
	}

	if last != b.N {
		b.Fatal("bad iterator", last, b.N)
	}
}

func BenchmarkIterators(b *testing.B) {
	iter := GetValuesIter(b.N)
	ctx := context.Background()

	b.ResetTimer()

	last := 0
	for {
		v, err := iter.Next(ctx)
		if err == io.EOF {
			break
		}

		last = v
	}

	if last != b.N {
		b.Fatal("bad iterator", last, b.N)
	}
}
