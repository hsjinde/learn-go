package main

import "fmt"

func main() {
	fileName := "報告.pdf"
	sizeBytes := 2048000
	var pageCount int
	fmt.Printf("pageCount 的零值：%d\n", pageCount)
	const bytesPerMB = 1024 * 1024
	sizeMB := float64(sizeBytes) / bytesPerMB
	fmt.Printf("檔案：%s，大小：%.2f MB\n", fileName, sizeMB)
}
