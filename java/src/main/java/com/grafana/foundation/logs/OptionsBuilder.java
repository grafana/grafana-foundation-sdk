// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.logs;

import com.grafana.foundation.common.LogsSortOrder;
import com.grafana.foundation.common.LogsDedupStrategy;
import java.util.List;

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
    
    public OptionsBuilder showLogContextToggle(Boolean showLogContextToggle) {
        this.internal.showLogContextToggle = showLogContextToggle;
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
    
    public OptionsBuilder enableInfiniteScrolling(Boolean enableInfiniteScrolling) {
        this.internal.enableInfiniteScrolling = enableInfiniteScrolling;
        return this;
    }
    
    public OptionsBuilder onClickFilterLabel(Object onClickFilterLabel) {
        this.internal.onClickFilterLabel = onClickFilterLabel;
        return this;
    }
    
    public OptionsBuilder onClickFilterOutLabel(Object onClickFilterOutLabel) {
        this.internal.onClickFilterOutLabel = onClickFilterOutLabel;
        return this;
    }
    
    public OptionsBuilder isFilterLabelActive(Object isFilterLabelActive) {
        this.internal.isFilterLabelActive = isFilterLabelActive;
        return this;
    }
    
    public OptionsBuilder onClickFilterString(Object onClickFilterString) {
        this.internal.onClickFilterString = onClickFilterString;
        return this;
    }
    
    public OptionsBuilder onClickFilterOutString(Object onClickFilterOutString) {
        this.internal.onClickFilterOutString = onClickFilterOutString;
        return this;
    }
    
    public OptionsBuilder onClickShowField(Object onClickShowField) {
        this.internal.onClickShowField = onClickShowField;
        return this;
    }
    
    public OptionsBuilder onClickHideField(Object onClickHideField) {
        this.internal.onClickHideField = onClickHideField;
        return this;
    }
    
    public OptionsBuilder logRowMenuIconsBefore(Object logRowMenuIconsBefore) {
        this.internal.logRowMenuIconsBefore = logRowMenuIconsBefore;
        return this;
    }
    
    public OptionsBuilder logRowMenuIconsAfter(Object logRowMenuIconsAfter) {
        this.internal.logRowMenuIconsAfter = logRowMenuIconsAfter;
        return this;
    }
    
    public OptionsBuilder onNewLogsReceived(Object onNewLogsReceived) {
        this.internal.onNewLogsReceived = onNewLogsReceived;
        return this;
    }
    
    public OptionsBuilder displayedFields(List<String> displayedFields) {
        this.internal.displayedFields = displayedFields;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
