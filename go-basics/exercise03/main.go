// 練習題 03：模擬「重新掃描資料夾後，比對哪些檔案變了」
//
// 下面已經準備好三份資料：
//   scanned   ：這次掃描到的所有檔案路徑
//   oldHashes ：上次掃描時存下來的「路徑 → hash」
//   newHashes ：這次掃描算出來的「路徑 → hash」（只有支援的格式才會算 hash）
//
// 1. 從 scanned 篩出支援的檔案（.pdf / .docx / .txt / .md，大小寫都要能判斷），
//    append 到新的 slice supported，印出「支援 N / 總共 M」
// 2. 用 map[string]int 統計 supported 裡每種副檔名（小寫）有幾個，
//    依副檔名排序後印出（提示：map 的 range 順序是隨機的）
// 3. 比對 newHashes 和 oldHashes，分別收集到三個 slice，排序後印出：
//    新增：在 new、不在 old
//    修改：兩邊都有，但 hash 不同
//    刪除：在 old、不在 new
//    （提示：comma-ok）
//
// 預期輸出：
//    支援 5 / 總共 7
//    .docx：1 個
//    .md：1 個
//    .pdf：2 個
//    .txt：1 個
//    新增： [docs/報價單.pdf notes/readme.txt]
//    修改： [docs/會議紀錄.docx]
//    刪除： [docs/已刪除.pdf]
//
// 寫完在 go-basics 底下執行：
//    go run ./exercise03
//    go vet ./...
//    gofmt -l .

package main

func main() {
	scanned := []string{
		"docs/規格書.PDF",
		"docs/會議紀錄.docx",
		"notes/待辦.md",
		"img/logo.png",
		"docs/報價單.pdf",
		"notes/readme.txt",
		"backup/old.zip",
	}
	oldHashes := map[string]string{
		"docs/規格書.PDF":   "h1",
		"docs/會議紀錄.docx": "h2",
		"notes/待辦.md":    "h3",
		"docs/已刪除.pdf":   "h9",
	}
	newHashes := map[string]string{
		"docs/規格書.PDF":     "h1",
		"docs/會議紀錄.docx":   "h2-v2",
		"notes/待辦.md":      "h3",
		"docs/報價單.pdf":     "h4",
		"notes/readme.txt": "h5",
	}

	// 這行只是讓骨架能編譯（宣告了沒用到會報錯）；開始寫之後就刪掉
	_, _, _ = scanned, oldHashes, newHashes

	// TODO: 從這裡開始寫
	// 提示：用到 fmt、path/filepath、strings、slices 時，記得加到 import
}
