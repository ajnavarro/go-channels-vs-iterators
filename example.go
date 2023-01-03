package main

import (
	"context"
	"io"
)

func GetValuesChan(num int) <-chan int {
	result := make(chan int)
	go func() {
		count := 0
		for {
			if count == num {
				break
			}

			count++
			result <- count
		}

		close(result)
	}()

	return result
}

func GetValuesIter(num int) IntIter {
	return &ValueIter{max: num}
}

type ValueIter struct {
	count int
	max   int
}

func (i *ValueIter) Next(ctx context.Context) (int, error) {
	if i.count == i.max {
		return 0, io.EOF
	}
	i.count++
	return i.count, nil
}

func (i *ValueIter) Close() error {
	return nil
}

type IntIter interface {
	// Next retrieves the next element. It will return io.EOF if there is no more elements to get.
	// After retrieving the last element, Close will be automatically called.
	Next(ctx context.Context) (int, error)
	io.Closer
}
