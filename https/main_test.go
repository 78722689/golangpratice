package https

import (
	"testing"
	"time"
)

var (
	input = [][]string{
		[]string{"111", "222", "333", "444", "555"},
		[]string{"888", "999", "333", "000", "666"},
		[]string{"111", "222", "111", "666", "555", "77"},
		[]string{"9"},
		[]string{"888"},
	}

	expect = [][]bool{
		[]bool{false, false, false, false, false},
		[]bool{false, false, true, false, false},
		[]bool{true, true, true, true, true, false},
		[]bool{false},
		[]bool{true},
	}

	API_HELLOZZ = "/hellozz"

	ADDRESS = "localhost:8443"

	https_client *TLSClient
)

func startSendInput(t *testing.T) {
	var (
		err error
		//tls  *TLSClient
		resp []bool
		fail bool
	)

	fail = false
	// Send row by row of input data
	for i, iv := range input {
		if resp, err = NewClient(iv); err != nil {
			t.Log("Send data failed.", err)
			t.FailNow()
		}
		if resp == nil || len(resp) != len(expect[i]) {
			t.Errorf("input[%d].resp!=expect[%d], %v!=%v", i, i, resp, expect[i])
			continue
		}

		// check resp with a corresponding row of expect
		equal := func() bool {
			for j, ev := range expect[i] {
				if ev != resp[j] {
					return false
				}
			}
			return true
		}

		if equal() {
			t.Logf("input[%d].resp==expect[%d], %v==%v", i, i, resp, expect[i])
		} else {
			t.Errorf("input[%d].resp!=expect[%d], %v!=%v", i, i, resp, expect[i])
			fail = true
		}
	}

	// Any response of row of input is not equal with expect, then case failed.
	if fail {
		t.Fail()
	}
}

// Test with input and expect
func Test_validate_api_hello(t *testing.T) {
	// Start https server
	// Any failures in the smain, then case fail
	go smain(ADDRESS)
	time.Sleep(2 * time.Second)

	// Begin to send data and check resp if it's same as expecting
	startSendInput(t)
}

// TODO
func Benchmark_api(b *testing.B) {

}
