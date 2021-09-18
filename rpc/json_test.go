package rpc

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStringArrayToJsonStringArray(t *testing.T) {
	dataJson := `["1","2","3"]`
	var arr []string
	_ = json.Unmarshal([]byte(dataJson), &arr)
	fmt.Println("Unmarshaled", "arr", arr)
}

func TestStringArrayToJsonIntArray(t *testing.T) {
	dataJson := `[1,2,3]`
	var arr []int64
	_ = json.Unmarshal([]byte(dataJson), &arr)
	fmt.Println("Unmarshaled", "arr", arr, len(arr))
}
