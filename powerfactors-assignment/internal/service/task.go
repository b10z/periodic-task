package service

import (
	"go.uber.org/zap"
	"time"
)

const timeFormat = "20060102T150405Z"

type TaskService struct {
	logger *zap.Logger
}

type TaskServiceInt interface {
	GetTimestamp(period string, timezone time.Location, timestamp1, timestamp2 time.Time) ([]string, error)
}

func NewTaskService(logger *zap.Logger) *TaskService {
	return &TaskService{
		logger: logger,
	}
}

func (ts *TaskService) GetTimestamps(period string, timezone *time.Location, timestamp1, timestamp2 time.Time) ([]string, error) {
	timestampInLocation1 := timestamp1.In(timezone)
	timestampInLocation2 := timestamp2.In(timezone)

	if timestampInLocation1.After(timestampInLocation2) {
		return nil, NewServiceError("t1 should be before t2")
	}

	var duration time.Duration

	switch period {
	case "1h":
		duration = time.Hour
		break
	case "1d":
		duration = timestamp1.AddDate(0, 0, 1).Sub(timestamp1)
		break
	case "1mo":
		duration = timestamp1.AddDate(0, 1, 0).Sub(timestamp1)
		break
	case "1y":
		duration = timestamp1.AddDate(1, 0, 0).Sub(timestamp1)
		break
	default:
		ts.logger.Error("parameter period is unknown")
		return nil, NewServiceError("invalid period parameter")
	}

	var generatedTimestamps []string
	for timestamp1.Before(timestamp2) {
		generatedTimestamps = append(generatedTimestamps, timestamp1.Format(timeFormat))
		timestamp1 = timestamp1.Add(duration)
	}

	return generatedTimestamps, nil
}
