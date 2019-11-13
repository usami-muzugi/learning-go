package main

import "fmt"

func main() {
	//赋值可以使用var 变量名 类型名 = 值
	var testVal int = 1
	fmt.Println(testVal)

	testVal = 2
	fmt.Println(testVal)

	//也可以省略var 和变量名 是用 :=来赋值
	test := 1
	fmt.Println(test)

	//Go的基本数据类型

	//布尔型
	b := true
	fmt.Println(b)

	//	数字型
	//	go中数字型只有整数型int 和浮点数float32、float64
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for item := range countryCapitalMap {
		fmt.Println(item, len(item))

	}

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}
