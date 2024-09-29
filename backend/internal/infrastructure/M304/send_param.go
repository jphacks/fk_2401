package m304

import (
	"fmt"
	"math"

	// "net/http"
	"strconv"
	"strings"
)

// 文字列を指定長に変換する(右埋め)
// s: 元の文字列, out_len: 出力文字列長, padchar: paddingに使う文字
func padding(s string, out_len int, padchar string) string {
	padding := out_len - len(s)
	if padding > 0 {
		return strings.Repeat(padchar, padding) + s
	} else {
		return s
	}
}

// 整数を2桁16進数の文字列に変換する
func byte_arrange(n int) string {
	hexstr := fmt.Sprintf("%x", n)
	value := padding(hexstr, 2, "0")
	return value
}

// 文字列を指定長の2桁16進数asciiコードの文字列に変換する
func string_arrange(s string, length int) string {
	rt := ""
	runes := []rune(s)
	for _, r := range runes {
		rt += padding(fmt.Sprintf("%x", r), 2, "0")
	}
	for len(rt) < length {
		rt += "0"
	}
	return rt
}

// 32bit浮動小数点数を16進数に変換
func float32_bin(f float32) string {
	bin := math.Float32bits(f)
	hexstr := fmt.Sprintf("%08x", bin)
	return hexstr
}

func send_BlockA(IP_ADDR string,
	LC_UECS_ID string,
	LC_MAC string,
	FIX_DHCP_FLAG int,
	FIXED_IPADDRESS string,
	FIXED_NETMASK string,
	FIXED_DEFGW string,
	FIXED_DNS string,
	VENDER_NAME string,
	NODE_NAME string) {
	address := 0x0000
	ih_uecs_id := padding(LC_UECS_ID, 12, "0")
	ih_mac := strings.Replace(LC_MAC, ":", "", -1)
	ih_dhcpflg := byte_arrange(FIX_DHCP_FLAG)
	sp_ip_addr := strings.Split(FIXED_IPADDRESS, ".")
	ih_ip_addr := ""
	for _, v := range sp_ip_addr {
		addrInt, _ := strconv.Atoi(v)
		ih_ip_addr += byte_arrange(addrInt)
	}

	sp_netmask := strings.Split(FIXED_NETMASK, ".")
	ih_netmask := ""
	for _, v := range sp_netmask {
		netmaskInt, _ := strconv.Atoi(v)
		ih_netmask += byte_arrange(netmaskInt)
	}

	sp_defgw := strings.Split(FIXED_DEFGW, ".")
	ih_defgw := ""
	for _, v := range sp_defgw {
		defgwInt, _ := strconv.Atoi(v)
		ih_defgw += byte_arrange(defgwInt)
	}

	sp_dns := strings.Split(FIXED_DNS, ".")
	ih_dns := ""
	for _, v := range sp_dns {
		dnsInt, _ := strconv.Atoi(v)
		ih_dns += byte_arrange(dnsInt)
	}

	ih_vender_name := string_arrange(VENDER_NAME, 32)
	ih_node_name := string_arrange(NODE_NAME, 32)

	ihtxt := ih_uecs_id + ih_mac + ih_dhcpflg + ih_ip_addr + ih_netmask + ih_defgw + ih_dns + ih_vender_name + ih_node_name

	for i := range 4 {
		tp := i * 32
		iht := ""
		if len(ihtxt) < tp+32 {
			iht = ihtxt[tp:]
		} else {
			iht = ihtxt[tp:(tp + 32)]
		}
		sz := padding(fmt.Sprintf("%x", len(iht)/2), 2, "0")
		addr := padding(fmt.Sprintf("%x", address+(tp/2)), 4, "0")
		ih := ":" + sz + addr + "00" + iht + "FF"
		// 送信処理
		url := "http://" + IP_ADDR + "/" + ih
		fmt.Println(url)
		// resp, _ := http.Get(url)
	}
}

