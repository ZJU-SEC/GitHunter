package model

import "time"

// Repo struct map the schema of a repository
type Repo struct {
	ID    uint   `gorm:"primaryKey"`
	Ref   string `gorm:"uniqueIndex"`
	Owner string

	// description
	Language      string
	Description   string
	DefaultBranch string

	// boolean items
	IsFork     bool
	IsArchive  bool
	IsTemplate bool

	// numeric items
	Size       uint
	StarCount  uint
	WatchCount uint
	ForkCount  uint

	// timestamp
	CreatedAt time.Time
	UpdatedAt time.Time
}
