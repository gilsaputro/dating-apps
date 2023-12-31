package models

import "github.com/jinzhu/gorm"

// UserMatchHistory struct to user match history information
type UserMatchHistory struct {
	gorm.Model
	UserID      uint
	PartnerID   uint
	PartnerName string
	Status      MatchStatus
}

type MatchStatus int

const (
	MatchStatusUnkown   MatchStatus = -1
	MatchStatusPending  MatchStatus = 1
	MatchStatusApproved MatchStatus = 2
	MatchStatusRejected MatchStatus = 3
)

var MatchStatusToString = map[MatchStatus]string{
	MatchStatusUnkown:   "UNKNOWN",
	MatchStatusPending:  "PENDING",
	MatchStatusApproved: "APPROVED",
	MatchStatusRejected: "REJECTED",
}

func (m MatchStatus) String() string {
	if val, ok := MatchStatusToString[m]; ok {
		return val
	}
	return MatchStatusToString[-1]
}
