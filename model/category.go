package model

type Category int

const (
	Paas Category = iota
	Rdb
	DocumentDb
	ObjectStorage
	Logging
	Monitoring
	Authentication
	CiCd
)
