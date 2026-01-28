// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.logs;

import com.grafana.foundation.common.LogsSortOrder;
import com.grafana.foundation.common.LogsDedupStrategy;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder showLabels(Boolean showLabels) {
        this.internal.showLabels = showLabels;
        return this;
    }
    
    public OptionsBuilder showCommonLabels(Boolean showCommonLabels) {
        this.internal.showCommonLabels = showCommonLabels;
        return this;
    }
    
    public OptionsBuilder showTime(Boolean showTime) {
        this.internal.showTime = showTime;
        return this;
    }
    
    public OptionsBuilder wrapLogMessage(Boolean wrapLogMessage) {
        this.internal.wrapLogMessage = wrapLogMessage;
        return this;
    }
    
    public OptionsBuilder prettifyLogMessage(Boolean prettifyLogMessage) {
        this.internal.prettifyLogMessage = prettifyLogMessage;
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
    public Options build() {
        return this.internal;
    }
}
