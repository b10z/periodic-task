package helper

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHelperError(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Basic error check",
			args: args{
				text: "Test Error Text",
			},
			wantErr: &helperError{data: "Test Error Text"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewHelperError(tt.args.text)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_helperError_Error(t *testing.T) {
	type fields struct {
		data string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Basic test for error casting",
			fields: fields{
				data: "Test Error Text",
			},
			want: errors.New("Test Error Text").Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &helperError{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, re.Error(), "Error()")
		})
	}
}
