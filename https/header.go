package https

import "sync"

var (
	CLIENT_CERT_FILE = "./cert/client.crt"
	CLIENT_KEY_FILE  = "./cert/client.key"
	ROOT_CA_FILE     = "./cert/rootca.crt"
	SERVER_KEY_FILE  = "./cert/localhost.key"
	SERVER_CA_FILE   = "./cert/localhost.crt"

	DEFAULT_BUFFER_SIZE = 32
	// To reduce memory usage and overcome GC time
	receiveDataBuffer = sync.Pool{
		New: func() interface{} {
			return make([]string, DEFAULT_BUFFER_SIZE)
		},
	}
	// TODO
	/*sendDataBuffer = sync.Pool{
		New: func() interface{} {
			return make([]bool, DEFAULT_BUFFER_SIZE)
		},
	}*/

	// To persist the proceeded data which are sent from client
	// Never released except application reboot
	// Shared by all goroutines, lock&unlock when using.
	memoryDB = make([]string, 0, 0)
	mutexDB  sync.Mutex
)
