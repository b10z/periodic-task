package service

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestNewTaskService(t *testing.T) {
	type args struct {
		logger *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want *TaskService
	}{
		{
			name: "NewTaskService test",
			args: args{
				logger: zap.NewNop(),
			},
			want: &TaskService{
				logger: zap.NewNop(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewTaskService(tt.args.logger), "NewTaskService(%v)", tt.args.logger)
		})
	}
}

func TestTaskService_GenerateTimestampService(t *testing.T) {
	type fields struct {
		logger *zap.Logger
	}
	type args struct {
		period    string
		timezone  *time.Location
		startDate time.Time
		endDate   time.Time
	}
	type testData struct {
		timezone string
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
			name: "GenerateTimestampService successful test with period of 1d",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1d",
				timezone:  nil,
				startDate: time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 02, 16, 20, 46, 03, 0, time.UTC),
			},
			testData: testData{
				timezone: "Europe/Athens",
			},
			want:    []string{"20210214T204603Z", "20210215T204603Z"},
			wantErr: nil,
		},
		{
			name: "GenerateTimestampService successful test with period of 1mo",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1mo",
				timezone:  nil,
				startDate: time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 03, 16, 20, 46, 03, 0, time.UTC),
			},
			testData: testData{
				timezone: "Europe/Athens",
			},
			want:    []string{"20210214T204603Z", "20210314T204603Z"},
			wantErr: nil,
		},
		{
			name: "GenerateTimestampService successful test with period of 1y",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1y",
				timezone:  nil,
				startDate: time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2022, 02, 16, 20, 46, 03, 0, time.UTC),
			},
			testData: testData{
				timezone: "Europe/Athens",
			},
			want:    []string{"20210214T204603Z", "20220214T204603Z"},
			wantErr: nil,
		},
		{
			name: "GenerateTimestampService successful test with period of 1h",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1h",
				timezone:  nil,
				startDate: time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 02, 15, 01, 46, 03, 0, time.UTC),
			},
			testData: testData{
				timezone: "Europe/Athens",
			},
			want:    []string{"20210214T204603Z", "20210214T214603Z", "20210214T224603Z", "20210214T234603Z", "20210215T004603Z"},
			wantErr: nil,
		},
		{
			name: "GenerateTimestampService failed test startDate after the endDate",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1d",
				timezone:  nil,
				startDate: time.Date(2021, 02, 17, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 02, 16, 20, 46, 03, 0, time.UTC),
			},
			testData: testData{
				timezone: "Europe/Athens",
			},
			want:    nil,
			wantErr: NewServiceError("t1 should be before t2"),
		},
		{
			name: "GenerateTimestampService failed test with invalid period argument",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1sec",
				timezone:  nil,
				startDate: time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 02, 16, 20, 46, 03, 0, time.UTC),
			},
			testData: testData{
				timezone: "Europe/Athens",
			},
			want:    nil,
			wantErr: NewServiceError("invalid period parameter"),
		},
		{
			name: "GenerateTimestampService failed test with period of 1d and dates out of range",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1d",
				timezone:  nil,
				startDate: time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 02, 14, 21, 46, 03, 0, time.UTC),
			},
			testData: testData{
				timezone: "Europe/Athens",
			},
			want:    nil,
			wantErr: NewServiceError("period parameter out of range"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TaskService{
				logger: tt.fields.logger,
			}
			tz, _ := time.LoadLocation(tt.testData.timezone)
			got, err := ts.GenerateTimestampService(tt.args.period, tz, tt.args.startDate, tt.args.endDate)

			assert.Equalf(t, tt.want, got, "GenerateTimestampService(%v, %v, %v, %v)", tt.args.period, tt.args.timezone, tt.args.startDate, tt.args.endDate)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
