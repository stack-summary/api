package model

import "time"

type Product struct {
	Name       string
	LogoUrl    string
	Url        string
	Version    string
	LastUpdate time.Time
	Specs      Specs
	Category   Category
}