func send_BlockB(B_ID int,
	IP_ADDR string,
	LC_VALID int,
	LC_ROOM int,
	LC_REGION int,
	LC_ORDER int,
	LC_PRIORITY int,
	LC_LV int,
	LC_CAST int,
	LC_SR string,
	LC_CCMTYPE string,
	LC_UNIT string,
	LC_STHR int,
	LC_STMN int,
	LC_EDHR int,
	LC_EDMN int,
	LC_INMN int,
	LC_DUMN int,
	LC_RLY_L int,
	LC_RLY_H int) {
	address := 0x1000
	recstep := 0x40
	ih_valid := byte_arrange(LC_VALID)
	ih_room := byte_arrange(LC_ROOM)
	ih_region := byte_arrange(LC_REGION)
	order_o := padding(fmt.Sprintf("%x", LC_ORDER), 4, "0")
	ih_order := order_o[2:4] + order_o[0:2]
	ih_priority := byte_arrange(LC_PRIORITY)
	ih_lv := byte_arrange(LC_LV)
	ih_cast := byte_arrange(LC_CAST)
	ih_sr := string_arrange(LC_SR, 2)
	ih_ccmtype := string_arrange(LC_CCMTYPE, 40)
	ih_unit := string_arrange(LC_UNIT, 20)
	ih_sthr := byte_arrange(LC_STHR)
	ih_stmn := byte_arrange(LC_STMN)
	ih_edhr := byte_arrange(LC_EDHR)
	ih_edmn := byte_arrange(LC_EDMN)
	ih_inmn := byte_arrange(LC_INMN)
	ih_dumn := byte_arrange(LC_DUMN)
	ih_rly_l := byte_arrange(LC_RLY_L)
	ih_rly_h := byte_arrange(LC_RLY_H)

	ihtxt := ih_valid + ih_room + ih_region + ih_order + ih_priority + ih_lv + ih_cast + ih_sr + ih_ccmtype + ih_unit + ih_sthr + ih_stmn + ih_edhr + ih_edmn + ih_inmn + ih_dumn + ih_rly_l + ih_rly_h
	for i := range 3 {
		tp := i * 32
		iht := ""
		if len(ihtxt) < tp+32 {
			iht = ihtxt[tp:]
		} else {
			iht = ihtxt[tp:(tp + 32)]
		}
		sz := padding(fmt.Sprintf("%x", len(iht)/2), 2, "0")
		addr := padding(fmt.Sprintf("%x", B_ID*recstep+address+(tp/2)), 4, "0")
		ih := ":" + sz + addr + "00" + iht + "FF"
		// 送信処理
		url := "http://" + IP_ADDR + "/" + ih
		fmt.Println(url)
		// resp, _ := http.Get(url)
	}
}

func send_BlockC(C_ID int,
	IP_ADDR string,
	LC_VALID int,
	LC_ROOM int,
	LC_REGION int,
	LC_ORDER int,
	LC_PRIORITY int,
	LC_LV int,
	LC_CAST int,
	LC_SR string,
	LC_CCMTYPE string,
	LC_UNIT string,
	LC_STHR int,
	LC_STMN int,
	LC_EDHR int,
	LC_EDMN int,
	LC_INMN int,
	LC_DUMN int,
	LC_RLY_L int,
	LC_RLY_H int) {
	address := 0x3000
	recstep := 0x40
	ih_valid := byte_arrange(LC_VALID)
	ih_room := byte_arrange(LC_ROOM)
	ih_region := byte_arrange(LC_REGION)
	order_o := padding(fmt.Sprintf("%x", LC_ORDER), 4, "0")
	ih_order := order_o[2:4] + order_o[0:2]
	ih_priority := byte_arrange(LC_PRIORITY)
	ih_lv := byte_arrange(LC_LV)
	ih_cast := byte_arrange(LC_CAST)
	ih_sr := string_arrange(LC_SR, 2)
	ih_ccmtype := string_arrange(LC_CCMTYPE, 40)
	ih_unit := string_arrange(LC_UNIT, 20)
	ih_sthr := byte_arrange(LC_STHR)
	ih_stmn := byte_arrange(LC_STMN)
	ih_edhr := byte_arrange(LC_EDHR)
	ih_edmn := byte_arrange(LC_EDMN)
	ih_inmn := byte_arrange(LC_INMN)
	ih_dumn := byte_arrange(LC_DUMN)
	ih_rly_l := byte_arrange(LC_RLY_L)
	ih_rly_h := byte_arrange(LC_RLY_H)

	ihtxt := ih_valid + ih_room + ih_region + ih_order + ih_priority + ih_lv + ih_cast + ih_sr + ih_ccmtype + ih_unit + ih_sthr + ih_stmn + ih_edhr + ih_edmn + ih_inmn + ih_dumn + ih_rly_l + ih_rly_h
	for i := range 3 {
		tp := i * 32
		iht := ""
		if len(ihtxt) < tp+32 {
			iht = ihtxt[tp:]
		} else {
			iht = ihtxt[tp:(tp + 32)]
		}
		sz := padding(fmt.Sprintf("%x", len(iht)/2), 2, "0")
		addr := padding(fmt.Sprintf("%x", C_ID*recstep+address+(tp/2)), 4, "0")
		ih := ":" + sz + addr + "00" + iht + "FF"
		// 送信処理
		url := "http://" + IP_ADDR + "/" + ih
		fmt.Println(url)
		// resp, _ := http.Get(url)
	}
}

