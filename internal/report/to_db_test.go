package report_test

import (
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/report"
	dbMocks "github.com/gitscan/mocks/db"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"testing"
)

func TestSaveFindings(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbMocks.NewMockDB(ctrl)
		dbFindings := dbMocks.NewMockFindingInterface(ctrl)
		dbLocation := dbMocks.NewMockLocationInterface(ctrl)

		db.EXPECT().Finding().Times(5).Return(dbFindings)
		db.EXPECT().Location().Times(1).Return(dbLocation)
		dbFindings.EXPECT().FindByInfoIDAndCommit(gomock.Any(), gomock.Any()).Times(2).Return(database.Finding{}, gorm.ErrRecordNotFound)
		dbFindings.EXPECT().FindByInfoIDAndCommit(gomock.Any(), gomock.Any()).Times(2).Return(database.Finding{}, nil)
		dbFindings.EXPECT().Creates(gomock.Any()).Return([]database.Finding{}, nil)
		dbLocation.EXPECT().Creates(gomock.Any()).Return([]database.Location{}, nil)

		report.SaveFindings("123", "commit", []report.Finding{
			{
				Type:   "",
				RuleId: "",
				Location: []report.Location{{
					Path: "",
					Positions: report.Position{
						Begin: []report.Begin{{Line: "1"}, {Line: "2"}},
					},
				}},
				Commit:   "",
				Metadata: report.Metadata{},
			},
			{
				Type:   "",
				RuleId: "",
				Location: []report.Location{{
					Path: "",
					Positions: report.Position{
						Begin: []report.Begin{},
					},
				}},
				Commit:   "",
				Metadata: report.Metadata{},
			},
		}, db)
	})
}
