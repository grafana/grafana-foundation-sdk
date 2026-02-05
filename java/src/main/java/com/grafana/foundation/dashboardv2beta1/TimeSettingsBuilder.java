// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class TimeSettingsBuilder implements com.grafana.foundation.cog.Builder<TimeSettingsSpec> {
    protected final TimeSettingsSpec internal;
    
    public TimeSettingsBuilder() {
        this.internal = new TimeSettingsSpec();
    }
    public TimeSettingsBuilder timezone(String timezone) {
        this.internal.timezone = timezone;
        return this;
    }
    
    public TimeSettingsBuilder from(String from) {
        this.internal.from = from;
        return this;
    }
    
    public TimeSettingsBuilder to(String to) {
        this.internal.to = to;
        return this;
    }
    
    public TimeSettingsBuilder autoRefresh(String autoRefresh) {
        this.internal.autoRefresh = autoRefresh;
        return this;
    }
    
    public TimeSettingsBuilder autoRefreshIntervals(List<String> autoRefreshIntervals) {
        this.internal.autoRefreshIntervals = autoRefreshIntervals;
        return this;
    }
    
    public TimeSettingsBuilder quickRanges(List<com.grafana.foundation.cog.Builder<TimeRangeOption>> quickRanges) {
        List<TimeRangeOption> quickRangesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TimeRangeOption> r1 : quickRanges) {
                TimeRangeOption quickRangesDepth1 = r1.build();
                quickRangesResources.add(quickRangesDepth1); 
        }
        this.internal.quickRanges = quickRangesResources;
        return this;
    }
    
    public TimeSettingsBuilder hideTimepicker(Boolean hideTimepicker) {
        this.internal.hideTimepicker = hideTimepicker;
        return this;
    }
    
    public TimeSettingsBuilder weekStart(TimeSettingsSpecWeekStart weekStart) {
        this.internal.weekStart = weekStart;
        return this;
    }
    
    public TimeSettingsBuilder fiscalYearStartMonth(Long fiscalYearStartMonth) {
        this.internal.fiscalYearStartMonth = fiscalYearStartMonth;
        return this;
    }
    
    public TimeSettingsBuilder nowDelay(String nowDelay) {
        this.internal.nowDelay = nowDelay;
        return this;
    }
    public TimeSettingsSpec build() {
        return this.internal;
    }
}
