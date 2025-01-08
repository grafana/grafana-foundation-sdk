// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class LineConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lineColor")
    public String lineColor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lineWidth")
    public Double lineWidth;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lineInterpolation")
    public LineInterpolation lineInterpolation;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lineStyle")
    public LineStyle lineStyle;
    // Indicate if null values should be treated as gaps or connected.
    // When the value is a number, it represents the maximum delta in the
    // X axis that should be considered connected.  For timeseries, this is milliseconds
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("spanNulls")
    public BoolOrFloat64 spanNulls;
    public LineConfig() {
    }
    
    public LineConfig(String lineColor,Double lineWidth,LineInterpolation lineInterpolation,LineStyle lineStyle,BoolOrFloat64 spanNulls) {
        this.lineColor = lineColor;
        this.lineWidth = lineWidth;
        this.lineInterpolation = lineInterpolation;
        this.lineStyle = lineStyle;
        this.spanNulls = spanNulls;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
