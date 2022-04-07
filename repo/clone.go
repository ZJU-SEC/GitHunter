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
	os.RemoveAll(config.CLONE_STORAGE_PATH)

	group := parallelizer.NewGroup(
		parallelizer.WithPoolSize(config.CLONE_WORKER_NUMBER),
		parallelizer.WithJobQueueSize(config.CLONE_BATCH_SIZE),
	)
	defer group.Close()

	rows, _ := model.DB.Model(&model.Repo{}).Where("star_count >= ? and language = ?", config.CLONE_LOWER_BOUND, config.LANGUAGE).Rows()

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

func getClonePath(repo *model.Repo) string {
	return config.CLONE_STORAGE_PATH + "/" + "\"" + repo.Ref + "\""
}

func getCloneURL(repo *model.Repo) string {
	addr := "https://github.com/" + repo.Ref + ".git"
	return addr
}

func cloneRepo(repo *model.Repo) {
	if config.DEBUG {
		print("Cloning " + repo.Ref + " to " + getClonePath(repo) + "\n")
	}
	if _, err := git.PlainClone(getClonePath(repo), false, &git.CloneOptions{
		URL:        getCloneURL(repo),
		NoCheckout: true,
	}); err != nil {
		switch err {
		//case transport.ErrEmptyRemoteRepository, transport.ErrAuthenticationRequired:
		//	script.Delete()
		default:
			fmt.Println("[ERR] cannot clone", repo.Ref, err)
		}
		os.RemoveAll(getClonePath(repo))
		return
	}
}
