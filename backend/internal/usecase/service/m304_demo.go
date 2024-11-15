package service

import (
	"fmt"
	"math"
	"net/http"

	m304 "github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/M304"
)

// m304にはあらかじめ値を入れておく
// const (
// 	Block    = "B"
// 	position = 1
// )

// 初期値に従い空ならhandlerで設定
type DemoData struct {
	IpAddr string
	StHr   int   //初期値0
	StMn   int   //初期値0
	EdHr   int   //初期値23
	EdMn   int   //初期値59
	InMn   int   //初期値1
	DuMn   int   //初期値1
	Rly    []int //onにするrlyのリスト(0~7), handler.goでonじゃないものは除外
}

const (
	inMn = 1
	duMn = 1
)

func SettingRlyDemo(rly []int) []int {
	rly_l := 0
	rly_h := 0
	rly_out := make([]int, 2)
	for _, v := range rly {
		if v < 4 {
			rly_l += int(math.Pow(4.0, (3.0-float64(v))) * 3.0)
		} else if v < 8 {
			rly_h += int(math.Pow(4.0, (7.0-float64(v))) * 3.0)
		}
	}
	rly_out[0] = rly_l
	rly_out[1] = rly_h
	return rly_out
}

func BuildDemoM304(demoData *DemoData) (int, error) {
	ip_addr := demoData.IpAddr
	sthr := demoData.StHr
	stmn := demoData.StMn
	edhr := demoData.EdHr
	edmn := demoData.EdMn
	inmn := inMn
	dumn := duMn
	rly := SettingRlyDemo(demoData.Rly)
	rly_l := rly[0]
	rly_h := rly[1]
	resp, err := m304.SendDemo(ip_addr, sthr, stmn, edhr, edmn, inmn, dumn, rly_l, rly_h)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error in response with status code: %d", resp.StatusCode)
	}
	return 1, nil
}
