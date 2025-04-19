package xsync

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	c := NewCond(&sync.Mutex{})

	ready := 0

	for i := 0; i < 10; i++ {
		idx := i
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			c.L.Lock()
			ready++
			c.L.Unlock()

			t.Logf("ready: %d", i)
			c.Broadcast()
		}(idx)
	}

	c.L.Lock()
	for ready != 10 {
		_ = c.Wait(context.Background())
		t.Logf("waiter wake up once")
	}
	c.L.Unlock()

	t.Logf("all waiter wake up")
}

func benchmarkCond(b *testing.B, waiterCnt int) {
	c := NewCond(&sync.Mutex{})
	done := make(chan bool)
	id := 0

	for r := 0; r < waiterCnt+1; r++ {
		go func() {
			for i := 0; i < b.N; i++ {
				c.L.Lock()
				if id == -1 {
					c.L.Unlock()
					break
				}

				id++

				if id == waiterCnt+1 {
					id = 0
					c.Broadcast()
				} else {
					_ = c.Wait(context.Background())
				}

				c.L.Unlock()
			}

			c.L.Lock()
			id = -1
			c.Broadcast()
			c.L.Unlock()

			done <- true
		}()
	}

	for r := 0; r < waiterCnt+1; r++ {
		<-done
	}
}

func BenchmarkCond_1(b *testing.B) {
	benchmarkCond(b, 1)
}

func BenchmarkCond_2(b *testing.B) {
	benchmarkCond(b, 2)
}

func BenchmarkCond_4(b *testing.B) {
	benchmarkCond(b, 4)
}

func BenchmarkCond_8(b *testing.B) {
	benchmarkCond(b, 8)
}

func BenchmarkCond_16(b *testing.B) {
	benchmarkCond(b, 16)
}

func BenchmarkCond_32(b *testing.B) {
	benchmarkCond(b, 32)
}

func BenchmarkCond_64(b *testing.B) {
	benchmarkCond(b, 64)
}

func BenchmarkCond_128(b *testing.B) {
	benchmarkCond(b, 128)
}
