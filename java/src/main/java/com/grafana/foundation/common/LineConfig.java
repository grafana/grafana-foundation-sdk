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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<LineConfig> {
        protected final LineConfig internal;
        
        public Builder() {
            this.internal = new LineConfig();
        }
    public Builder lineColor(String lineColor) {
    this.internal.lineColor = lineColor;
        return this;
    }
    
    public Builder lineWidth(Double lineWidth) {
    this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public Builder lineInterpolation(LineInterpolation lineInterpolation) {
    this.internal.lineInterpolation = lineInterpolation;
        return this;
    }
    
    public Builder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
    this.internal.lineStyle = lineStyle.build();
        return this;
    }
    
    public Builder spanNulls(BoolOrFloat64 spanNulls) {
    this.internal.spanNulls = spanNulls;
        return this;
    }
    public LineConfig build() {
            return this.internal;
        }
    }
}
