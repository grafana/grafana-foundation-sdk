// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class PointsConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showPoints")
    public VisibilityMode showPoints;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointSize")
    public Double pointSize;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointColor")
    public String pointColor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointSymbol")
    public String pointSymbol;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<PointsConfig> {
        private final PointsConfig internal;
        
        public Builder() {
            this.internal = new PointsConfig();
        }
    public Builder showPoints(VisibilityMode showPoints) {
    this.internal.showPoints = showPoints;
        return this;
    }
    
    public Builder pointSize(Double pointSize) {
    this.internal.pointSize = pointSize;
        return this;
    }
    
    public Builder pointColor(String pointColor) {
    this.internal.pointColor = pointColor;
        return this;
    }
    
    public Builder pointSymbol(String pointSymbol) {
    this.internal.pointSymbol = pointSymbol;
        return this;
    }
    public PointsConfig build() {
            return this.internal;
        }
    }
}
