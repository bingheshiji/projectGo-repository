package main

import "fmt"
// 有两个从小到大排好序的 int 数组（每个数组自身没有重复元素）。请找出所有在
// 这两个数组中都出现过的数。请写一个函数，输入为两个数组。
func main() {
	// 定义两个从小到大排序好的数组
	a := []int{1,2,3,4,5,6,7,8,9}
	b := []int{5,6,7,8,9,10}
	// 循环遍历是否存在
	for i:=0; i<len(a); i++ {
		for j:=0; j<len(b);j++  {
			if a[i] == b[j] {
				fmt.Println(a[i])
			} else if a[i] < b[j]  {
				break
			}
		}
	}
}