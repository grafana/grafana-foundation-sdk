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
    public Options() {
        this.showLabels = false;
        this.showCommonLabels = false;
        this.showTime = false;
        this.wrapLogMessage = false;
        this.prettifyLogMessage = false;
        this.enableLogDetails = false;
        this.sortOrder = LogsSortOrder.DESCENDING;
        this.dedupStrategy = LogsDedupStrategy.NONE;
    }
    public Options(Boolean showLabels,Boolean showCommonLabels,Boolean showTime,Boolean wrapLogMessage,Boolean prettifyLogMessage,Boolean enableLogDetails,LogsSortOrder sortOrder,LogsDedupStrategy dedupStrategy) {
        this.showLabels = showLabels;
        this.showCommonLabels = showCommonLabels;
        this.showTime = showTime;
        this.wrapLogMessage = wrapLogMessage;
        this.prettifyLogMessage = prettifyLogMessage;
        this.enableLogDetails = enableLogDetails;
        this.sortOrder = sortOrder;
        this.dedupStrategy = dedupStrategy;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
