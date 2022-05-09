package lock

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var value1, value2, value3 int32
var lock sync.Mutex

func IncreValue1() {
	value1++
}
func IncreValue2() {
	atomic.AddInt32(&value2, 1)
}
func IncreValue3() {
	lock.Lock()
	value3++
	lock.Unlock()
}

func TestLocking(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go IncreValue1()
		go IncreValue2()
		go IncreValue3()
	}

	time.Sleep(1000)
	fmt.Println("thread unsafe: ", value1)
	fmt.Println("optimistic lock: ", value2)
	fmt.Println("pesimistic lock: ", value3)

}
