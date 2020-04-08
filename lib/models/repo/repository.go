package repo

import (
	"crypto/md5"
	"dragoncore/lib/models/md"
	"encoding/hex"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// repository
type repository struct {
	url RepositoryURL
	repo *git.Repository
}

// New returns a repository
func New(url RepositoryURL) *repository {
	return &repository{url, nil}
}

// getBasePath gets the path where repositories are stored
func (_ repository) getBasePath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path.Join(dir, ".external")
}

// getRepoDir gets the local repo directory
func (r *repository) getRepoPath() string {
	hash := md5.New()
	hash.Write([]byte(r.url))
	return path.Join(r.getBasePath(), hex.EncodeToString(hash.Sum(nil)))
}


// Clone clones the repository if it hasn't been cloned yet
func (r *repository) Clone() (err error) {
	if r.repo != nil {
		return nil
	}

	c := make(chan interface{})
	r.repo, err = git.PlainClone(r.getRepoPath(), false, &git.CloneOptions{
		URL:      string(r.url),
		Progress: cloneProgress{c},
	})

	if err != nil {
		if err.Error() == git.ErrRepositoryAlreadyExists.Error() {
			if r.repo, err = git.PlainOpen(r.getRepoPath()); err != nil {
				return err
			}
			return nil
		}
		return err
	}
	<-c
	return nil
}

// Update updates the content of the repository
func (r *repository) Update() (err error) {
	if err = r.Clone(); err != nil {
		return err
	}

	err = r.repo.Fetch(&git.FetchOptions{})
	return err
}

// Documentation returns all MD documentation
func (r *repository) Documentation() (files []md.Type, err error) {
	if err = r.Clone(); err != nil {
		return nil, err
	}

	dir := r.getRepoPath()
	visit := func (path string, f os.FileInfo, err error) error {
		if f.IsDir() && f.Name() == ".git" {
			return filepath.SkipDir
		}
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".md") {
			files = append(files, md.New(path, dir))
		}
		return nil
	}

	_ = filepath.Walk(dir, visit)
	return files, nil
}

