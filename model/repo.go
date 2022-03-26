package model

import (
	"fmt"
	"sync"
)

// Repo struct map the schema of a repository
type Repo struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Ref   string `gorm:"uniqueIndex" json:"full_name"`
	Owner string

	// description
	Language      string `json:"language"`
	Description   string `json:"description"`
	DefaultBranch string `json:"default_branch"`

	// boolean items
	IsFork     bool `json:"fork"`
	IsArchive  bool `json:"archived"`
	IsTemplate bool `json:"is_template"`
	OrgProj    bool

	// numeric items
	Size      uint `json:"size"`
	StarCount uint `json:"stargazers_count"`
	ForkCount uint `json:"forks"`

	// timestamp
	CreatedAt int64 // unix timestamp of the creation time
	UpdatedAt int64 // unix timestamp of the update time

	// raw data, process required
	RawOwner     Owner  `gorm:"-" json:"owner"`
	RawCreatedAt string `gorm:"-" json:"created_at"`
	RawUpdatedAt string `gorm:"-" json:"pushed_at"`
}

type Owner struct {
	Name      string `json:"login"`
	OwnerType string `json:"type"`
}

func (r *Repo) preprocess() {
	// TODO owner
	// TODO OrgProj
	// TODO CreatedAt
	// TODO UpdatedAt
	fmt.Println(r)
}

func CreateRepoBatch(repos []Repo) {
	// preprocess the records
	for i := 0; i <= len(repos); i++ {
		repos[i].preprocess()
	}

	var mutex sync.Mutex
	mutex.Lock()

	// TODO create repos record

	mutex.Unlock()
}
