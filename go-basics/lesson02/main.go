package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// ===== 1. if：條件不用括號，大括號一定要有 =====
	sizeBytes := 60 * 1024 * 1024
	const maxSize = 50 * 1024 * 1024

	if sizeBytes > maxSize {
		fmt.Println("檔案太大，略過")
	} else if sizeBytes == 0 {
		fmt.Println("空檔案，略過")
	} else {
		fmt.Println("大小 OK，準備處理")
	}

	// ===== 2. if 的初始化敘述：ext 只活在這個 if/else 裡 =====
	fileName := "會議紀錄.DOCX"
	if ext := strings.ToLower(filepath.Ext(fileName)); ext == ".docx" {
		fmt.Println("Word 檔，副檔名：", ext)
	} else {
		fmt.Println("不是 Word 檔，副檔名：", ext)
	}
	// fmt.Println(ext) // ← 編譯錯誤：undefined: ext

	// ===== 3. for 的四種形式 =====
	// 3-1 經典三段式
	for i := 0; i < 3; i++ {
		fmt.Println("三段式 i =", i)
	}

	// 3-2 只有條件（等於 Python 的 while）
	retry := 0
	for retry < 2 {
		retry++
		fmt.Println("while 式 retry =", retry)
	}

	// 3-3 無窮迴圈 + break / continue
	n := 0
	for {
		n++
		if n == 2 {
			continue // 跳過 2
		}
		if n > 4 {
			break
		}
		fmt.Println("無窮迴圈 n =", n)
	}

	// 3-4 range 整數（Go 1.22+），等於 Python 的 for i in range(3)
	for i := range 3 {
		fmt.Println("range 整數 i =", i)
	}

	// ===== 4. switch：不需要 break、可以多值 case =====
	name := "筆記.md" // 試著改成 a.pdf、b.txt、c.exe 再跑一次
	switch strings.ToLower(filepath.Ext(name)) {
	case ".pdf":
		fmt.Println(name, "→ PDF 解析器")
	case ".txt", ".md": // 多值 case：任一個符合就進來
		fmt.Println(name, "→ 純文字解析器")
	default:
		fmt.Println(name, "→ 不支援，略過")
	}

	// ===== 5. 無條件 switch：取代一長串 if-else if =====
	score := 0.83
	switch {
	case score >= 0.8:
		fmt.Println("高度相關")
	case score >= 0.5:
		fmt.Println("可能相關")
	default:
		fmt.Println("不相關")
	}
}
