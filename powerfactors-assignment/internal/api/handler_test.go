package api

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"powerfactors/assignment/internal/service"
	mock_service "powerfactors/assignment/mocks/internal_/service"
	"testing"
	"time"
)

func TestNewTimestampHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ts service.TaskServiceInt
	}
	tests := []struct {
		name string
		args args
		want *TimestampHandler
	}{
		{
			name: "NewTimestampHandler test",
			args: args{
				ts: mock_service.NewMockTaskServiceInt(ctrl),
			},
			want: &TimestampHandler{
				taskService: mock_service.NewMockTaskServiceInt(ctrl),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTimestampHandler(tt.args.ts)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestTimestampHandler_GetTimestamp(t *testing.T) {
	ctrl := gomock.NewController(t)
	type fields struct {
		taskService   service.TaskServiceInt
		serviceResult []string
		serviceError  error
		startDate     time.Time
		endDate       time.Time
		timezone      string
		period        string
		statusCode    int
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wWant  http.ResponseWriter
	}{
		{
			name: "Successful request (200-OK)",
			fields: fields{
				taskService:   mock_service.NewMockTaskServiceInt(ctrl),
				serviceResult: []string{"20210214T204603Z"},
				startDate:     time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:       time.Date(2021, 02, 15, 20, 46, 03, 0, time.UTC),
				period:        "1d",
				timezone:      "Europe/Athens",
				statusCode:    200,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/ptlist?tz=Europe/Athens&t1=20210214T204603Z&t2=20210215T204603Z&period=1d", nil),
			},
			wWant: &httptest.ResponseRecorder{
				HeaderMap: http.Header{"Content-Type": []string{"application/json"}},
				Body: bytes.NewBufferString(`["20210214T204603Z"]
`),
				Flushed: false,
			},
		},
		{
			name: "Failed request with missing parameter (400-BadRequest)",
			fields: fields{
				taskService:  mock_service.NewMockTaskServiceInt(ctrl),
				serviceError: errors.New("testError"),
				startDate:    time.Now(),
				endDate:      time.Date(2021, 02, 15, 20, 46, 03, 0, time.UTC),
				period:       "1d",
				timezone:     "Europe/Athens",
				statusCode:   400,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/ptlist?tz=Europe/Athens&t2=20210215T204603Z&period=1d", nil),
			},
			wWant: &httptest.ResponseRecorder{
				HeaderMap: http.Header{"Content-Type": []string{"application/json"}},
				Body: bytes.NewBufferString(`{"status":"error","desc":"invalid number of parameters"}
`),
				Flushed: false,
			},
		},
		{
			name: "Failed request with invalid parameter (400-BadRequest)",
			fields: fields{
				taskService:   mock_service.NewMockTaskServiceInt(ctrl),
				serviceResult: []string{"20210214T204603Z"},
				startDate:     time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:       time.Date(2021, 02, 15, 20, 46, 03, 0, time.UTC),
				period:        "1d",
				timezone:      "Europe/Athens",
				statusCode:    400,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/ptlist?tz=Europe/Athens&t1=20210214T204603Z&t2=2021021522T204603Z&period=1d", nil),
			},
			wWant: &httptest.ResponseRecorder{
				HeaderMap: http.Header{"Content-Type": []string{"application/json"}},
				Body: bytes.NewBufferString(`{"status":"error","desc":"invalid t2 parameter"}
`),
				Flushed: false,
			},
		},
		{
			name: "Failed request with service error (400-BadRequest)",
			fields: fields{
				taskService:  mock_service.NewMockTaskServiceInt(ctrl),
				serviceError: errors.New("testError"),
				startDate:    time.Date(2021, 02, 14, 20, 46, 03, 0, time.UTC),
				endDate:      time.Date(2021, 02, 15, 20, 46, 03, 0, time.UTC),
				period:       "1d",
				timezone:     "Europe/Athens",
				statusCode:   400,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/ptlist?tz=Europe/Athens&t1=20210214T204603Z&t2=20210215T204603Z&period=1d", nil),
			},
			wWant: &httptest.ResponseRecorder{
				HeaderMap: http.Header{"Content-Type": []string{"application/json"}},
				Body: bytes.NewBufferString(`{"status":"error","desc":"error while generating the timestamps"}
`),
				Flushed: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := &TimestampHandler{
				taskService: tt.fields.taskService,
			}
			tz, _ := time.LoadLocation(tt.fields.timezone)
			tt.wWant.WriteHeader(tt.fields.statusCode)
			th.taskService.(*mock_service.MockTaskServiceInt).EXPECT().GenerateTimestampService(tt.fields.period, tz, tt.fields.startDate, tt.fields.endDate).Return(tt.fields.serviceResult, tt.fields.serviceError).AnyTimes()
			th.GetTimestamp(tt.args.w, tt.args.r)
			assert.Equal(t, tt.wWant.Header(), tt.args.w.Header())
			assert.EqualValues(t, tt.wWant, tt.args.w)
		})
	}
}
