package mypprof

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
)

func StartNetworkProfile() {
	//runtime.GOMAXPROCS(2) // 限制 CPU 使用数，避免过载
	/*runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪
	*/
	go func() {
		fmt.Println("0.0.0.0:6060 started")
		if err := http.ListenAndServe("0.0.0.0:6060", nil); err != nil {
			panic(err)
		}
	}()

}

func StartCPUProfile() {
	runtime.SetCPUProfileRate(1000000)

	cpuprofile := `./cpu.profile`
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}

	}
}

func StartMemoryProfile() {
	memprofile := `./mem.profile`
	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

}
