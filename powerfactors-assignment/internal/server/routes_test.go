package server

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"powerfactors/assignment/internal/api"
	mock_api "powerfactors/assignment/mocks/internal_/api"
	"reflect"
	"testing"
)

func TestNewListener(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		rt        *mux.Router
		log       *zap.Logger
		server    *http.Server
		tmHandler api.TimestampHandlerInt
	}
	tests := []struct {
		name string
		args args
		want Listener
	}{
		{
			name: "Basic NewListener test",
			args: args{
				rt:  mux.NewRouter(),
				log: zap.NewNop(),
				server: &http.Server{
					Addr: "testAddress",
				},
				tmHandler: mock_api.NewMockTimestampHandlerInt(ctrl),
			},
			want: Listener{
				router: mux.NewRouter(),
				logger: zap.NewNop(),
				httpServer: &http.Server{
					Addr: "testAddress",
				},
				timestampHandler: mock_api.NewMockTimestampHandlerInt(ctrl),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListener(tt.args.rt, tt.args.log, tt.args.server, tt.args.tmHandler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
