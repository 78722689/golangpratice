package mysync

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var l sync.Mutex

func r() {

	defer l.Unlock()
	l.Lock()

	d := time.Duration(rand.Intn(5))
	time.Sleep(d * time.Second)
	fmt.Println("Sleep", d)
}

func TestSync() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < 10; i++ {
		go r()
	}
	time.Sleep(11 * time.Second)
}
