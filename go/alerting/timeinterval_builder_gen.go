// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeInterval] = (*TimeIntervalBuilder)(nil)

type TimeIntervalBuilder struct {
	internal *TimeInterval
	errors   map[string]cog.BuildErrors
}

func NewTimeIntervalBuilder() *TimeIntervalBuilder {
	resource := NewTimeInterval()
	builder := &TimeIntervalBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *TimeIntervalBuilder) Build() (TimeInterval, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimeInterval{}, err
	}

	return *builder.internal, nil
}

func (builder *TimeIntervalBuilder) Times(times []cog.Builder[TimeRange]) *TimeIntervalBuilder {
	timesResources := make([]TimeRange, 0, len(times))
	for _, r1 := range times {
		timesDepth1, err := r1.Build()
		if err != nil {
			builder.errors["times"] = err.(cog.BuildErrors)
			return builder
		}
		timesResources = append(timesResources, timesDepth1)
	}
	builder.internal.Times = timesResources

	return builder
}

func (builder *TimeIntervalBuilder) Weekdays(weekdays []cog.Builder[WeekdayRange]) *TimeIntervalBuilder {
	weekdaysResources := make([]WeekdayRange, 0, len(weekdays))
	for _, r1 := range weekdays {
		weekdaysDepth1, err := r1.Build()
		if err != nil {
			builder.errors["weekdays"] = err.(cog.BuildErrors)
			return builder
		}
		weekdaysResources = append(weekdaysResources, weekdaysDepth1)
	}
	builder.internal.Weekdays = weekdaysResources

	return builder
}

func (builder *TimeIntervalBuilder) DaysOfMonth(daysOfMonth []cog.Builder[DayOfMonthRange]) *TimeIntervalBuilder {
	daysOfMonthResources := make([]DayOfMonthRange, 0, len(daysOfMonth))
	for _, r1 := range daysOfMonth {
		daysOfMonthDepth1, err := r1.Build()
		if err != nil {
			builder.errors["days_of_month"] = err.(cog.BuildErrors)
			return builder
		}
		daysOfMonthResources = append(daysOfMonthResources, daysOfMonthDepth1)
	}
	builder.internal.DaysOfMonth = daysOfMonthResources

	return builder
}

func (builder *TimeIntervalBuilder) Months(months []cog.Builder[MonthRange]) *TimeIntervalBuilder {
	monthsResources := make([]MonthRange, 0, len(months))
	for _, r1 := range months {
		monthsDepth1, err := r1.Build()
		if err != nil {
			builder.errors["months"] = err.(cog.BuildErrors)
			return builder
		}
		monthsResources = append(monthsResources, monthsDepth1)
	}
	builder.internal.Months = monthsResources

	return builder
}

func (builder *TimeIntervalBuilder) Years(years []cog.Builder[YearRange]) *TimeIntervalBuilder {
	yearsResources := make([]YearRange, 0, len(years))
	for _, r1 := range years {
		yearsDepth1, err := r1.Build()
		if err != nil {
			builder.errors["years"] = err.(cog.BuildErrors)
			return builder
		}
		yearsResources = append(yearsResources, yearsDepth1)
	}
	builder.internal.Years = yearsResources

	return builder
}

func (builder *TimeIntervalBuilder) Location(location Location) *TimeIntervalBuilder {
	builder.internal.Location = &location

	return builder
}
