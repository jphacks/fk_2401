package repository

import "database/sql"

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

func PointerToNullFloat64(n *float64) sql.NullFloat64 {
	nf := sql.NullFloat64{}
	if n != nil {
		nf.Float64 = *n
		nf.Valid = true
		return nf
	}
	nf.Float64 = float64(0)
	nf.Valid = false
	return nf
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

func NullFloat64ToPointer(ni sql.NullFloat64) *float64 {
	var pf *float64
	if ni.Valid {
		n := ni.Float64
		pf = &n
	}
	return pf
}
