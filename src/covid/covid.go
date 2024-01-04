package covid
import (
    "encoding/csv"
    "errors"
    "io"
    "net/http"
    "strconv"
    "strings"
)
// COVID-19 全球病例線上報表路徑
const url = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv"
// QueryCovidCase 會搜尋並傳回指定國家的病例數
func QueryCovidCase(region string) (int, error) {
    // 用 HTTP GET 取回 CSV 檔內容
    r, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer r.Body.Close()
    reader := csv.NewReader(r.Body)
    header := true
    // 走訪 CSV 檔, 找到相符國家名稱就傳回病例數
    // 有錯誤或找不到則傳回 0
    for {
        record, rowErr := reader.Read()
        if rowErr == io.EOF {
            break
        } else if rowErr != nil {
            return 0, rowErr
        }
        // 跳過 CSV 標頭
        if header {
            header = false
            continue
        }
        // 用小寫國名比對, 並跳過個別省分的資料
        if strings.ToLower(record[1]) == strings.ToLower(region) &&
            record[0] == "" {
            cases, convErr := strconv.Atoi(record[len(record)-1])
            if convErr != nil {
                return 0, convErr
            }
            return cases, nil
        }
    }
    return 0, errors.New("查無該國名稱")
}