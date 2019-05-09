package main

import "bytes"

type BufferPool struct {
	c chan *bytes.Buffer
	cap int
}

func NewBufferPool(size int) (bp *BufferPool) {
	return &BufferPool{
		c: make(chan *bytes.Buffer, size),
	}
}

func (bp *BufferPool) Get() (b *bytes.Buffer) {
	select {
	case b = <-bp.c:
		// 嘗試拿取
	default:
		// 未拿到，自己建立一個
		b = bytes.NewBuffer(make([]byte, 0, bp.cap))
	}
}

func (bp *BufferPool) Put(b *bytes.Buffer) {
	b.Reset()

	select {
		// 嘗試還回去
	case bp.c <- b:
		// 放不進去則丟掉
	default:

	}
}
