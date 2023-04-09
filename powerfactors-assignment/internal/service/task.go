package service

import (
	"go.uber.org/zap"
	"powerfactors/assignment/internal/helper"
	"time"
)

type TaskService struct {
	logger             *zap.Logger
	timestampGenerator helper.TimestampGeneratorInt
}

type TaskServiceInt interface {
	PeriodicTaskService(period string, timezone *time.Location, startDate, endDate time.Time) ([]string, error)
}

func NewTaskService(logger *zap.Logger, timestampGen helper.TimestampGeneratorInt) *TaskService {
	return &TaskService{
		logger:             logger,
		timestampGenerator: timestampGen,
	}
}

func (ts *TaskService) PeriodicTaskService(period string, timezone *time.Location, startDate, endDate time.Time) ([]string, error) {
	timestampInLocation1 := startDate.In(timezone)
	timestampInLocation2 := endDate.In(timezone)

	if timestampInLocation1.After(timestampInLocation2) {
		ts.logger.Error("startDate is after the endDate")
		return nil, NewServiceError("t1 should be before t2")
	}

	timestamps, err := ts.timestampGenerator.GenerateTimestamps(period, timestampInLocation1, endDate)
	if err != nil {
		ts.logger.Error("error while getting duration from period")
		return nil, NewServiceError("error while generating the timestamps")
	}

	return timestamps, nil
}
