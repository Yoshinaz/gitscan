package repo_test

import (
	"errors"
	"github.com/gitscan/constants/reqError"
	"github.com/gitscan/constants/status"
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/report"
	"github.com/gitscan/internal/service/repo"
	dbMocks "github.com/gitscan/mocks/db"
	repoMocks "github.com/gitscan/mocks/repo"
	ruleMocks "github.com/gitscan/mocks/rule"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	gitLib "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestRepo_Scan(t *testing.T) {
	commit := "1234"
	url := "github.com/test"
	name := "test repo"
	//should make a mock repository object instead of actual clone
	repository, err := gitLib.Clone(memory.NewStorage(), nil, &gitLib.CloneOptions{
		URL: "https://github.com/Yoshinaz/test_secret",
	})
	assert.Nil(t, err)

	t.Run("clone if repository is nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		rulesData := ruleMocks.NewMockInterface(ctrl)

		db.EXPECT().Info().Times(1).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{Status: status.QUEUED.String(), Commit: commit}, nil)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().GetHeadCommitHash().Return(commit, nil)
		repoMock.EXPECT().Repo().Return(nil)
		repoMock.EXPECT().Clone()
		repoMock.EXPECT().Rules().Return(rulesData)
		rulesData.EXPECT().RuleSet().Return("1234")

		w := make(chan bool, 1)
		s, err := repo.Scan(repoMock, "false", db, w)

		assert.Nil(t, err)
		assert.Equal(t, status.QUEUED.String(), s)
		assert.Equal(t, 0, len(w))
	})

	t.Run("Found repository in queue", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		rulesData := ruleMocks.NewMockInterface(ctrl)

		db.EXPECT().Info().Times(1).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{Status: status.QUEUED.String(), Commit: commit}, nil)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().GetHeadCommitHash().Return(commit, nil)
		repoMock.EXPECT().Repo().Return(&gitLib.Repository{})
		repoMock.EXPECT().Rules().Return(rulesData)
		rulesData.EXPECT().RuleSet().Return("1234")

		w := make(chan bool, 1)
		s, err := repo.Scan(repoMock, "false", db, w)

		assert.Nil(t, err)
		assert.Equal(t, status.QUEUED.String(), s)
		assert.Equal(t, 0, len(w))
	})

	t.Run("Found repository in progress", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		rulesData := ruleMocks.NewMockInterface(ctrl)

		db.EXPECT().Info().Times(1).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{Status: status.INPROGRESS.String(), Commit: commit}, nil)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().GetHeadCommitHash().Return(commit, nil)
		repoMock.EXPECT().Repo().Return(&gitLib.Repository{})
		repoMock.EXPECT().Rules().Return(rulesData)
		rulesData.EXPECT().RuleSet().Return("1234")

		w := make(chan bool, 1)
		s, err := repo.Scan(repoMock, "false", db, w)

		assert.Nil(t, err)
		assert.Equal(t, status.INPROGRESS.String(), s)
		assert.Equal(t, 0, len(w))
	})

	t.Run("Found repository success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		rulesData := ruleMocks.NewMockInterface(ctrl)

		db.EXPECT().Info().Times(1).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{Status: status.SUCCESS.String(), Commit: commit}, nil)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().GetHeadCommitHash().Return(commit, nil)
		repoMock.EXPECT().Repo().Return(&gitLib.Repository{})
		repoMock.EXPECT().Rules().Return(rulesData)
		rulesData.EXPECT().RuleSet().Return("1234")

		w := make(chan bool, 1)
		s, err := repo.Scan(repoMock, "false", db, w)

		assert.Nil(t, err)
		assert.Equal(t, status.SUCCESS.String(), s)
		assert.Equal(t, 0, len(w))
	})

	t.Run("Rescan : found with failed status", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		rulesData := ruleMocks.NewMockInterface(ctrl)

		infoData := database.Info{Status: status.FAILED.String(), Commit: commit}
		db.EXPECT().Info().Times(1).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(infoData, nil)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().GetHeadCommitHash().Return(commit, nil)
		repoMock.EXPECT().Repo().Return(&gitLib.Repository{})
		repoMock.EXPECT().Rules().Return(rulesData)
		rulesData.EXPECT().RuleSet().Return("1234")

		//process expected call
		numCommit, err := getNumberOfCommit(repository)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Repo().Times(2).Return(repository)
		db.EXPECT().Info().Times(2).Return(dbInfo)
		dbInfo.EXPECT().Update(gomock.Any()).Times(2).Return(infoData, nil)
		repoMock.EXPECT().Rules().Times(numCommit).Return(rulesData)
		rulesData.EXPECT().Process(gomock.Any()).Times(numCommit).Return([]report.Finding{}, nil)

		w := make(chan bool, 1)
		s, err := repo.Scan(repoMock, "false", db, w)
		time.Sleep(2 * time.Second)

		assert.Nil(t, err)
		assert.Equal(t, status.QUEUED.String(), s)

	})

	t.Run("Rescan : found with different hash commit", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		rulesData := ruleMocks.NewMockInterface(ctrl)

		infoData := database.Info{Status: status.SUCCESS.String(), Commit: "some commit"}
		db.EXPECT().Info().Times(1).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{Status: status.SUCCESS.String(), Commit: "some commit"}, nil)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().GetHeadCommitHash().Return(commit, nil)
		repoMock.EXPECT().Repo().Return(&gitLib.Repository{})
		repoMock.EXPECT().Rules().Return(rulesData)
		rulesData.EXPECT().RuleSet().Return("1234")

		//process expected call
		numCommit, err := getNumberOfCommit(repository)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Repo().Times(2).Return(repository)
		db.EXPECT().Info().Times(2).Return(dbInfo)
		dbInfo.EXPECT().Update(gomock.Any()).Times(2).Return(infoData, nil)
		repoMock.EXPECT().Rules().Times(numCommit).Return(rulesData)
		rulesData.EXPECT().Process(gomock.Any()).Times(numCommit).Return([]report.Finding{}, nil)

		w := make(chan bool, 1)
		s, err := repo.Scan(repoMock, "false", db, w)
		time.Sleep(2 * time.Second)

		assert.Nil(t, err)
		assert.Equal(t, status.QUEUED.String(), s)

	})

	t.Run("First time: do not found in database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		rulesData := ruleMocks.NewMockInterface(ctrl)

		db.EXPECT().Info().Times(2).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{}, gorm.ErrRecordNotFound)
		dbInfo.EXPECT().Create(gomock.Any()).Return(database.Info{}, nil)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().GetHeadCommitHash().Return(commit, nil)
		repoMock.EXPECT().Repo().Return(&gitLib.Repository{})
		repoMock.EXPECT().Rules().Return(rulesData)
		rulesData.EXPECT().RuleSet().Return("1234")

		//process expected call
		numCommit, err := getNumberOfCommit(repository)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Repo().Times(2).Return(repository)
		db.EXPECT().Info().Times(2).Return(dbInfo)
		dbInfo.EXPECT().Update(gomock.Any()).Times(2).Return(database.Info{}, nil)
		repoMock.EXPECT().Rules().Times(numCommit).Return(rulesData)
		rulesData.EXPECT().Process(gomock.Any()).Times(numCommit).Return([]report.Finding{}, nil)

		w := make(chan bool, 1)
		s, err := repo.Scan(repoMock, "false", db, w)
		time.Sleep(2 * time.Second)

		assert.Nil(t, err)
		assert.Equal(t, status.QUEUED.String(), s)
	})

	t.Run("Error: clone fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		rulesData := ruleMocks.NewMockInterface(ctrl)

		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().Repo().Return(nil)
		repoMock.EXPECT().Clone().Return(errors.New("clone error"))
		repoMock.EXPECT().Rules().Return(rulesData)
		rulesData.EXPECT().RuleSet().Return("1234")

		w := make(chan bool, 1)
		s, err := repo.Scan(repoMock, "false", db, w)

		assert.NotNil(t, err)
		assert.Equal(t, reqError.URL_ERROR.String(), s)
	})

	t.Run("Error: get head error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		rulesData := ruleMocks.NewMockInterface(ctrl)

		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().Repo().Return(&gitLib.Repository{})
		repoMock.EXPECT().GetHeadCommitHash().Return("", errors.New("head error"))
		repoMock.EXPECT().Rules().Return(rulesData)
		rulesData.EXPECT().RuleSet().Return("1234")

		w := make(chan bool, 1)
		s, err := repo.Scan(repoMock, "false", db, w)

		assert.NotNil(t, err)
		assert.Equal(t, reqError.HEAD_ERROR.String(), s)
	})

	t.Run("Error: find db error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		rulesData := ruleMocks.NewMockInterface(ctrl)

		db.EXPECT().Info().Times(1).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{}, errors.New("db error"))
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().Repo().Return(&gitLib.Repository{})
		repoMock.EXPECT().GetHeadCommitHash().Return(commit, nil)
		repoMock.EXPECT().Rules().Return(rulesData)
		rulesData.EXPECT().RuleSet().Return("1234")

		w := make(chan bool, 1)
		s, err := repo.Scan(repoMock, "false", db, w)

		assert.NotNil(t, err)
		assert.Equal(t, reqError.INTERNAL_ERROR.String(), s)
	})

	t.Run("Error: create db error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		rulesData := ruleMocks.NewMockInterface(ctrl)

		db.EXPECT().Info().Times(2).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{}, gorm.ErrRecordNotFound)
		dbInfo.EXPECT().Create(gomock.Any()).Return(database.Info{}, errors.New("create db error"))
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().GetHeadCommitHash().Return(commit, nil)
		repoMock.EXPECT().Repo().Return(&gitLib.Repository{})
		repoMock.EXPECT().Rules().Return(rulesData)
		rulesData.EXPECT().RuleSet().Return("1234")

		w := make(chan bool, 1)
		s, err := repo.Scan(repoMock, "false", db, w)
		assert.NotNil(t, err)
		assert.Equal(t, reqError.INTERNAL_ERROR.String(), s)
	})

}

func getNumberOfCommit(repository *gitLib.Repository) (int, error) {
	ref, err := repository.Head()
	if err != nil {
		return 0, err
	}
	cIter, err := repository.Log(&gitLib.LogOptions{From: ref.Hash()})
	count := 0
	cIter.ForEach(func(c *object.Commit) error {
		count++
		return nil
	})

	return count, err
}
