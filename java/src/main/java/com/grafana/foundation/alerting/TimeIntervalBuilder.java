// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;
import java.util.LinkedList;

public class TimeIntervalBuilder implements com.grafana.foundation.cog.Builder<TimeInterval> {
    protected final TimeInterval internal;
    
    public TimeIntervalBuilder() {
        this.internal = new TimeInterval();
    }
    public TimeIntervalBuilder times(List<com.grafana.foundation.cog.Builder<TimeRange>> times) {
        List<TimeRange> timesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TimeRange> r1 : times) {
                TimeRange timesDepth1 = r1.build();
                timesResources.add(timesDepth1); 
        }
        this.internal.times = timesResources;
        return this;
    }
    
    public TimeIntervalBuilder weekdays(List<com.grafana.foundation.cog.Builder<WeekdayRange>> weekdays) {
        List<WeekdayRange> weekdaysResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<WeekdayRange> r1 : weekdays) {
                WeekdayRange weekdaysDepth1 = r1.build();
                weekdaysResources.add(weekdaysDepth1); 
        }
        this.internal.weekdays = weekdaysResources;
        return this;
    }
    
    public TimeIntervalBuilder daysOfMonth(List<com.grafana.foundation.cog.Builder<DayOfMonthRange>> daysOfMonth) {
        List<DayOfMonthRange> daysOfMonthResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<DayOfMonthRange> r1 : daysOfMonth) {
                DayOfMonthRange daysOfMonthDepth1 = r1.build();
                daysOfMonthResources.add(daysOfMonthDepth1); 
        }
        this.internal.daysOfMonth = daysOfMonthResources;
        return this;
    }
    
    public TimeIntervalBuilder months(List<com.grafana.foundation.cog.Builder<MonthRange>> months) {
        List<MonthRange> monthsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<MonthRange> r1 : months) {
                MonthRange monthsDepth1 = r1.build();
                monthsResources.add(monthsDepth1); 
        }
        this.internal.months = monthsResources;
        return this;
    }
    
    public TimeIntervalBuilder years(List<com.grafana.foundation.cog.Builder<YearRange>> years) {
        List<YearRange> yearsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<YearRange> r1 : years) {
                YearRange yearsDepth1 = r1.build();
                yearsResources.add(yearsDepth1); 
        }
        this.internal.years = yearsResources;
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
