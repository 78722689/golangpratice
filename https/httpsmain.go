package https

import (
	"flag"
	"log"
	"os"
	"strings"
)

func HTTPSMain() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Lmicroseconds | log.LstdFlags)

	var start string
	var send string
	cmd := flag.NewFlagSet("https", flag.ExitOnError)
	cmd.StringVar(&start, "start", "server", "Start https server or https client. Options server/client, default server")
	cmd.StringVar(&send, "send", "", "Send data, values seperate with \",\", such as \"aaa,111,ccc,222,333\" ")
	cmd.Parse(os.Args[2:])

	value := strings.ToLower(start)
	switch value {
	case "server":
		smain(":8443")
	case "client":
		s := strings.Split(send, ",")
		if len(s) == 0 {
			log.Fatal("Not valid data")
		}

		if result, err := NewClient(s); err != nil {
			log.Fatal(err)
		} else {
			log.Println(result)
		}

	default:
		log.Fatal("Unsupported start, options are server and client.")
	}
}
