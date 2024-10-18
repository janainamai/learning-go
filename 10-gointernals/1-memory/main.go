package main

import (
	"fmt"
	"sync"
)

type (
	span struct {
		size      int
		allocated bool
	}

	mheap struct {
		spans []*span
		lock  sync.Mutex
	}

	mcentral struct {
		spans []*span
		lock  sync.Mutex
	}

	mcache struct {
		spans []*span
	}
)

func NewHeap(size int) *mheap {
	h := &mheap{}

	for i := 0; i < size; i++ {
		h.spans = append(h.spans, &span{size: i + 1})
	}

	return h
}

func (h *mheap) getSpanFromMHeap(size int) *span {
	h.lock.Lock()
	defer h.lock.Unlock()

	for _, span := range h.spans {
		if span.size == size && !span.allocated {
			span.allocated = true
			return span
		}
	}

	return nil
}

func (c *mcentral) getSpanFromMCentral(size int) *span {
	c.lock.Lock()
	defer c.lock.Unlock()

	for _, span := range c.spans {
		if span.size == size && !span.allocated {
			span.allocated = true
			return span
		}
	}

	return nil
}

func (ca *mcache) getSpanFromMCache(size int) *span {
	for _, span := range ca.spans {
		if span.size == size && !span.allocated {
			span.allocated = true
			return span
		}
	}

	return nil
}

func main() {
	heap := NewHeap(10)
	mcentral := &mcentral{spans: heap.spans}
	mcache := &mcache{}

	requestedSize := 5
	// requestedSize := 20

	requestSpan := mcache.getSpanFromMCache(requestedSize)
	if requestSpan == nil {
		requestSpan = mcentral.getSpanFromMCentral(requestedSize)
	}
	if requestSpan == nil {
		requestSpan = heap.getSpanFromMHeap(requestedSize)
	}
	if requestSpan == nil {
		panic("Out of memory")
	}

	fmt.Println("Allocated span")
}
