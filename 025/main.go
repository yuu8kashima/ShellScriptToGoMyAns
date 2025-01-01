package main

import (
	"fmt"
)

/**
 * 5つの整数を入力し、それらの合計と平均を計算するプログラムを作成してください。
 * メモ: 標準入力からの読み込みを試した
 */
func main() {
	// number_count分だけ標準入力から読む
	number_count := 5
	numbers := [5]int{}
	for i := 0; i < number_count; i++ {
		fmt.Scan(&numbers[i])
	}
	// 合計と平均計算して出力
	sum := 0
	mean := float64(0)
	for _, v := range numbers {
		sum += v
		mean += float64(v)
	}
	mean /= float64(len(numbers))
	fmt.Printf("sum: %d\nmean: %f\n", sum, mean)
}
