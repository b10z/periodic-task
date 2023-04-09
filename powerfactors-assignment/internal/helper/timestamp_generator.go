package helper

import (
	"go.uber.org/zap"
	"time"
)

const timeFormat = "20060102T150405Z"

type TimestampGenerator struct {
	logger *zap.Logger
}

type TimestampGeneratorInt interface {
	GenerateTimestamps(period string, startDate, endDate time.Time) ([]string, error)
}

func NewTimestampGenerator(logger *zap.Logger) *TimestampGenerator {
	return &TimestampGenerator{
		logger: logger,
	}
}

type periodicIncrease struct {
	year    int
	month   int
	week    int
	day     int
	hour    int
	minutes int
	seconds int
}

func (tg TimestampGenerator) GenerateTimestamps(period string, startDate, endDate time.Time) ([]string, error) {
	startDate, duration, err := getDurationFromPeriod(period, startDate, endDate)
	if err != nil {
		tg.logger.Error("error while getting duration from period")
		return nil, err
	}

	var generatedTimestamps []string
	for startDate.Before(endDate) {
		generatedTimestamps = append(generatedTimestamps, startDate.UTC().Format(timeFormat))
		startDate = startDate.AddDate(duration.year, duration.month, duration.day)
		startDate = startDate.Add(time.Duration(duration.hour)*time.Hour +
			time.Duration(duration.minutes)*time.Minute +
			time.Duration(duration.seconds)*time.Second)
	}
	return generatedTimestamps, nil
}

// In getDurationFromPeriod new periods can easily be declared. Just add a new case inside the switch, and return the new startingTimestamp and the periodic increase (the step)
// as done for the other periods. Then, the algorithm will handle accordingly any increment given.
func getDurationFromPeriod(period string, startDate, endDate time.Time) (time.Time, *periodicIncrease, error) {
	switch period {
	case "1h":
		startingTimestamp := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), startDate.Hour()+1, 0, 0, 0, startDate.Location())
		periodicIncrement := &periodicIncrease{
			year:    0,
			month:   0,
			week:    0,
			day:     0,
			hour:    1,
			minutes: 0,
			seconds: 0,
		}
		if startingTimestamp.After(endDate) {
			return time.Time{}, &periodicIncrease{}, NewHelperError("period parameter out of range")
		}
		return startingTimestamp, periodicIncrement, nil
	case "1d":
		timeTest := time.Date(startDate.Year(), startDate.Month(), startDate.Day()+1, 0, 0, 0, 0, startDate.Location())
		periodicIncrement := &periodicIncrease{
			year:    0,
			month:   0,
			week:    0,
			day:     1,
			hour:    0,
			minutes: 0,
			seconds: 0,
		}
		if timeTest.After(endDate) {
			return time.Time{}, &periodicIncrease{}, NewHelperError("period parameter out of range")
		}
		return timeTest, periodicIncrement, nil
	case "1mo":
		startingTimestamp := time.Date(startDate.Year(), startDate.Month()+1, 1, 0, 0, 0, 0, startDate.Location())
		periodicIncrement := &periodicIncrease{
			year:    0,
			month:   1,
			week:    0,
			day:     0,
			hour:    0,
			minutes: 0,
			seconds: 0,
		}
		if startingTimestamp.After(endDate) {
			return time.Time{}, &periodicIncrease{}, NewHelperError("period parameter out of range")
		}
		return startingTimestamp, periodicIncrement, nil
	case "1y":
		startingTimestamp := time.Date(startDate.Year()+1, 1, 1, 0, 0, 0, 0, startDate.Location())
		periodicIncrement := &periodicIncrease{
			year:    1,
			month:   0,
			week:    0,
			day:     0,
			hour:    0,
			minutes: 0,
			seconds: 0,
		}
		if startingTimestamp.After(endDate) {
			return time.Time{}, &periodicIncrease{}, NewHelperError("period parameter out of range")
		}
		return startingTimestamp, periodicIncrement, nil
	default:
		return time.Time{}, &periodicIncrease{}, NewHelperError("invalid period parameter")
	}
}
