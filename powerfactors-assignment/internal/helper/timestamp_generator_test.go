package helper

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestNewTimestampGenerator(t *testing.T) {
	type args struct {
		logger *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want *TimestampGenerator
	}{
		{
			name: "Basic NewTimestampGenerator test",
			args: args{
				logger: zap.NewNop(),
			},
			want: &TimestampGenerator{
				logger: zap.NewNop(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewTimestampGenerator(tt.args.logger), "NewTimestampGenerator(%v)", tt.args.logger)
		})
	}
}

func TestTimestampGenerator_GenerateTimestamps(t *testing.T) {
	type fields struct {
		logger *zap.Logger
	}
	type args struct {
		period    string
		startDate time.Time
		endDate   time.Time
	}
	type testData struct {
		tz string
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
			name: "Successful timestamp generation for 1h",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1h",
				startDate: time.Date(2021, 07, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 07, 15, 12, 34, 56, 0, time.UTC),
			},
			want: []string{
				"20210714T210000Z",
				"20210714T220000Z",
				"20210714T230000Z",
				"20210715T000000Z",
				"20210715T010000Z",
				"20210715T020000Z",
				"20210715T030000Z",
				"20210715T040000Z",
				"20210715T050000Z",
				"20210715T060000Z",
				"20210715T070000Z",
				"20210715T080000Z",
				"20210715T090000Z",
				"20210715T100000Z",
				"20210715T110000Z",
				"20210715T120000Z",
			},
			wantErr: nil,
		},
		{
			name: "Successful timestamp generation for 1d",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1d",
				startDate: time.Date(2021, 10, 10, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 11, 15, 12, 34, 56, 0, time.UTC),
			},
			testData: testData{
				tz: "Europe/Athens",
			},
			want: []string{
				"20211010T210000Z",
				"20211011T210000Z",
				"20211012T210000Z",
				"20211013T210000Z",
				"20211014T210000Z",
				"20211015T210000Z",
				"20211016T210000Z",
				"20211017T210000Z",
				"20211018T210000Z",
				"20211019T210000Z",
				"20211020T210000Z",
				"20211021T210000Z",
				"20211022T210000Z",
				"20211023T210000Z",
				"20211024T210000Z",
				"20211025T210000Z",
				"20211026T210000Z",
				"20211027T210000Z",
				"20211028T210000Z",
				"20211029T210000Z",
				"20211030T210000Z",
				"20211031T220000Z",
				"20211101T220000Z",
				"20211102T220000Z",
				"20211103T220000Z",
				"20211104T220000Z",
				"20211105T220000Z",
				"20211106T220000Z",
				"20211107T220000Z",
				"20211108T220000Z",
				"20211109T220000Z",
				"20211110T220000Z",
				"20211111T220000Z",
				"20211112T220000Z",
				"20211113T220000Z",
				"20211114T220000Z",
			},
			wantErr: nil,
		},
		{
			name: "Successful timestamp generation for 1mo",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1mo",
				startDate: time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 11, 15, 12, 34, 56, 0, time.UTC),
			},
			testData: testData{
				tz: "Europe/Athens",
			},
			want: []string{
				"20210228T220000Z",
				"20210331T210000Z",
				"20210430T210000Z",
				"20210531T210000Z",
				"20210630T210000Z",
				"20210731T210000Z",
				"20210831T210000Z",
				"20210930T210000Z",
				"20211031T220000Z",
			},
			wantErr: nil,
		},
		{
			name: "Successful timestamp generation for 1y",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1y",
				startDate: time.Date(2018, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 11, 15, 12, 34, 56, 0, time.UTC),
			},
			testData: testData{
				tz: "Europe/Athens",
			},
			want: []string{
				"20181231T220000Z",
				"20191231T220000Z",
				"20201231T220000Z",
			},
			wantErr: nil,
		},
		/////
		{
			name: "Failed timestamp generation for Invalid parameter",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "InvalidParameterHere",
				startDate: time.Date(2021, 07, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 07, 15, 12, 34, 56, 0, time.UTC),
			},
			want:    nil,
			wantErr: NewHelperError("invalid period parameter"),
		},
		{
			name: "Failed timestamp generation for 1h",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1h",
				startDate: time.Date(2021, 07, 15, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 07, 15, 12, 34, 56, 0, time.UTC),
			},
			want:    nil,
			wantErr: NewHelperError("period parameter out of range"),
		},
		{
			name: "Failed timestamp generation for 1d",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1d",
				startDate: time.Date(2021, 11, 15, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 11, 15, 12, 34, 56, 0, time.UTC),
			},
			testData: testData{
				tz: "Europe/Athens",
			},
			want:    nil,
			wantErr: NewHelperError("period parameter out of range"),
		},
		{
			name: "Failed timestamp generation for 1mo",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1mo",
				startDate: time.Date(2021, 11, 20, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 11, 15, 12, 34, 56, 0, time.UTC),
			},
			testData: testData{
				tz: "Europe/Athens",
			},
			want:    nil,
			wantErr: NewHelperError("period parameter out of range"),
		},
		{
			name: "Failed timestamp generation for 1y",
			fields: fields{
				logger: zap.NewNop(),
			},
			args: args{
				period:    "1y",
				startDate: time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:   time.Date(2021, 11, 15, 12, 34, 56, 0, time.UTC),
			},
			testData: testData{
				tz: "Europe/Athens",
			},
			want:    nil,
			wantErr: NewHelperError("period parameter out of range"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tg := TimestampGenerator{
				logger: tt.fields.logger,
			}
			tz, _ := time.LoadLocation(tt.testData.tz)
			got, err := tg.GenerateTimestamps(tt.args.period, tt.args.startDate.In(tz), tt.args.endDate)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