func send_BlockD(D_ID int,
	IP_ADDR string,
	LC_COPE_VALID int,
	LC_COPE_ROOM int,
	LC_COPE_REGION int,
	LC_COPE_ORDER int,
	LC_COPE_PRIORITY int,
	LC_COPE_CCMTYPE string,
	LC_COPE_OPE int,
	LC_COPE_FVAL float32) {
	address := 0x5000
	recstep := 0x20
	ih_cope_valid := byte_arrange(LC_COPE_VALID)
	ih_cope_room := byte_arrange(LC_COPE_ROOM)
	ih_cope_region := byte_arrange(LC_COPE_REGION)
	cope_order_o := padding(fmt.Sprintf("%x", LC_COPE_ORDER), 4, "0")
	ih_cope_order := cope_order_o[2:4] + cope_order_o[0:2]
	ih_cope_priority := byte_arrange(LC_COPE_PRIORITY)
	ih_cope_ccmtype := string_arrange(LC_COPE_CCMTYPE, 40)
	ih_cope_ope := byte_arrange(LC_COPE_OPE)
	ih_cope_fval := float32_bin(LC_COPE_FVAL)

	ihtxt := ih_cope_valid + ih_cope_room + ih_cope_region + ih_cope_order + ih_cope_priority + ih_cope_ccmtype + ih_cope_ope + ih_cope_fval
	for i := range 2 {
		tp := i * 32
		iht := ""
		if len(ihtxt) < tp+32 {
			iht = ihtxt[tp:]
		} else {
			iht = ihtxt[tp:(tp + 32)]
		}
		sz := padding(fmt.Sprintf("%x", len(iht)/2), 2, "0")
		addr := padding(fmt.Sprintf("%x", D_ID*recstep+address+(tp/2)), 4, "0")
		ih := ":" + sz + addr + "00" + iht + "FF"
		// 送信処理
		url := "http://" + IP_ADDR + "/" + ih
		fmt.Println(url)
		// resp, _ := http.Get(url)
	}
}

func test_control() {
	print("A\n")
	send_BlockA("192.168.1.14", "10100C00000B", "02:A2:73:0B:00:2A", 0, "192.168.38.50", "255.255.255.0", "192.168.11.1", "192.168.11.1", "AMPSD", "TESTA123")
	print("B\n")
	send_BlockB(0, "192.168.1.14", 1, 1, 1, 1, 15, 3, 0, "R", "InAirHumid", "%", 0, 0, 23, 59, 1, 1, 252, 0)
	send_BlockB(1, "192.168.1.14", 1, 1, 1, 1, 15, 3, 0, "R", "InAirHumid", "%", 0, 0, 23, 59, 1, 1, 0, 255)
	print("C\n")
	send_BlockC(0, "192.168.1.14", 1, 1, 1, 1, 15, 3, 0, "R", "TEST456", "", 0, 0, 23, 59, 1, 1, 252, 0)
	send_BlockC(1, "192.168.1.14", 1, 1, 1, 1, 15, 3, 0, "R", "TEST789", "", 0, 0, 12, 00, 1, 1, 0, 255)
	print("D\n")
	var f float32 = 1.500
	send_BlockD(0, "192.168.1.14", 1, 1, 1, 1, 15, "InAirHumid", 3, f)
}
