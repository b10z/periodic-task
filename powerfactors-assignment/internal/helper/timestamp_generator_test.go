package helper

import (
	"fmt"
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
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "",
			fields:  fields{},
			args:    args{},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tg := TimestampGenerator{
				logger: tt.fields.logger,
			}
			got, err := tg.GenerateTimestamps(tt.args.period, tt.args.startDate, tt.args.endDate)
			if !tt.wantErr(t, err, fmt.Sprintf("GenerateTimestamps(%v, %v, %v)", tt.args.period, tt.args.startDate, tt.args.endDate)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GenerateTimestamps(%v, %v, %v)", tt.args.period, tt.args.startDate, tt.args.endDate)
		})
	}
}

func Test_getDurationFromPeriod(t *testing.T) {
	type args struct {
		period    string
		startDate time.Time
		endDate   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		want1   *periodicIncrease
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getDurationFromPeriod(tt.args.period, tt.args.startDate, tt.args.endDate)
			if !tt.wantErr(t, err, fmt.Sprintf("getDurationFromPeriod(%v, %v, %v)", tt.args.period, tt.args.startDate, tt.args.endDate)) {
				return
			}
			assert.Equalf(t, tt.want, got, "getDurationFromPeriod(%v, %v, %v)", tt.args.period, tt.args.startDate, tt.args.endDate)
			assert.Equalf(t, tt.want1, got1, "getDurationFromPeriod(%v, %v, %v)", tt.args.period, tt.args.startDate, tt.args.endDate)
		})
	}
}
