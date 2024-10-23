package repository

import (
	"context"
	"database/sql"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type M304Repository struct {
	queries *mysqlc.Queries
}

func NewM304Repository(queries *mysqlc.Queries) *M304Repository {
	return &M304Repository{
		queries: queries,
	}
}

func PointerToNullString(str *string) sql.NullString {
	ns := sql.NullString{}
	if str != nil {
		ns.String = *str
		ns.Valid = true
		return ns
	}
	ns.String = ""
	ns.Valid = false
	return ns
}

func PointerToNullInt32(n *int) sql.NullInt32 {
	ni := sql.NullInt32{}
	if n != nil {
		ni.Int32 = int32(*n)
		ni.Valid = true
		return ni
	}
	ni.Int32 = int32(0)
	ni.Valid = false
	return ni
}

func NullStringToPointer(ns sql.NullString) *string {
	var ps *string
	if ns.Valid {
		s := ns.String
		ps = &s
	}
	return ps
}

func NullInt32ToPointer(ni sql.NullInt32) *int {
	var pi *int
	if ni.Valid {
		n := int(ni.Int32)
		pi = &n
	}
	return pi
}

func (mr M304Repository) CreateM304(newM304 domain.M304) (int64, error) {
	ctx := context.Background()

	arg := mysqlc.CreateM304Params{
		UecsID:     newM304.UecsID,
		MacAddr:    newM304.MacAddr,
		DhcpFlg:    newM304.DhcpFlg,
		IpAddr:     PointerToNullString(newM304.IpAddr),
		NetMask:    PointerToNullString(newM304.NetMask),
		Defgw:      PointerToNullString(newM304.Defgw),
		Dns:        PointerToNullString(newM304.Dns),
		VenderName: newM304.VenderName,
		NodeName:   PointerToNullString(newM304.NodeName),
		Rly0:       PointerToNullInt32(newM304.Rly0),
		Rly1:       PointerToNullInt32(newM304.Rly1),
		Rly2:       PointerToNullInt32(newM304.Rly2),
		Rly3:       PointerToNullInt32(newM304.Rly3),
		Rly4:       PointerToNullInt32(newM304.Rly4),
		Rly5:       PointerToNullInt32(newM304.Rly5),
		Rly6:       PointerToNullInt32(newM304.Rly6),
		Rly7:       PointerToNullInt32(newM304.Rly7),
	}

	id, err := mr.queries.CreateM304(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (mr M304Repository) GetM304FromUecsDevice(uecsDeviceID int) ([]*domain.M304, error) {
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
		m304s[i] = domain.NewM304WithID(
			int(v.ID),
			v.UecsID,
			v.MacAddr,
			v.DhcpFlg,
			NullStringToPointer(v.IpAddr),
			NullStringToPointer(v.NetMask),
			NullStringToPointer(v.Defgw),
			NullStringToPointer(v.Dns),
			v.VenderName,
			NullStringToPointer(v.NodeName),
			NullInt32ToPointer(v.Rly0),
			NullInt32ToPointer(v.Rly1),
			NullInt32ToPointer(v.Rly2),
			NullInt32ToPointer(v.Rly3),
			NullInt32ToPointer(v.Rly4),
			NullInt32ToPointer(v.Rly5),
			NullInt32ToPointer(v.Rly6),
			NullInt32ToPointer(v.Rly7),
		)
	}

	return m304s, nil
}
