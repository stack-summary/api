package model

type Category int

const (
	Paas Category = iota
	Serverless
	Rdb
	DocumentDb
	ObjectStorage
	Logging
	Monitoring
	Authentication
	CiCd
)
