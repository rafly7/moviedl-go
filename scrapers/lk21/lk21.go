package lk21

import (
	"moviedl/models/lk21"
)

type LK21 interface {
	Latest(page int) *[]lk21.Latest
	Download(slug string) *[]lk21.LinkBypass
	SearchByTitle(title string) *[]lk21.Movie
	Popular() *[]lk21.Popular
}

type ScrapType interface {
	LK21
}

func CollectorImpl() *Collector {
	return &Collector{}
}

type Collector struct {
}
