package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func ChannelMain() {
	//readClosedChannel()
	//writeClosedChannel()

	//closeClosedChannel()

	//readNilChannel()
	//writeNilChannel()
	checkClosedChannel()
}

func readClosedChannel() {
	c := make(chan int, 10)
	c <- rand.Intn(100)
	c <- rand.Intn(100)
	close(c)

	/*for i := 10; i < 10; i++ {
		// 如果closed并且没有数据接收， ok=false. 此方法可以判断channel是否已经关闭
		// 只要有数据，即使已经close，ok也为true
		if v, ok := <-c; ok {
			fmt.Println(v)
		}
	}*/

	// 虽然已经closed，仍然可以读取出里面已经存入数据
	for i := range c {
		fmt.Println(i)
	}
}

func writeClosedChannel() {
	c := make(chan int)

	go func() {
		for {
			select {
			case <-time.Tick(1 * time.Second):
				c <- rand.Intn(100)
			}
		}
	}()

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println("Received", v)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(c)
	time.Sleep(5 * time.Second)
}

func closeClosedChannel() {
	c := make(chan int)

	close(c)
	close(c)
}

func readNilChannel() {
	c := make(chan int)
	c = nil

	fmt.Println("Reading...")
	fmt.Println(<-c)
	fmt.Println("Exit...")
}

func writeNilChannel() {
	c := make(chan int)
	c = nil

	fmt.Println("Writing...")
	c <- 100
	fmt.Println("Exit...")
}

func checkClosedChannel() {
	c := make(chan int, 10)

	/*go func() {

		fmt.Println("closing")
		close(c)
		fmt.Println("closed")
	}()*/

	go func() {
		closed := false
		for {
			if v, ok := <-c; ok {
				fmt.Println("ok", v)
			} else {
				fmt.Println("nok", v)
			}
			time.Sleep(1 * time.Second)
			if !closed {
				close(c)
				fmt.Println("closed")
				closed = true
			}
		}
	}()

	c <- 5
	//fmt.Println("set 5")
	//time.Sleep(1 * time.Second)
	c <- 8
	c <- 10
	c <- 2
	c <- 3
	c <- 5
	c <- 88
	c <- 99
	time.Sleep(11 * time.Second)
}
