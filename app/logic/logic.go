package logic

import (
	"encoding/json"
	"fmt"
)

var PointTab int = 0

type Gamebob struct {
	Count int `json:"points"`
}

func Point(arr []byte) {
	var g Gamebob
	err := json.Unmarshal(arr, &g)
	if err != nil {
		fmt.Println("Erorr:", err)
		return
	}

	count := g.Count
	PointTab += count
	fmt.Println(PointTab)
}
func GetPointTab() int {
	return PointTab
}
