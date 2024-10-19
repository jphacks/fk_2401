package repository

import (
	"context"
	"database/sql"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type M304Repositoty struct {
	queries *mysqlc.Queries
}

func NewM304Repository(queries *mysqlc.Queries) *M304Repositoty {
	return &M304Repositoty{
		queries: queries,
	}
}

func (mr M304Repositoty) CreateM304(newM304 domain.M304) (int64, error) {
	ctx := context.Background()

	uecsid := sql.NullString{
		String: newM304.UecsID,
		Valid:  false,
	}
	if newM304.UecsID != "" {
		uecsid.Valid = true
	}

	mac_addr := sql.NullString{
		String: newM304.MacAddr,
		Valid:  false,
	}
	if newM304.MacAddr != "" {
		mac_addr.Valid = true
	}

	ip_addr := sql.NullString{
		String: newM304.IpAddr,
		Valid:  false,
	}
	if newM304.IpAddr != "" {
		ip_addr.Valid = true
	}

	net_mask := sql.NullString{
		String: newM304.NetMask,
		Valid:  false,
	}
	if newM304.NetMask != "" {
		net_mask.Valid = true
	}

	defgw := sql.NullString{
		String: newM304.Defgw,
		Valid:  false,
	}
	if newM304.Defgw != "" {
		defgw.Valid = true
	}

	dns := sql.NullString{
		String: newM304.Dns,
		Valid:  false,
	}
	if newM304.Dns != "" {
		dns.Valid = true
	}

	vender_name := sql.NullString{
		String: newM304.VenderName,
		Valid:  false,
	}
	if newM304.VenderName != "" {
		vender_name.Valid = true
	}

	node_name := sql.NullString{
		String: newM304.NodeName,
		Valid:  false,
	}
	if newM304.NodeName != "" {
		node_name.Valid = true
	}

	rly0 := sql.NullInt32{
		Int32: int32(newM304.Rly0),
		Valid: false,
	}
	if newM304.Rly0 != 0 {
		rly0.Valid = true
	}

	rly1 := sql.NullInt32{
		Int32: int32(newM304.Rly1),
		Valid: false,
	}
	if newM304.Rly1 != 0 {
		rly1.Valid = true
	}

	rly2 := sql.NullInt32{
		Int32: int32(newM304.Rly2),
		Valid: false,
	}
	if newM304.Rly2 != 0 {
		rly2.Valid = true
	}

	rly3 := sql.NullInt32{
		Int32: int32(newM304.Rly3),
		Valid: false,
	}
	if newM304.Rly3 != 0 {
		rly3.Valid = true
	}

	rly4 := sql.NullInt32{
		Int32: int32(newM304.Rly4),
		Valid: false,
	}
	if newM304.Rly4 != 0 {
		rly4.Valid = true
	}

	rly5 := sql.NullInt32{
		Int32: int32(newM304.Rly5),
		Valid: false,
	}
	if newM304.Rly5 != 0 {
		rly5.Valid = true
	}

	rly6 := sql.NullInt32{
		Int32: int32(newM304.Rly6),
		Valid: false,
	}
	if newM304.Rly6 != 0 {
		rly6.Valid = true
	}

	rly7 := sql.NullInt32{
		Int32: int32(newM304.Rly7),
		Valid: false,
	}
	if newM304.Rly7 != 0 {
		rly7.Valid = true
	}

	arg := mysqlc.CreateM304Params{
		UecsID:     uecsid,
		MacAddr:    mac_addr,
		DhcpFlg:    newM304.DhcpFlg,
		IpAddr:     ip_addr,
		NetMask:    net_mask,
		Defgw:      defgw,
		Dns:        dns,
		VenderName: vender_name,
		NodeName:   node_name,
		Rly0:       rly0,
		Rly1:       rly1,
		Rly2:       rly2,
		Rly3:       rly3,
		Rly4:       rly4,
		Rly5:       rly5,
		Rly6:       rly6,
		Rly7:       rly7,
	}

	id, err := mr.queries.CreateM304(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (mr M304Repositoty) GetM304FromUecsDevice(uecsDeviceID int) ([]*domain.M304, error) {
	ctx := context.Background()

	rly := sql.NullInt32{
		Int32: int32(uecsDeviceID),
		Valid: true,
	}
	m304sRow, err := mr.queries.GetM304FromUecsDevice(ctx, rly)
	if err != nil {
		return nil, err
	}

	m304s := make([]*domain.M304, len(m304sRow))
	for i, v := range m304sRow {
		uecsid := ""
		if v.UecsID.Valid {
			uecsid = v.UecsID.String
		}
		mac_addr := ""
		if v.MacAddr.Valid {
			mac_addr = v.MacAddr.String
		}
		ip_addr := ""
		if v.IpAddr.Valid {
			ip_addr = v.IpAddr.String
		}
		net_mask := ""
		if v.NetMask.Valid {
			net_mask = v.NetMask.String
		}
		defgw := ""
		if v.Defgw.Valid {
			defgw = v.Defgw.String
		}
		dns := ""
		if v.Dns.Valid {
			dns = v.Dns.String
		}
		vender_name := ""
		if v.VenderName.Valid {
			vender_name = v.VenderName.String
		}
		node_name := ""
		if v.NodeName.Valid {
			node_name = v.NodeName.String
		}
		rly0 := 0
		if v.Rly0.Valid {
			rly0 = int(v.Rly0.Int32)
		}
		rly1 := 0
		if v.Rly1.Valid {
			rly1 = int(v.Rly1.Int32)
		}
		rly2 := 0
		if v.Rly2.Valid {
			rly2 = int(v.Rly2.Int32)
		}
		rly3 := 0
		if v.Rly3.Valid {
			rly3 = int(v.Rly3.Int32)
		}
		rly4 := 0
		if v.Rly4.Valid {
			rly4 = int(v.Rly4.Int32)
		}
		rly5 := 0
		if v.Rly5.Valid {
			rly5 = int(v.Rly5.Int32)
		}
		rly6 := 0
		if v.Rly6.Valid {
			rly6 = int(v.Rly6.Int32)
		}
		rly7 := 0
		if v.Rly7.Valid {
			rly7 = int(v.Rly7.Int32)
		}
		m304s[i] = domain.NewM304WithID(
			int(v.ID),
			uecsid,
			mac_addr,
			v.DhcpFlg,
			ip_addr,
			net_mask,
			defgw,
			dns,
			vender_name,
			node_name,
			rly0,
			rly1,
			rly2,
			rly3,
			rly4,
			rly5,
			rly6,
			rly7,
		)
	}

	return m304s, nil
}
