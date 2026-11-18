# Go 基礎教學計畫

這個資料夾是「先學 Go 基礎、再做 nas-doc-search 專案」的學習區。
目前階段：**學 Go 基礎，不做專案**。等 11 課都完成再開始專案階段 1。

## 學習者背景

- 會 Python，Go 是初學者
- 開發環境：Windows 11、PowerShell、VS Code（Go 擴充套件），Go 版本 1.27.4
- 部署環境：Linux VPS（Docker）
- 最終目標：完成 nas-doc-search（地端 AI 文件搜尋），作為應徵群暉科技「Product Developer（生成式 AI 應用開發）」的作品

## 教學方式（每次都要遵守）

1. **一次只教一課**，不要預先產生後面的課程
2. **先解釋為什麼這樣設計**，再給程式碼
3. 程式碼加**中文註解**；遇到 Go 特有觀念時，**用 Python 寫法對照**
4. 每課都要說明這個觀念在 nas-doc-search 專案的**哪裡會用到**
5. 範例程式由學習者**自己動手輸入**；Claude 先在 scratchpad 驗證範例能編譯、能執行，不直接寫進學習者的資料夾
6. 每課最後留一題**練習題**，讓學習者自己寫，寫完貼回來 review
7. Review 時要實際執行 `go run`、`go vet ./...`、`gofmt -l .`，並指出：
   - 有沒有符合題目要求
   - 不符合 Go 慣例的地方（命名、多餘的轉型、格式）
   - 先肯定做對的地方，再列出問題；問題依嚴重程度標示 ❌ / ⚠️ / 💡
8. 指令一律給 **PowerShell 版本**（curl 用 `curl.exe`），路徑含空格要加引號
9. 不確定的套件版本或 API，提醒學習者查官方文件（https://go.dev/doc）
10. 回覆一律使用繁體中文（台灣用語），程式碼與識別字維持原文

## 資料夾結構

```
D:\learn go\
  CLAUDE.md
  go-basics\            ← 課程模組（module 名稱：gobasics）
    go.mod
    lesson01\main.go    ← 課堂範例
    exercise01\main.go  ← 練習題（獨立資料夾、獨立的 func main）
    lesson02\main.go
    exercise02\main.go
    ...
```

- 一個資料夾 = 一個 package；`package main` 的資料夾裡只能有一個 `func main()`
- 執行方式：在 `go-basics` 底下執行 `go run ./lessonNN` 或 `go run ./exerciseNN`

## 課程路線

| 課 | 主題 | 對應專案用途 | 狀態 |
|---|---|---|---|
| 01 | 程式結構、變數、型別、零值、常數、fmt | 一切的基礎 | ✅ 完成 |
| 02 | if / for / switch | 依副檔名分派、切 chunk、失敗重試 | ✅ 完成 |
| 03 | slice、map、range | 存檔案清單、hash 對照表 | ⏳ 進行中（等練習題） |
| 04 | 函式、多回傳值、error 處理 | 「錯誤一律處理」的規範 | 未開始 |
| 05 | struct、method、pointer | 定義 `File{Path, Size, ModTime}` | 未開始 |
| 06 | interface | `Extractor` 介面設計 | 未開始 |
| 07 | package、大小寫匯出、專案結構 | `internal/store`、`internal/watcher` 拆分 | 未開始 |
| 08 | 檔案操作、`filepath.WalkDir`、`filepath.Join` | 階段 1 的資料夾掃描 | 未開始 |
| 09 | goroutine、channel、`sync.WaitGroup` | 限制 worker 數量的任務佇列 | 未開始 |
| 10 | `net/http`、JSON | `GET /files` API | 未開始 |
| 11 | testing | 評估工具、程式品質 | 未開始 |

每完成一課，更新上表的「狀態」欄，並在下方「學習紀錄」補充重點。

## 學習紀錄

### 第 1 課
- 已學：`package main`、`func main`、`var` / `:=`、零值、`const`、`Println` / `Printf`
- 作業第 2 次 review（2026-11-13）：
  - ✅ 已修：`%.2f`、拿掉 `float64(bytesPerMB)` 多餘轉型 → 計算結果 1.95 MB 正確
  - 學習者自己完成：搬到 `exercise01\main.go`、`pageCount` 加上說明文字
  - Claude 協助收尾（2026-11-13）：`printBanner` 改為 `func main`、刪除過時註解與 lesson01 裡的 `printBanner()` 呼叫、`mian.go` 改名、刪除多餘的根目錄 `go.mod`、`go fmt`
  - 教訓：搬檔案後要一併檢查「呼叫端」和「被呼叫端」，搬完先 `go vet ./...` 再交作業
- 已加上 `.gitattributes`（`*.go` 一律 LF），之後 gofmt 不會再因 CRLF 報錯

## Git

- Repo：https://github.com/hsjinde/learn-go （分支 `main`）
- 每完成一課（範例 + 練習題通過 review）就 commit + push 一次

### 第 2 課
- 已教：if（含初始化敘述與作用範圍）、for 的四種形式、`range` 整數（Go 1.22+）、switch（不需要 break、多值 case、無條件 switch）
- 練習題：`exercise02` 模擬「大小判斷 → 依副檔名分派 → 失敗重試 3 次」
- 作業 review（2026-11-18，第 4 次通過）：
  - 卡過的點：`const x := ...`（const 只能用 `=`）、用 `or` 而不是 `||`、同一範圍重複 `:=`、把沒用到的 import 註解掉而不是去用它
  - 卡最久的點：`main` 裡的 `return` = 程式結束；沒 return 的話判斷完還會往下跑
  - 曾直接複製 lesson 範例的變數（`name`），導致判斷的不是同一個檔案 → 提醒「複製後要對齊自己的變數」
  - 學習者表示看不懂條列式 review 時，改用「第幾行 → 改成什麼 → 預期輸出」的步驟式說明比較有效

## 常見易錯點（Python 背景）

- 宣告了沒用到的變數、import 了沒用到的套件，都會**編譯失敗**
- 整數除以整數還是整數（像 Python 的 `//`），要先 `float64()` 轉型
- `:=` 只能在函式內使用，而且同一個作用範圍內不能重複宣告同一個名稱
- Windows 換行（CRLF）會讓 `gofmt -l` 把檔案列出來

## 完成基礎後

回到 nas-doc-search 專案的階段 1：掃描資料夾存 SQLite（`modernc.org/sqlite`）、fsnotify 監控、`GET /files` API、Dockerfile。
