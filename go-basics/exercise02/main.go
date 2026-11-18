// 練習題 02：模擬處理「一個檔案」的流程
//
// 1. 宣告 fileName（例如 "規格書.PDF"）和 sizeBytes，並用 const 定義 50 MB 上限
// 2. 大小判斷：超過上限、或大小為 0 → 印出原因並結束（提示：main 裡的 return）
// 3. 依副檔名分派（用 switch，大小寫都要能判斷）：
//    .pdf        → 使用 PDF 解析器
//    .docx       → 使用 Word 解析器
//    .txt / .md  → 使用純文字解析器
//    其他        → 不支援，結束
// 4. 失敗重試：const maxRetries = 3，再用常數 succeedOn 模擬「第幾次會成功」
//    每次印「第 N 次嘗試…」，成功印「成功」並跳出；3 次都失敗印「放棄」
//    把 succeedOn 分別改成 2 和 5 各跑一次，確認輸出正確
//
// 寫完在 go-basics 底下執行：
//    go run ./exercise02
//    go vet ./...
//    gofmt -l .

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// TODO: 從這裡開始寫
	// 提示：用到 fmt、path/filepath、strings 時，記得加到 import
	fileName := "test.exe"
	fmt.Println(fileName)
	const sizeBytes = 40 * 1024 * 1024
	const maxSize = 50 * 1024 * 1024
	if sizeBytes > maxSize {
		fmt.Println("檔案過大")
		return
	} else if sizeBytes == 0 {
		fmt.Println("檔案為空")
		return
	} else {
		fmt.Println("處理中...")
	}

	//name := "test.exe"
	switch strings.ToLower(filepath.Ext(fileName)) {
	case ".pdf":
		fmt.Println(fileName, "→ PDF 解析器")
	case ".docx":
		fmt.Println(fileName, "→ 使用 Word 解析器")
	case ".txt", ".md": // 多值 case：任一個符合就進來
		fmt.Println(fileName, "→ 純文字解析器")
	default:
		fmt.Println(fileName, "→ 不支援，略過")
		return
	}

	const maxRetries = 3
	const succeedOn = 2
	retry := 0
	for retry < maxRetries {
		retry++
		fmt.Printf("第 %d 次嘗試…\n", retry)
		if retry == succeedOn {
			fmt.Println("成功")
			break
		}
		if retry == maxRetries {
			fmt.Println("放棄")
			break
		}
	}

}
