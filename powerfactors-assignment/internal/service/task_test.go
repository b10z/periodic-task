package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"powerfactors/assignment/internal/helper"
	mock_helper "powerfactors/assignment/mocks/internal_/helper"
	"testing"
	"time"
)

func TestNewTaskService(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		logger             *zap.Logger
		timestampGenerator helper.TimestampGeneratorInt
	}
	tests := []struct {
		name string
		args args
		want *TaskService
	}{
		{
			name: "NewTaskService test",
			args: args{
				logger:             zap.NewNop(),
				timestampGenerator: mock_helper.NewMockTimestampGeneratorInt(ctrl),
			},
			want: &TaskService{
				logger:             zap.NewNop(),
				timestampGenerator: mock_helper.NewMockTimestampGeneratorInt(ctrl),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewTaskService(tt.args.logger, tt.args.timestampGenerator), "NewTaskService(%v)", tt.args.logger)
		})
	}
}

func TestTaskService_PeriodicTaskService(t *testing.T) {
	ctrl := gomock.NewController(t)
	type fields struct {
		logger             *zap.Logger
		timestampGenerator helper.TimestampGeneratorInt
	}
	type args struct {
		period    string
		timezone  *time.Location
		startDate time.Time
		endDate   time.Time
	}
	type testData struct {
		timezone   string
		period     string
		mockResult []string
		mockError  error
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		testData testData
		want     []string
		wantErr  error
	}{
		{
			name: "Successful timestamp generation",
			fields: fields{
				logger:             zap.NewNop(),
				timestampGenerator: mock_helper.NewMockTimestampGeneratorInt(ctrl),
			},
			args: args{
				period:    "1d",
				timezone:  nil,
				startDate: time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 02, 16, 20, 46, 03, 0, time.UTC),
			},
			testData: testData{
				timezone:   "Europe/Athens",
				period:     "1d",
				mockResult: []string{"20210214T000000Z", "20210215T000000Z"},
				mockError:  nil,
			},
			want:    []string{"20210214T000000Z", "20210215T000000Z"},
			wantErr: nil,
		},
		{
			name: "Failed timestamp generation",
			fields: fields{
				logger:             zap.NewNop(),
				timestampGenerator: mock_helper.NewMockTimestampGeneratorInt(ctrl),
			},
			args: args{
				period:    "1day",
				timezone:  nil,
				startDate: time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 02, 16, 20, 46, 03, 0, time.UTC),
			},
			testData: testData{
				timezone:   "Europe/Athens",
				period:     "1day",
				mockResult: nil,
				mockError:  helper.NewHelperError("invalid period parameter"),
			},
			want:    nil,
			wantErr: NewServiceError("error while generating the timestamps"),
		},
		{
			name: "startDate is after the endDate failed test",
			fields: fields{
				logger:             zap.NewNop(),
				timestampGenerator: mock_helper.NewMockTimestampGeneratorInt(ctrl),
			},
			args: args{
				period:    "1day",
				timezone:  nil,
				startDate: time.Date(2021, 06, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 02, 16, 20, 46, 03, 0, time.UTC),
			},
			testData: testData{
				timezone:   "Europe/Athens",
				period:     "1day",
				mockResult: nil,
				mockError:  helper.NewHelperError("invalid period parameter"),
			},
			want:    nil,
			wantErr: NewServiceError("t1 should be before t2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TaskService{
				logger:             tt.fields.logger,
				timestampGenerator: tt.fields.timestampGenerator,
			}
			tz, _ := time.LoadLocation(tt.testData.timezone)
			ts.timestampGenerator.(*mock_helper.MockTimestampGeneratorInt).EXPECT().GenerateTimestamps(tt.testData.period, tt.args.startDate.In(tz), tt.args.endDate).Return(tt.testData.mockResult, tt.testData.mockError).AnyTimes()
			got, err := ts.PeriodicTaskService(tt.args.period, tz, tt.args.startDate, tt.args.endDate)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
