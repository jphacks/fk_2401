package m304

import (
	"fmt"
	"net/http"
)

const (
	addrStmn = 0x1067
)

func SendDemo(ipAddr string, sthr int, stmn int, edhr int, edmn int, inmn int, dumn int, rlyL int, rlyH int) (*http.Response, error) {
	address := addrStmn
	ih_sthr := ByteArrange(sthr)
	ih_stmn := ByteArrange(stmn)
	ih_edhr := ByteArrange(edhr)
	ih_edmn := ByteArrange(edmn)
	ih_inmn := ByteArrange(inmn)
	ih_dumn := ByteArrange(dumn)
	ih_rly_l := ByteArrange(rlyL)
	ih_rly_h := ByteArrange(rlyH)
	ihtxt := ih_sthr + ih_stmn + ih_edhr + ih_edmn + ih_inmn + ih_dumn + ih_rly_l + ih_rly_h
	iht := ""
	if len(ihtxt) < blockSize {
		iht = ihtxt
	} else {
		return nil, fmt.Errorf("ihtxt length exceeds blockSize")
	}
	sz := Padding(fmt.Sprintf("%x", len(iht)/2), 2, "0")
	addr := Padding(fmt.Sprintf("%x", address), 4, "0")
	ih := ":" + sz + addr + "00" + iht + "FF"
	// 送信処理
	url := "http://" + ipAddr + "/" + ih
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return resp, nil
}
