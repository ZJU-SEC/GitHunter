package repo

import (
	"GitHunter/config"
	"GitHunter/model"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/shomali11/parallelizer"
)

func Clone() {
	group := parallelizer.NewGroup(
		parallelizer.WithPoolSize(config.WORKER),
		parallelizer.WithJobQueueSize(config.QUEUE_SIZE),
	)
	defer group.Close()

	rows, _ := model.DB.Model(&model.Repo{}).Where("checked = ?", false).Rows()

	for rows.Next() {
		var s model.Repo
		model.DB.ScanRows(rows, &s)

		group.Add(func() {
			s := s
			cloneRepo(&s)
		})
	}

	group.Wait()
}


func cloneRepo(repo *model.Repo) {
    if _, err := os.Stat(repo.LocalPath()); !os.IsNotExist(err) {
		repo.Check()
		return
	}

	if _, err := git.PlainClone(repo.LocalPath(), false, &git.CloneOptions{
		URL: repo.GitURL(),
	}); err != nil {
		switch err {
		//case transport.ErrEmptyRemoteRepository, transport.ErrAuthenticationRequired:
		//	script.Delete()
		default:
			fmt.Println("[ERR] cannot clone", repo.Ref, err)
		}
		os.RemoveAll(repo.LocalPath())
		return
	}
    repo.Check()
    fmt.Println(repo.Ref, "cloned")
}
