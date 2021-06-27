package main

import (
	"golangpratice/channel"
	"golangpratice/gc"
	"golangpratice/kafka"
	"golangpratice/mycontext"
	"golangpratice/mysync"
	"golangpratice/rpc"
	"golangpratice/tip"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(2)
	// 设置CPU采样率，默认是100, 设置低了可能采集不到CPU数据，推荐设置高一点。
	runtime.SetCPUProfileRate(500)
	log.SetFlags(log.Lshortfile | log.Ltime | log.Lmicroseconds | log.LstdFlags)
	/*
		cpuprofile := `./cpu.profile`
		//if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			panic(err)
		}

		// f千万不能在局部被关闭，否则采样数据定不进profile
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}

		//}
		defer pprof.StopCPUProfile()
	*/
	//mypprof.StartNetworkProfile()

	time.Sleep(1 * time.Second)

	if len(os.Args) < 2 ||
		os.Args[1] == "-help" ||
		os.Args[1] == "-h" {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "kafka":
		arg := &kafka.Arguments{}
		arg.ParseKafkaArg()
		kafka.KafkaMain(arg)
	case "context":
		mycontext.ContextMain()
	case "channel":
		channel.ChannelMain()
	case "tip":
		tip.TipMain()
	case "sync":
		mysync.TestSync()
	case "gc":
		gc.GCMain(10000)
	case "rpc":
		if len(os.Args) < 3 {
			log.Fatal("Arguments are not enough")
		}
		tls := false
		if len(os.Args) >= 4 && os.Args[3] == "tls" {
			tls = true
		}
		if os.Args[2] == "server" {
			rpc.RPCServerMain(tls)
		} else {
			rpc.RPCClientMain(os.Args[4], tls)
		}
	default:
		usage()
	}

	/*memprofile := `./mem.profile`
	//if memprofile != "" {
	f, err = os.Create(memprofile)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	*/
	//}

	time.Sleep(10 * time.Second)
	/*runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}*/
	//time.Sleep(3 * time.Second)
}
