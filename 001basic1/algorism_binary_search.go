
package main

import "fmt"

//go run algorism_binary_search.go
//2分探索

func main() {
	data := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	//最後に%の文字列が出るのは何故だ。。。
	fmt.Print(binary_search(data, 90))

}

func binary_search(data []int, value int) int {
	left := 0 //探索する範囲の左端
	right := len(data) - 1

	for left <= right {
		mid := (left + right) / 2
		// fmt.Print(mid)
		if data[mid] == value {
			return mid
		} else if data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
    return -1
}


