package gc

import (
	"fmt"
	"math/rand"
	"time"
)

type test struct {
	var1 []string
	var2 []int
	var3 bool
}

func cpuhigh(randV1 int) {
	randV2 := rand.Intn(100000)
	for {
		r := (randV1 + randV2) / (randV1 + 1) * randV2
		_ = r
		if (((r*1000)/10)*randV2)%8 == 0 {
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func heapescape(id, randLength int) (*test, []int) {
	t := &test{}
	t.var1 = make([]string, 65535)
	for i := 0; i < 500+randLength; i++ {
		t.var1 = append(t.var1, "aaaa")
		t.var2 = append(t.var2, i)
		fmt.Println(id, "var1", len(t.var1), cap(t.var1))
		//time.Sleep(20 * time.Millisecond)
	}
	t.var3 = true
	n := make([]int, 1024*1024*1)
	return t, n
}

func GCMain(num int) {
	//length := 100
	/*go func() {
		for i := 0; i < 100000; i++ {
			r := rand.Intn(100000)
			go cpuhigh(r)
		}
	}()
	*/
	go func() {
		for i := 0; i < num; i++ {
			r := rand.Intn(10000)
			go heapescape(i, r)

		}
	}()
}

func sliceRetrunEscape() []int {
	s := []int{1, 2, 3}

	return s
}

//maxStackVarSize,maxImplicitStackVarSize定义在https://github.com/golang/go/blob/master/src/cmd/compile/internal/gc/go.go
func sliceSizeExceedLimitEscape() {
	// var 申明的变量，占用stack超过maxStackVarSize=int64(10*1024*1024)，将escape到heap
	go func() {
		var x [10 * 1024 * 1024]byte // no escape, as maxStackVarSize=int64(10*1024*1024)
		_ = x
		var y [10*1024*1024 + 1]byte // escape
		_ = y
		z := [10*1024*1024 + 1]byte{} // escape
		_ = z
	}()

	// make 申明的变量，占用stack超过maxImplicitStackVarSize=int64(64*1024)，65535，将escape到heap
	go func() {
		_ = make([]byte, 64*1024-1) // no escape
		_ = make([]byte, 64*1024)   // escape
		_ = make([]int, 10000)      // escape, 8*10000=80000>65535 as 64 OS, int is 8 byte.
	}()

	//sliceOverFunction(100)
}

func t1(v1 []int) {
}

func t2(v2 []int) []int {
	x := v2
	x[0] = 100
	return x
}
func t3(v3 *[]int) *[]int {
	v := v3
	return v
}
func t4(v3 []int) *[]int {
	v := v3
	return &v
}

func sliceOverFunction(len int) ([]int, [2]int, *[2]int) {
	_ = make([]int, len) // escape, variable length with slice, even if call the function with len=100

	s1 := []int{1, 2}  // escape, as return outside
	s2 := [2]int{1, 2} // no escape, as it's array and return value to outside
	s3 := [2]int{1, 2} // escape, although it's array, it return address to outside
	s4 := make([]int, 2)
	t1(s4) // no escape
	s5 := []int{1, 2}
	t2(s5) // no escape

	s6 := []int{1, 2}
	t3(&s6) // no escape

	s7 := []int{1, 2}
	t4(s7) // escape

	return s1, s2, &s3
}

func test1() {

}

func closureEscape() {

	// func escape
	func() {
		x := 100
		_ = x
	}()

	// func escape
	go func() {
		x := 100
		_ = x
	}()

	// no escape
	go test1()
}
