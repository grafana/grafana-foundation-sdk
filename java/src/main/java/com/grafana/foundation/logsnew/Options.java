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
    @JsonProperty("showTime")
    public Boolean showTime;
    @JsonProperty("wrapLogMessage")
    public Boolean wrapLogMessage;
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
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onNewLogsReceived")
    public Object onNewLogsReceived;
    public Options() {
        this.showTime = false;
        this.wrapLogMessage = false;
        this.enableLogDetails = false;
        this.sortOrder = LogsSortOrder.DESCENDING;
        this.dedupStrategy = LogsDedupStrategy.NONE;
    }
    public Options(Boolean showTime,Boolean wrapLogMessage,Boolean enableLogDetails,LogsSortOrder sortOrder,LogsDedupStrategy dedupStrategy,Boolean enableInfiniteScrolling,Object onNewLogsReceived) {
        this.showTime = showTime;
        this.wrapLogMessage = wrapLogMessage;
        this.enableLogDetails = enableLogDetails;
        this.sortOrder = sortOrder;
        this.dedupStrategy = dedupStrategy;
        this.enableInfiniteScrolling = enableInfiniteScrolling;
        this.onNewLogsReceived = onNewLogsReceived;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
