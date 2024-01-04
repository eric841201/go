package main
import (
    "fmt"
    "syscall/js" // WASM 函式庫
    "wasm/src/covid" // covid 套件
)
// 傳回 queryCovidCase 的 JavaScript 版函式並處理 Go 語言錯誤
func jsFuncWrapper() js.Func {
    // 傳回 JavaScript 函式
     return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        // 取得 JavaScript DOM 文件元素
        alert := js.Global().Get("alert")
        doc := js.Global().Get("document")
        label := doc.Call("getElementById", "result")
        if !label.Truthy() {
            alert.Invoke("網頁未包含 id='result' 元素")
            return nil
        }
        label.Set("innerHTML", "")
        // 限制呼叫函式的引數數量
        if len(args) != 1 {
            alert.Invoke("WASM 函式引數數量錯誤")
            return nil
        }
        // 開一個新的 Goroutine, 以免 http.Get 卡死 js.FuncOf
        go func() {
            // 呼叫 queryCovidCase
            result, err := covid.QueryCovidCase(args[0].String())
            if err != nil {
                alert.Invoke("WASM 執行錯誤: " + err.Error())
                return
            }
            // 將查詢結果寫到網頁的 DOM 元素
            label.Set("innerHTML", fmt.Sprintf("案例數: %v", result))
        }()
        return nil
    })
}
func main() {
// 註冊 JavaScript 函式, 結束時釋出資源
    jsFunc := jsFuncWrapper()
    js.Global().Set("queryCovidCase", jsFunc)
    defer jsFunc.Release()
    // 用空 select 卡住主程式
    select {}
}