// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.logs;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.LogsSortOrder;
import com.grafana.foundation.common.LogsDedupStrategy;
import java.util.List;

public class Options {
    @JsonProperty("showLabels")
    public Boolean showLabels;
    @JsonProperty("showCommonLabels")
    public Boolean showCommonLabels;
    @JsonProperty("showTime")
    public Boolean showTime;
    @JsonProperty("showLogContextToggle")
    public Boolean showLogContextToggle;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showControls")
    public Boolean showControls;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("controlsStorageKey")
    public String controlsStorageKey;
    @JsonProperty("wrapLogMessage")
    public Boolean wrapLogMessage;
    @JsonProperty("prettifyLogMessage")
    public Boolean prettifyLogMessage;
    @JsonProperty("enableLogDetails")
    public Boolean enableLogDetails;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("sortOrder")
    public LogsSortOrder sortOrder;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("dedupStrategy")
    public LogsDedupStrategy dedupStrategy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("enableInfiniteScrolling")
    public Boolean enableInfiniteScrolling;
    // TODO: figure out how to define callbacks
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onClickFilterLabel")
    public Object onClickFilterLabel;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onClickFilterOutLabel")
    public Object onClickFilterOutLabel;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("isFilterLabelActive")
    public Object isFilterLabelActive;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onClickFilterString")
    public Object onClickFilterString;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onClickFilterOutString")
    public Object onClickFilterOutString;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onClickShowField")
    public Object onClickShowField;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onClickHideField")
    public Object onClickHideField;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onLogOptionsChange")
    public Object onLogOptionsChange;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("logRowMenuIconsBefore")
    public Object logRowMenuIconsBefore;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("logRowMenuIconsAfter")
    public Object logRowMenuIconsAfter;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onNewLogsReceived")
    public Object onNewLogsReceived;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("displayedFields")
    public List<String> displayedFields;
    public Options() {
    }
    public Options(Boolean showLabels,Boolean showCommonLabels,Boolean showTime,Boolean showLogContextToggle,Boolean showControls,String controlsStorageKey,Boolean wrapLogMessage,Boolean prettifyLogMessage,Boolean enableLogDetails,LogsSortOrder sortOrder,LogsDedupStrategy dedupStrategy,Boolean enableInfiniteScrolling,Object onClickFilterLabel,Object onClickFilterOutLabel,Object isFilterLabelActive,Object onClickFilterString,Object onClickFilterOutString,Object onClickShowField,Object onClickHideField,Object onLogOptionsChange,Object logRowMenuIconsBefore,Object logRowMenuIconsAfter,Object onNewLogsReceived,List<String> displayedFields) {
        this.showLabels = showLabels;
        this.showCommonLabels = showCommonLabels;
        this.showTime = showTime;
        this.showLogContextToggle = showLogContextToggle;
        this.showControls = showControls;
        this.controlsStorageKey = controlsStorageKey;
        this.wrapLogMessage = wrapLogMessage;
        this.prettifyLogMessage = prettifyLogMessage;
        this.enableLogDetails = enableLogDetails;
        this.sortOrder = sortOrder;
        this.dedupStrategy = dedupStrategy;
        this.enableInfiniteScrolling = enableInfiniteScrolling;
        this.onClickFilterLabel = onClickFilterLabel;
        this.onClickFilterOutLabel = onClickFilterOutLabel;
        this.isFilterLabelActive = isFilterLabelActive;
        this.onClickFilterString = onClickFilterString;
        this.onClickFilterOutString = onClickFilterOutString;
        this.onClickShowField = onClickShowField;
        this.onClickHideField = onClickHideField;
        this.onLogOptionsChange = onLogOptionsChange;
        this.logRowMenuIconsBefore = logRowMenuIconsBefore;
        this.logRowMenuIconsAfter = logRowMenuIconsAfter;
        this.onNewLogsReceived = onNewLogsReceived;
        this.displayedFields = displayedFields;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
