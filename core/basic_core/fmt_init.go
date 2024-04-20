package basic_core

import "fmt"

type FmtInit struct{}

func (f FmtInit) BasicPrint() {
	test1()
	test2()
	test3()
	test4()

}

// 打印
func test1() {
	fmt.Println("Basic Print")
	name := "John Doe"
	fmt.Printf("Hello %s!\n", name)
	fmt.Println("在终端打印单独的一行")
}

// 传入的数据生成并返回要给字符串
func test2() {
	s1 := fmt.Sprint("枯藤")
	name := "John Doe"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	fmt.Println(s1, s2, "over")
}
func test3() {
	err := fmt.Errorf("这是一个错误%s", " nihaoa ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("test3")
}
func test4() {
	//var (
	//	name    string
	//	age     int
	//	married bool
	//)
	////就是让你输入的东西
	//fmt.Scan(&name, &age, &married)
	//fmt.Printf("扫描结果 name：%s age：%d married：%t\n", name, age, married)

}
