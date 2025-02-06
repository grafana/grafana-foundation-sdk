// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;

public class TimeIntervalBuilder implements com.grafana.foundation.cog.Builder<TimeInterval> {
    protected final TimeInterval internal;
    
    public TimeIntervalBuilder() {
        this.internal = new TimeInterval();
    }
    public TimeIntervalBuilder daysOfMonth(List<String> daysOfMonth) {
        this.internal.daysOfMonth = daysOfMonth;
        return this;
    }
    
    public TimeIntervalBuilder location(String location) {
        this.internal.location = location;
        return this;
    }
    
    public TimeIntervalBuilder months(List<String> months) {
        this.internal.months = months;
        return this;
    }
    
    public TimeIntervalBuilder times(com.grafana.foundation.cog.Builder<List<TimeRange>> times) {
        this.internal.times = times.build();
        return this;
    }
    
    public TimeIntervalBuilder weekdays(List<String> weekdays) {
        this.internal.weekdays = weekdays;
        return this;
    }
    
    public TimeIntervalBuilder years(List<String> years) {
        this.internal.years = years;
        return this;
    }
    public TimeInterval build() {
        return this.internal;
    }
}
