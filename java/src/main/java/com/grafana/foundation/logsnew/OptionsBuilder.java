// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.logsnew;

import com.grafana.foundation.common.LogsSortOrder;
import com.grafana.foundation.common.LogsDedupStrategy;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder showTime(Boolean showTime) {
        this.internal.showTime = showTime;
        return this;
    }
    
    public OptionsBuilder wrapLogMessage(Boolean wrapLogMessage) {
        this.internal.wrapLogMessage = wrapLogMessage;
        return this;
    }
    
    public OptionsBuilder enableLogDetails(Boolean enableLogDetails) {
        this.internal.enableLogDetails = enableLogDetails;
        return this;
    }
    
    public OptionsBuilder sortOrder(LogsSortOrder sortOrder) {
        this.internal.sortOrder = sortOrder;
        return this;
    }
    
    public OptionsBuilder dedupStrategy(LogsDedupStrategy dedupStrategy) {
        this.internal.dedupStrategy = dedupStrategy;
        return this;
    }
    
    public OptionsBuilder enableInfiniteScrolling(Boolean enableInfiniteScrolling) {
        this.internal.enableInfiniteScrolling = enableInfiniteScrolling;
        return this;
    }
    
    public OptionsBuilder onNewLogsReceived(Object onNewLogsReceived) {
        this.internal.onNewLogsReceived = onNewLogsReceived;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
