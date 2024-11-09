package domain

type M304 struct {
	ID         int
	UecsID     string
	MacAddr    string
	DhcpFlg    bool
	IpAddr     *string
	NetMask    *string
	Defgw      *string
	Dns        *string
	VenderName string
	NodeName   *string
}

func NewM304(uecsID string, macAddr string, dhcpFlg bool, ipAddr *string, netMask *string, defgw *string, dns *string, venderName string, nodeName *string) *M304 {
	return &M304{
		UecsID:     uecsID,
		MacAddr:    macAddr,
		DhcpFlg:    dhcpFlg,
		IpAddr:     ipAddr,
		NetMask:    netMask,
		Defgw:      defgw,
		Dns:        dns,
		VenderName: venderName,
		NodeName:   nodeName,
	}
}

func NewM304WithID(id int, uecsID string, macAddr string, dhcpFlg bool, ipAddr *string, netMask *string, defgw *string, dns *string, venderName string, nodeName *string) *M304 {
	return &M304{
		ID:         id,
		UecsID:     uecsID,
		MacAddr:    macAddr,
		DhcpFlg:    dhcpFlg,
		IpAddr:     ipAddr,
		NetMask:    netMask,
		Defgw:      defgw,
		Dns:        dns,
		VenderName: venderName,
		NodeName:   nodeName,
	}
}
