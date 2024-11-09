package repository

import (
	"context"

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
	}

	id, err := mr.queries.CreateM304(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (mr M304Repository) GetM304FromID(ID int) (*domain.M304, error) {
	ctx := context.Background()

	m304, err := mr.queries.GetM304FromID(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	getM304 := domain.NewM304(
		m304.UecsID,
		m304.MacAddr,
		m304.DhcpFlg,
		NullStringToPointer(m304.IpAddr),
		NullStringToPointer(m304.NetMask),
		NullStringToPointer(m304.Defgw),
		NullStringToPointer(m304.Dns),
		m304.VenderName,
		NullStringToPointer(m304.NodeName),
	)

	return getM304, nil
}
