// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;

public class TimeIntervalBuilder implements com.grafana.foundation.cog.Builder<TimeInterval> {
    protected final TimeInterval internal;
    
    public TimeIntervalBuilder() {
        this.internal = new TimeInterval();
    }
    public TimeIntervalBuilder times(com.grafana.foundation.cog.Builder<List<TimeRange>> times) {
        this.internal.times = times.build();
        return this;
    }
    
    public TimeIntervalBuilder weekdays(com.grafana.foundation.cog.Builder<List<WeekdayRange>> weekdays) {
        this.internal.weekdays = weekdays.build();
        return this;
    }
    
    public TimeIntervalBuilder daysOfMonth(com.grafana.foundation.cog.Builder<List<DayOfMonthRange>> daysOfMonth) {
        this.internal.daysOfMonth = daysOfMonth.build();
        return this;
    }
    
    public TimeIntervalBuilder months(com.grafana.foundation.cog.Builder<List<MonthRange>> months) {
        this.internal.months = months.build();
        return this;
    }
    
    public TimeIntervalBuilder years(com.grafana.foundation.cog.Builder<List<YearRange>> years) {
        this.internal.years = years.build();
        return this;
    }
    
    public TimeIntervalBuilder location(String location) {
        this.internal.location = location;
        return this;
    }
    public TimeInterval build() {
        return this.internal;
    }
}
