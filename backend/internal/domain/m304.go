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
	Rly0       *int
	Rly1       *int
	Rly2       *int
	Rly3       *int
	Rly4       *int
	Rly5       *int
	Rly6       *int
	Rly7       *int
}

func NewM304(uecsID string, macAddr string, dhcpFlg bool, ipAddr *string, netMask *string, defgw *string, dns *string, venderName string, nodeName *string, rly0 *int, rly1 *int, rly2 *int, rly3 *int, rly4 *int, rly5 *int, rly6 *int, rly7 *int) *M304 {
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
		Rly0:       rly0,
		Rly1:       rly1,
		Rly2:       rly2,
		Rly3:       rly3,
		Rly4:       rly4,
		Rly5:       rly5,
		Rly6:       rly6,
		Rly7:       rly7,
	}
}

func NewM304WithID(id int, uecsID string, macAddr string, dhcpFlg bool, ipAddr *string, netMask *string, defgw *string, dns *string, venderName string, nodeName *string, rly0 *int, rly1 *int, rly2 *int, rly3 *int, rly4 *int, rly5 *int, rly6 *int, rly7 *int) *M304 {
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
		Rly0:       rly0,
		Rly1:       rly1,
		Rly2:       rly2,
		Rly3:       rly3,
		Rly4:       rly4,
		Rly5:       rly5,
		Rly6:       rly6,
		Rly7:       rly7,
	}
}
