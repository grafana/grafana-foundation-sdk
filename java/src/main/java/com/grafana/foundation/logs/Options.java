// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.logs;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.LogsSortOrder;
import com.grafana.foundation.common.LogsDedupStrategy;

public class Options {
    @JsonProperty("showLabels")
    public Boolean showLabels;
    @JsonProperty("showCommonLabels")
    public Boolean showCommonLabels;
    @JsonProperty("showTime")
    public Boolean showTime;
    @JsonProperty("showLogContextToggle")
    public Boolean showLogContextToggle;
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
    // Display controls to jump to the last or first log line, and filters by log level.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showControls")
    public Boolean showControls;
    // Show a component to manage the displayed fields from the logs.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showFieldSelector")
    public Boolean showFieldSelector;
    // Use a predefined coloring scheme to highlight relevant parts of the log lines.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("syntaxHighlighting")
    public Boolean syntaxHighlighting;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fontSize")
    public OptionsFontSize fontSize;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("detailsMode")
    public OptionsDetailsMode detailsMode;
    public Options() {
        this.showLabels = false;
        this.showCommonLabels = false;
        this.showTime = false;
        this.showLogContextToggle = false;
        this.wrapLogMessage = false;
        this.prettifyLogMessage = false;
        this.enableLogDetails = false;
        this.sortOrder = LogsSortOrder.DESCENDING;
        this.dedupStrategy = LogsDedupStrategy.NONE;
    }
    public Options(Boolean showLabels,Boolean showCommonLabels,Boolean showTime,Boolean showLogContextToggle,Boolean wrapLogMessage,Boolean prettifyLogMessage,Boolean enableLogDetails,LogsSortOrder sortOrder,LogsDedupStrategy dedupStrategy,Boolean enableInfiniteScrolling,Boolean showControls,Boolean showFieldSelector,Boolean syntaxHighlighting,OptionsFontSize fontSize,OptionsDetailsMode detailsMode) {
        this.showLabels = showLabels;
        this.showCommonLabels = showCommonLabels;
        this.showTime = showTime;
        this.showLogContextToggle = showLogContextToggle;
        this.wrapLogMessage = wrapLogMessage;
        this.prettifyLogMessage = prettifyLogMessage;
        this.enableLogDetails = enableLogDetails;
        this.sortOrder = sortOrder;
        this.dedupStrategy = dedupStrategy;
        this.enableInfiniteScrolling = enableInfiniteScrolling;
        this.showControls = showControls;
        this.showFieldSelector = showFieldSelector;
        this.syntaxHighlighting = syntaxHighlighting;
        this.fontSize = fontSize;
        this.detailsMode = detailsMode;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
