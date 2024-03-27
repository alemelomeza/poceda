package domain

import "time"

type Subscription struct {
	CaseID int
	LDAP string
	CreatedAt time.Time
}
