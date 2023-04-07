package api

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAPIError(t *testing.T) {
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
			wantErr: errors.New("Test Error Text"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewAPIError(tt.args.text)
			assert.Equal(t, tt.wantErr.Error(), err.Error())
		})
	}
}

func Test_apiError_Error(t *testing.T) {
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
			re := &apiError{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, re.Error(), "Error()")
		})
	}
}
