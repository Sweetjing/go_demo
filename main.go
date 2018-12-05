package main

import "fmt"

/*
1.（1）
	《1》go语言是以包为管理单位的
	《2》每个文件必须声明一个生命包
	《3》每一个程序必须要有一个main包
（2）go只有一个入口函数main
（3）PrintIn()函数会自动换行
	调用函数大多数需要导入包
（4）go语言结尾没有分号
2.类型转换
（1）bool类型不能转换为整形《===》整形也不能转化为bool类型：这种类型是不可转换的类型，叫做不兼容类型
	0就是假，非0就是真
3.
 */
func main()  {

	var flag bool
	flag = true
	fmt.Printf("%t",flag)   //打印true



	
}
