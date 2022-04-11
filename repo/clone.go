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

	rows, _ := model.DB.Model(&model.Repo{}).Rows()

	for rows.Next() {
		var r model.Repo
		model.DB.ScanRows(rows, &r)

		group.Add(func() {
			r := r
			cloneRepo(&r)
		})
	}

	group.Wait()
}


func cloneRepo(repo *model.Repo) {
    if _, err := os.Stat(repo.LocalPath()); !os.IsNotExist(err) {
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
    fmt.Println(repo.Ref, "cloned")
}
