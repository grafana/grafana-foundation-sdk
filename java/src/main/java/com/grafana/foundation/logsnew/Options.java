// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.logsnew;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.LogsSortOrder;
import com.grafana.foundation.common.LogsDedupStrategy;

public class Options {
    @JsonProperty("showControls")
    public Boolean showControls;
    @JsonProperty("showTime")
    public Boolean showTime;
    @JsonProperty("wrapLogMessage")
    public Boolean wrapLogMessage;
    @JsonProperty("enableLogDetails")
    public Boolean enableLogDetails;
    @JsonProperty("syntaxHighlighting")
    public Boolean syntaxHighlighting;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("sortOrder")
    public LogsSortOrder sortOrder;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("dedupStrategy")
    public LogsDedupStrategy dedupStrategy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("grammar")
    public Object grammar;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("enableInfiniteScrolling")
    public Boolean enableInfiniteScrolling;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onLogOptionsChange")
    public Object onLogOptionsChange;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onNewLogsReceived")
    public Object onNewLogsReceived;
    public Options() {
    }
    public Options(Boolean showControls,Boolean showTime,Boolean wrapLogMessage,Boolean enableLogDetails,Boolean syntaxHighlighting,LogsSortOrder sortOrder,LogsDedupStrategy dedupStrategy,Object grammar,Boolean enableInfiniteScrolling,Object onLogOptionsChange,Object onNewLogsReceived) {
        this.showControls = showControls;
        this.showTime = showTime;
        this.wrapLogMessage = wrapLogMessage;
        this.enableLogDetails = enableLogDetails;
        this.syntaxHighlighting = syntaxHighlighting;
        this.sortOrder = sortOrder;
        this.dedupStrategy = dedupStrategy;
        this.grammar = grammar;
        this.enableInfiniteScrolling = enableInfiniteScrolling;
        this.onLogOptionsChange = onLogOptionsChange;
        this.onNewLogsReceived = onNewLogsReceived;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
