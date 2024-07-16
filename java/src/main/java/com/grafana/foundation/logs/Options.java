// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.logs;

import com.grafana.foundation.common.LogsSortOrder;
import com.grafana.foundation.common.LogsDedupStrategy;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

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
    @JsonProperty("sortOrder")
    public LogsSortOrder sortOrder; 
    @JsonProperty("dedupStrategy")
    public LogsDedupStrategy dedupStrategy;
    // TODO: figure out how to define callbacks 
    @JsonProperty("onClickFilterLabel")
    public Object onClickFilterLabel; 
    @JsonProperty("onClickFilterOutLabel")
    public Object onClickFilterOutLabel; 
    @JsonProperty("isFilterLabelActive")
    public Object isFilterLabelActive; 
    @JsonProperty("onClickFilterString")
    public Object onClickFilterString; 
    @JsonProperty("onClickFilterOutString")
    public Object onClickFilterOutString;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
