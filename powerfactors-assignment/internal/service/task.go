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
	GenerateTimestampService(period string, timezone *time.Location, startDate, endDate time.Time) ([]string, error)
}

func NewTaskService(logger *zap.Logger) *TaskService {
	return &TaskService{
		logger: logger,
	}
}

func (ts *TaskService) GenerateTimestampService(period string, timezone *time.Location, startDate, endDate time.Time) ([]string, error) {
	timestampInLocation1 := startDate.In(timezone)
	timestampInLocation2 := endDate.In(timezone)

	if timestampInLocation1.After(timestampInLocation2) {
		ts.logger.Error("startDate is after the endDate")
		return nil, NewServiceError("t1 should be before t2")
	}

	duration, err := getDurationFromPeriod(period, startDate, endDate)
	if err != nil {
		ts.logger.Error("error while getting duration from period")
		return nil, err
	}

	var generatedTimestamps []string
	for startDate.Before(endDate) {
		generatedTimestamps = append(generatedTimestamps, startDate.Format(timeFormat))
		startDate = startDate.Add(duration)
	}

	return generatedTimestamps, nil
}

// In getDurationFromPeriod new periods can easily be declared. Just add a new case inside the switch, and return the duration from the startDate
// as done for the other periods. Then, the algorithm will handle accordingly any duration given.
func getDurationFromPeriod(period string, startDate, endDate time.Time) (time.Duration, error) {
	switch period {
	case "1h":
		timeDuration := startDate.Add(time.Hour)
		if timeDuration.After(endDate) {
			return 0, NewServiceError("period parameter out of range")
		}
		return timeDuration.Sub(startDate), nil
	case "1d":
		timeDuration := startDate.AddDate(0, 0, 1)
		if timeDuration.After(endDate) {
			return 0, NewServiceError("period parameter out of range")
		}
		return timeDuration.Sub(startDate), nil
	case "1mo":
		timeDuration := startDate.AddDate(0, 1, 0)
		if timeDuration.After(endDate) {
			return 0, NewServiceError("period parameter out of range")
		}
		return timeDuration.Sub(startDate), nil
	case "1y":
		timeDuration := startDate.AddDate(1, 0, 0)
		if timeDuration.After(endDate) {
			return 0, NewServiceError("period parameter out of range")
		}
		return timeDuration.Sub(startDate), nil
	default:
		return 0, NewServiceError("invalid period parameter")
	}
}
