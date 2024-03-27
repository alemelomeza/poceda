package domain

import "time"

type Action struct {
	CaseID     int
	LDAP       string
	ActionType string
	SLA        time.Time
	Comment    string
	CreatedAt  string
}
