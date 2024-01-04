package covid
import (
    "testing"
)
// 單元測試
func TestQueryCovidCase(t *testing.T) {
    regions := []string{"US", "Japan", "Taiwan*"}
    for _, region := range regions {
        result, err := QueryCovidCase(region)
        if err != nil {
            // err 不為 nil 時印出訊息並回報測試失敗
            t.Errorf("error: %v", err)
        }
        t.Logf("result for %v = %v\n", region, result)
    }
}