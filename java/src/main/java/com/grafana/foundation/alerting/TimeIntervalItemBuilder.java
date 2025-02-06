// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;

public class TimeIntervalItemBuilder implements com.grafana.foundation.cog.Builder<TimeIntervalItem> {
    protected final TimeIntervalItem internal;
    
    public TimeIntervalItemBuilder() {
        this.internal = new TimeIntervalItem();
    }
    public TimeIntervalItemBuilder daysOfMonth(List<String> daysOfMonth) {
        this.internal.daysOfMonth = daysOfMonth;
        return this;
    }
    
    public TimeIntervalItemBuilder location(String location) {
        this.internal.location = location;
        return this;
    }
    
    public TimeIntervalItemBuilder months(List<String> months) {
        this.internal.months = months;
        return this;
    }
    
    public TimeIntervalItemBuilder times(com.grafana.foundation.cog.Builder<List<TimeIntervalTimeRange>> times) {
        this.internal.times = times.build();
        return this;
    }
    
    public TimeIntervalItemBuilder weekdays(List<String> weekdays) {
        this.internal.weekdays = weekdays;
        return this;
    }
    
    public TimeIntervalItemBuilder years(List<String> years) {
        this.internal.years = years;
        return this;
    }
    public TimeIntervalItem build() {
        return this.internal;
    }
}
