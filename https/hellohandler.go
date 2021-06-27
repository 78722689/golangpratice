package https

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type helloZZHandler struct {
}

// API /hellozz
// To handle GET and POST
func (h *helloZZHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hellozz" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	switch req.Method {
	case "GET":
		if bytes, err := json.Marshal(memoryDB); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write(bytes)
		}
	case "POST":
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatalln("Received data corrupted", err)
			return
		}

		values := receiveDataBuffer.Get().([]string)
		// Return memory back
		defer receiveDataBuffer.Put(values)

		if err := json.Unmarshal(reqBody, &values); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalln("Decode data occured failure ", err)
			return
		}

		log.Println(values)
		if len(values) != 0 {
			result := h.processData(values)
			// Return memory back
			//defer sendDataBuffer.Put(result)

			if r, err := json.Marshal(result); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			} else {
				if _, err := w.Write(r); err != nil {
					log.Println(err)
					return
				}
			}

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}

func (h *helloZZHandler) existedInDB(value string) bool {
	for _, v := range memoryDB {
		if v == value {
			return true
		}
	}

	return false
}

// DO NOT change data in the function
func (h *helloZZHandler) processData(data []string) []bool {
	//result := sendDataBuffer.Get().([]bool) //make([]bool, len(data))
	result := make([]bool, len(data))

	for i, v := range data {
		duplicated := func() bool {
			for _, sv := range data[0:(i - 1)] {
				if v == sv {
					return true
				}
			}
			return false
		}

		if i > 0 && duplicated() {
			result[i] = true
			continue
		}

		mutexDB.Lock()
		if h.existedInDB(v) {
			result[i] = true
		} else {
			memoryDB = append(memoryDB, v)
			result[i] = false
		}
		mutexDB.Unlock()
	}

	return result
}
