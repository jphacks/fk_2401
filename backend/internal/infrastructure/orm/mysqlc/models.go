// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package mysqlc

import (
	"database/sql"
	"time"
)

type ClimateData struct {
	ID   int32
	Name string
	Unit string
}

type Device struct {
	ID            int32
	HouseID       int32
	ClimateDataID int32
	DeviceName    sql.NullString
	SetPoint      sql.NullFloat64
	Duration      sql.NullInt32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type House struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type M304 struct {
	ID         int32
	UecsID     string
	MacAddr    string
	DhcpFlg    bool
	IpAddr     sql.NullString
	NetMask    sql.NullString
	Defgw      sql.NullString
	Dns        sql.NullString
	VenderName string
	NodeName   sql.NullString
	Rly0       sql.NullInt32
	Rly1       sql.NullInt32
	Rly2       sql.NullInt32
	Rly3       sql.NullInt32
	Rly4       sql.NullInt32
	Rly5       sql.NullInt32
	Rly6       sql.NullInt32
	Rly7       sql.NullInt32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UecsDevice struct {
	ID        int32
	Ccmtype   string
	Room      int32
	Region    int32
	Order     int32
	Priority  int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
