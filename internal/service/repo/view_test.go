package repo_test

import (
	"encoding/json"
	"github.com/gitscan/constants/status"
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/report"
	"github.com/gitscan/internal/service/repo"
	dbMocks "github.com/gitscan/mocks/db"
	repoMocks "github.com/gitscan/mocks/repo"
	ruleMocks "github.com/gitscan/mocks/rule"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepo_View(t *testing.T) {
	url := "github.com/test"
	name := "test"
	t.Run("success: queue", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)

		db.EXPECT().Info().Times(1).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{Status: status.QUEUED.String()}, nil)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().URL().Return(url)

		info, err := repo.ViewReport(repoMock, "true", db)
		assert.Nil(t, err)
		assert.Equal(t, status.QUEUED.String(), info.Status)
	})

	t.Run("success: in progress", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)

		db.EXPECT().Info().Times(1).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{Status: status.INPROGRESS.String()}, nil)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().URL().Return(url)

		info, err := repo.ViewReport(repoMock, "true", db)
		assert.Nil(t, err)
		assert.Equal(t, status.INPROGRESS.String(), info.Status)
	})

	t.Run("success: failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)

		db.EXPECT().Info().Times(1).Return(dbInfo)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{Status: status.FAILED.String()}, nil)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().URL().Return(url)

		info, err := repo.ViewReport(repoMock, "true", db)
		assert.Nil(t, err)
		assert.Equal(t, status.FAILED.String(), info.Status)
	})

	t.Run("success: view all commit", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		dbFinding := dbMocks.NewMockFindingInterface(ctrl)
		dbLocation := dbMocks.NewMockLocationInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		ruleMock := ruleMocks.NewMockInterface(ctrl)

		infoID := "info_1234"

		db.EXPECT().Info().Times(1).Return(dbInfo)
		db.EXPECT().Finding().Times(1).Return(dbFinding)
		db.EXPECT().Location().Times(2).Return(dbLocation)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{
			Model: database.Model{
				ID: infoID,
			},
			Name:        "test",
			URL:         "github.com/test",
			Status:      status.SUCCESS.String(),
			Commit:      "1",
			Description: "",
			EnqueuedAt:  time.Time{},
			StartedAt:   nil,
			FinishedAt:  nil,
		}, nil)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Rules().Times(2).Return(ruleMock)
		dbFinding.EXPECT().FindByInfoID(infoID).Return([]database.Finding{
			{
				Model: database.Model{
					ID: "finding_1",
				},
				InfoID: infoID,
				Type:   "sasl",
				RuleID: "G401",
				Commit: "1",
			}, {
				Model: database.Model{
					ID: "finding_2",
				},
				InfoID: "info_1234",
				Type:   "sasl",
				RuleID: "G401",
				Commit: "2",
			},
		}, nil)
		dbLocation.EXPECT().FindByFindingID("finding_1").Return([]database.Location{
			{
				Model: database.Model{
					ID: "location_1",
				},
				FindingID: "finding_1",
				Path:      "test/test1.go",
				Lines:     "13,15,38",
				Status:    status.ACTIVE.String(),
			},
		}, nil)
		dbLocation.EXPECT().FindByFindingID("finding_2").Return([]database.Location{
			{
				Model: database.Model{
					ID: "location_2",
				},
				FindingID: "finding_2",
				Path:      "test/test2.go",
				Lines:     "20,28,38",
				Status:    status.ACTIVE.String(),
			},
			{
				Model: database.Model{
					ID: "location_3",
				},
				FindingID: "finding_2",
				Path:      "test/test1.go",
				Lines:     "1,2",
				Status:    status.ACTIVE.String(),
			},
		}, nil)
		ruleMock.EXPECT().GetMetaData(gomock.Any()).Times(2).Return(report.Metadata{
			Description: "test",
			Severity:    "HIGH",
		})

		info, err := repo.ViewReport(repoMock, "true", db)
		infoJson, _ := json.Marshal(info)
		expected := `{"Name":"test","URL":"github.com/test","Status":"SUCCESS","Description":"","EnqueuedAt":"0001-01-01T00:00:00Z","StartedAt":null,"FinishedAt":null,"CreatedAt":"0001-01-01T00:00:00Z","Findings":[{"Type":"sasl","RuleId":"G401","Location":[{"Path":"test/test1.go","Positions":{"Begin":[{"Line":"13"},{"Line":"15"},{"Line":"38"}]}}],"Commit":"1","Metadata":{"Description":"test","Severity":"HIGH"}},{"Type":"sasl","RuleId":"G401","Location":[{"Path":"test/test2.go","Positions":{"Begin":[{"Line":"20"},{"Line":"28"},{"Line":"38"}]}},{"Path":"test/test1.go","Positions":{"Begin":[{"Line":"1"},{"Line":"2"}]}}],"Commit":"2","Metadata":{"Description":"test","Severity":"HIGH"}}]}`
		assert.Nil(t, err)
		assert.Equal(t, status.SUCCESS.String(), info.Status)
		assert.Equal(t, expected, string(infoJson))
	})

	t.Run("success: view latest commit", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbInfo := dbMocks.NewMockInfoInterface(ctrl)
		dbFinding := dbMocks.NewMockFindingInterface(ctrl)
		dbLocation := dbMocks.NewMockLocationInterface(ctrl)
		repoMock := repoMocks.NewMockInterface(ctrl)
		ruleMock := ruleMocks.NewMockInterface(ctrl)

		infoID := "info_1234"

		db.EXPECT().Info().Times(1).Return(dbInfo)
		db.EXPECT().Finding().Times(1).Return(dbFinding)
		db.EXPECT().Location().Times(1).Return(dbLocation)
		dbInfo.EXPECT().Find(database.Info{URL: url}).Return(database.Info{
			Model: database.Model{
				ID: infoID,
			},
			Name:        "test",
			URL:         "github.com/test",
			Status:      status.SUCCESS.String(),
			Commit:      "1",
			Description: "",
			EnqueuedAt:  time.Time{},
			StartedAt:   nil,
			FinishedAt:  nil,
		}, nil)
		repoMock.EXPECT().Name().Return(name)
		repoMock.EXPECT().URL().Return(url)
		repoMock.EXPECT().Rules().Times(1).Return(ruleMock)
		dbFinding.EXPECT().FindByInfoIDAndCommit(infoID, "1").Return(database.Finding{
			Model: database.Model{
				ID: "finding_1",
			},
			InfoID: infoID,
			Type:   "sasl",
			RuleID: "G401",
			Commit: "1",
		}, nil)
		dbLocation.EXPECT().FindByFindingID("finding_1").Return([]database.Location{
			{
				Model: database.Model{
					ID: "location_1",
				},
				FindingID: "finding_1",
				Path:      "test/test1.go",
				Lines:     "13,15,38",
				Status:    status.ACTIVE.String(),
			},
		}, nil)
		ruleMock.EXPECT().GetMetaData(gomock.Any()).Times(1).Return(report.Metadata{
			Description: "test",
			Severity:    "HIGH",
		})

		info, err := repo.ViewReport(repoMock, "false", db)
		infoJson, _ := json.Marshal(info)
		expected := `{"Name":"test","URL":"github.com/test","Status":"SUCCESS","Description":"","EnqueuedAt":"0001-01-01T00:00:00Z","StartedAt":null,"FinishedAt":null,"CreatedAt":"0001-01-01T00:00:00Z","Findings":[{"Type":"sasl","RuleId":"G401","Location":[{"Path":"test/test1.go","Positions":{"Begin":[{"Line":"13"},{"Line":"15"},{"Line":"38"}]}}],"Commit":"1","Metadata":{"Description":"test","Severity":"HIGH"}}]}`
		assert.Nil(t, err)
		assert.Equal(t, status.SUCCESS.String(), info.Status)
		assert.Equal(t, expected, string(infoJson))
	})
}
