package api

import "time"

type StuApplicationResponse struct {
	ApplicationID      int64     `json:"applicationID"`
	UserName           string    `json:"username"`
	University         string    `json:"university"`
	Major              string    `json:"major"`
	OriginalUniversity string    `json:"originalUniversity"`
	OriginalMajor      string    `json:"originalMajor"`
	IsAdmitted         int64     `json:"isAdmitted"`
	IsConfirmed        int64     `json:"isConfirmed"`
	SubmitTime         time.Time `json:"submitTime"`
}
