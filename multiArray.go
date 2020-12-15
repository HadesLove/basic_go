package main

import "fmt"

func main() {
	a := [3][4]string{{"深圳", "广州", "佛山", "东莞"},{"杭州", "宁波", "温州", "嘉兴"},{"苏州", "南京", "无锡", "南通"}}
	fmt.Println(a)
	fmt.Printf("二维数组的地址：%p\n", &a)
	fmt.Printf("二维数组的长度：%d\n", len(a))

	fmt.Printf("一维数组的长度：%d\n", len(a[0]))
	fmt.Println(a[0][3])
	fmt.Println(a[1][2])
	fmt.Println(a[2][1])

	for i := 0; i < len(a) ; i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(a[i][j], "\t")
		}
		fmt.Println()
	}

	fmt.Println("------------------")

	for _, arr := range a {
		for _, val := range arr {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}
}