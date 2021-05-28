package tip

import (
	"encoding/json"
	"fmt"
)

func TipMain() {
	//show1()
	//show2(student{})
	//jsonTest()
	//maptest()
	slicetest()
}

type Param map[string]interface{}

type Show struct {
	Param
}

func show1() {
	s := new(Show)

	// 如果没有下面这行给map申请内存，对Param赋值将会Panic
	//s.Param = make(map[string]interface{})
	s.Param["RMB"] = 20000

	fmt.Println(s)
}

type student struct {
	Name string
}

func show2(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		fmt.Println(1111, msg)
	}
}

func jsonTest() {
	// 使用json.Unmarshal，struct的字段必须为导出字段，否则读取不到值
	type People struct {
		name string `json:"Name"`
	}

	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p.name)
}

type Student struct {
	Name string
}

func maptest() {
	m := make(map[string]Student)
	m["1"] = Student{"abc"}
	m["2"] = Student{"xyz"}
	// m["1"].Name = "111"  // 这句会导致编译错误，因为无法取到Name的地址
	fmt.Println(m)
}

func slicetest() {
	m := make([]int, 10000)
	for i := 0; i < 20; i++ {
		m = append(m, i)
		fmt.Printf("%p=%d\n", m, m[i])
	}

}
