package main

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

func main() {
	// ===== 1. slice：可以長大的清單（像 Python 的 list）=====
	var empty []string                  // 零值是 nil，長度 0，可以直接 append
	files := []string{"a.pdf", "b.txt"} // 直接給初始值
	fmt.Println(empty == nil, len(empty), files, len(files))

	// append 會「回傳」新的 slice，一定要接回原變數
	// Python：files.append("c.md")   ← 原地修改
	// Go    ：files = append(files, "c.md")
	files = append(files, "c.md")
	files = append(files, "d.docx", "e.png") // 一次加多個
	fmt.Println("files =", files, "len =", len(files))

	// 已知大概數量時，先用 make 預留容量，append 時比較不用一直搬家
	supported := make([]string, 0, len(files)) // 長度 0、容量 len(files)
	fmt.Println("supported len/cap =", len(supported), cap(supported))

	// ===== 2. range slice：同時拿到 index 和值 =====
	// Python：for i, f in enumerate(files):
	for i, f := range files {
		fmt.Println(i, f)
	}
	// 不需要 index 時用 _ 丟掉（宣告了不用會編譯失敗，_ 是例外）
	// Python：for f in files:
	for _, f := range files {
		ext := strings.ToLower(filepath.Ext(f))
		if ext == ".pdf" || ext == ".txt" || ext == ".md" || ext == ".docx" {
			supported = append(supported, f)
		}
	}
	fmt.Println("supported =", supported)

	// ===== 3. 切片 s[a:b]：和 Python 一樣含頭不含尾 =====
	fmt.Println(files[1:3], files[:2], files[3:])
	// ⚠️ 和 Python 最大的不同：Go 的切片「不會複製」，和原本的 slice 共用同一塊記憶體
	part := files[:2]
	part[0] = "改掉了.pdf"
	fmt.Println("files[0] =", files[0]) // 原本的 files 也被改了！
	// 需要獨立的副本時用 slices.Clone（Python 的 files[:2] 本來就是副本）
	safe := slices.Clone(files[:2])
	safe[0] = "只改副本.pdf"
	fmt.Println("files[0] =", files[0], "safe[0] =", safe[0])

	// ===== 4. 實戰：把文字切成固定大小的 chunk =====
	words := strings.Fields("地端 AI 文件 搜尋 系統 把 文件 切成 小塊 再 做 向量化")
	const chunkSize = 5
	for start := 0; start < len(words); start += chunkSize {
		end := min(start+chunkSize, len(words)) // min 是內建函式（Go 1.21+），避免最後一塊超出範圍
		fmt.Printf("chunk %d：%v\n", start/chunkSize, words[start:end])
	}

	// ===== 5. map：key → value 對照表（像 Python 的 dict）=====
	hashes := map[string]string{
		"docs/a.pdf": "aaa111",
		"docs/b.txt": "bbb222",
	}
	hashes["notes/c.md"] = "ccc333" // 新增或覆蓋
	fmt.Println("len(hashes) =", len(hashes))

	// 讀不存在的 key 不會報錯，而是回傳「零值」（string 的零值是 ""）
	// Python 的 d["x"] 會 KeyError，比較像 d.get("x", "")
	fmt.Printf("不存在的 key：%q\n", hashes["nope.pdf"])

	// comma-ok：分辨「真的不存在」還是「存在但值剛好是零值」
	// Python：if "docs/a.pdf" in hashes:
	if h, ok := hashes["docs/a.pdf"]; ok {
		fmt.Println("找到了，hash =", h)
	}
	if _, ok := hashes["nope.pdf"]; !ok {
		fmt.Println("nope.pdf 不在 map 裡")
	}

	delete(hashes, "docs/b.txt") // Python：del hashes["docs/b.txt"]（key 不存在也不會報錯）
	fmt.Println("刪除後 len =", len(hashes))

	// ⚠️ nil map 可以讀，但「寫入會 panic」→ 要用 make 或字面值建立
	// var bad map[string]int
	// bad["x"] = 1 // panic: assignment to entry in nil map
	counts := make(map[string]int)

	// ===== 6. 用 map 計數：int 零值是 0，所以可以直接 ++ =====
	// Python：counts[ext] = counts.get(ext, 0) + 1
	for _, f := range []string{"a.PDF", "b.pdf", "c.md", "d.txt", "e.md", "f.pdf"} {
		counts[strings.ToLower(filepath.Ext(f))]++
	}

	// ===== 7. range map：順序是隨機的！每次跑都可能不同 =====
	for ext, n := range counts {
		fmt.Println("（隨機順序）", ext, n)
	}
	// 需要固定順序時：先把 key 收集起來排序，再依序讀
	keys := make([]string, 0, len(counts))
	for ext := range counts { // 只要 key 時可以省略 value
		keys = append(keys, ext)
	}
	slices.Sort(keys)
	for _, ext := range keys {
		fmt.Printf("（排序後）%s：%d 個\n", ext, counts[ext])
	}
}
