package models

import "errors"

var (
	NotFound       = errors.New("NotFound")
	ErrGettingSite = errors.New("Error getting site with max access time")
)
