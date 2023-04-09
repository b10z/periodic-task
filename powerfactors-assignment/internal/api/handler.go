package api

import (
	"encoding/json"
	"net/http"
	"powerfactors/assignment/internal/service"
	"time"
)

const timeFormat = "20060102T150405Z"

type TimestampHandler struct {
	taskService service.TaskServiceInt
}

type TimestampHandlerInt interface {
	GetTimestamp(w http.ResponseWriter, r *http.Request)
}

func NewTimestampHandler(ts service.TaskServiceInt) *TimestampHandler {
	return &TimestampHandler{
		taskService: ts,
	}
}

type errorJSON struct {
	Status string `json:"status"`
	Desc   string `json:"desc"`
}

func (th *TimestampHandler) GetTimestamp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(r.URL.Query()) < 4 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errorJSON{
			Status: "error",
			Desc:   NewAPIError("invalid number of parameters").Error(),
		})
		return
	}

	period := r.URL.Query().Get("period")
	if period == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errorJSON{
			Status: "error",
			Desc:   NewAPIError("invalid period parameter").Error(),
		})
		return
	}

	timezone, err := time.LoadLocation(r.URL.Query().Get("tz"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errorJSON{
			Status: "error",
			Desc:   NewAPIError("invalid tz parameter").Error(),
		})
		return
	}

	timestamp1, err := time.Parse(timeFormat, r.URL.Query().Get("t1"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errorJSON{
			Status: "error",
			Desc:   NewAPIError("invalid t1 parameter").Error(),
		})
		return
	}

	timestamp2, err := time.Parse(timeFormat, r.URL.Query().Get("t2"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errorJSON{
			Status: "error",
			Desc:   NewAPIError("invalid t2 parameter").Error(),
		})
		return
	}
	generatedTimestamps, err := th.taskService.PeriodicTaskService(period, timezone, timestamp1, timestamp2)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errorJSON{
			Status: "error",
			Desc:   NewAPIError("error while generating the timestamps").Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(generatedTimestamps)
}
