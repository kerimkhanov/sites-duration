package models

const (
	MinName = "minName"
	MinTime = "minTime"
	MaxName = "maxName"
	MaxTime = "maxTime"
)

type SiteAccessTime struct {
	Site       string
	AccessTime float64
}

func (s *SiteAccessTime) SetSite(site string) {
	s.Site = site
}

func (s *SiteAccessTime) SetAccessTime(accessTime float64) {
	s.AccessTime = accessTime
}
