/**
*   编写一个小程序
*   给定一个字符串数组
*   ["I","am","stupid","and","weak"]
*   用 for 循环遍历该数组并修改为
*   ["I","am","smart","and","strong"]
 */

package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"I", "am", "stupid", "and", "weak"}
	fmt.Printf("mySlice %+v\n", mySlice)
	for index := range mySlice {
		if mySlice[index] == "stupid" {
			mySlice[index] = "smart"
		}
		if mySlice[index] == "weak" {
			mySlice[index] = "strong"
		}
	}
	fmt.Printf("mySlice %+v\n", mySlice)
}
