// 每個 .go 檔第一行都要宣告自己屬於哪個 package。
// package main 很特別：代表「這是一支可以直接執行的程式」。
package main

// import 引入標準庫。注意：import 了卻沒用到，會編譯失敗！
import "fmt"

// func main() 是程式進入點，相當於 Python 的 if __name__ == "__main__":
func main() {
	// 方式一：var 名稱 型別 = 值（完整寫法）
	var projectName string = "nas-doc-search"

	// 方式二：:= 短宣告，Go 自動推斷型別（最常用，但只能在函式內使用）
	// Python 對照：file_count = 42
	fileCount := 42
	sizeMB := 3.5 // 推斷為 float64

	// 方式三：只宣告不給值 → 自動得到「零值」
	// Python 沒有這個概念（Python 要先 x = None）
	var status string // 零值是 ""
	var retries int   // 零值是 0
	var done bool     // 零值是 false

	// 常數：編譯時就決定，之後不能改
	const defaultModel = "qwen3.5:4b"

	// fmt.Println 類似 Python 的 print()
	fmt.Println("專案：", projectName)

	// fmt.Printf 類似 Python 的 f-string / % 格式化
	// %d 整數、%.1f 小數一位、%s 字串、%q 加引號的字串、%v 任何值、%T 型別
	fmt.Printf("檔案數：%d，大小：%.1f MB\n", fileCount, sizeMB)
	fmt.Printf("零值 → status=%q retries=%d done=%v\n", status, retries, done)
	fmt.Printf("sizeMB 的型別是 %T\n", sizeMB)
	fmt.Println("預設模型：", defaultModel)
}
