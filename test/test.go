package main

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

type A struct{}
type B struct {
	A // 只写出类型, 没有变量名 -- > 匿名字段
}

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	// 一般情况下，定义结构体的时候，字段名和类型一一对应，如下:Person为类型，并无字段名。
	Person //匿名字段,name默认Student就包含了Person的所有字段。
	id     int
	addr   string
}

func (p *A) Show() {
	fmt.Println("start ...")
	p.Show2()
}
func (p *A) Show2() {
	fmt.Println("AAA")
}

func (p *B) Show2() { // 名称也是Show2 (属于结构体函数, 可以这样写, 不会报错)
	fmt.Println("BBB")
}

func testObject() {
	b := &B{}
	b.Show() // 相当于b先访问了自己的匿名字段A, A再调用Show()

	//顺序初始化
	s1 := Student{Person{"ck_god", 1, 18}, 1, "sz"}
	fmt.Printf("s1=%+v\n", s1)

	//部分成员初始化1
	s2 := Student{Person: Person{"xiaobai", 0, 20}, id: 2, addr: "sz"}

	//部分成员初始化2
	s3 := Student{Person: Person{name: "kavai"}, id: 3}
	fmt.Println(s2, s3)

	//访问成员变量
	var s4 Student
	s4.name = "ck_god"
	s4.sex = 1
	s4.age = 18
	s4.id = 1
	s4.addr = "sz"
	fmt.Println(s4)

	//访问成员变量2
	var s5 Student
	s5.Person = Person{"god_girl", 1, 23}
	s5.id = 2
	s5.addr = "wz"
	fmt.Println(s5)
}

func main() {
	testJson()
}

func testJson() {
	//拼凑json body为map数组
	var rbody []map[string]interface{}
	t := make(map[string]interface{})
	t["device_id"] = "dddddd"
	t["device_hid"] = "ddddddd"
	rbody = append(rbody, t)
	t1 := make(map[string]interface{})
	t1["device_id"] = "aaaaa"
	t1["device_hid"] = "aaaaa"
	rbody = append(rbody, t1)
	cnnJson := make(map[string]interface{})
	cnnJson["code_msg"] = ""
	cnnJson["body"] = rbody
	cnnJson["page"] = 0
	cnnJson["page_size"] = 0
	b, _ := json.Marshal(cnnJson)
	cnnn := string(b)
	fmt.Println("cnnn:%s", cnnn)
	cn_json, _ := simplejson.NewJson([]byte(cnnn))
	cn_body, _ := cn_json.Get("body").Array()
	for _, di := range cn_body { //就在这里对di进行类型判断
		newdi, _ := di.(map[string]interface{})
		device_id := newdi["device_id"]
		device_hid := newdi["device_hid"]
		fmt.Println(device_hid, device_id)
	}
}
