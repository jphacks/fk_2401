package m304

import (
	"fmt"
	"net/http"
	"strings"
)

func padding(s string, out_len int, padchar string) string {
	padding := out_len - len(s)
	if padding > 0 {
		return strings.Repeat(padchar, padding) + s
	} else {
		return s
	}
}

func byte_arrange(n int) string {
	hexstr := fmt.Sprintf("%x", n)
	value := padding(hexstr, 2, "0")
	return value
}

func send_BlockA(LC_UECS_ID string,
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
		ih_ip_addr += byte_arrange(int(v))
	}

	url := ""
	resp, _ := http.Get(url)
}

func send_BlockB() {

}

func send_BlockC() {

}

func send_BlockD() {

}
