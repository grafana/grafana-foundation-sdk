// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// TODO docs
public class PointsConfig { 
    @JsonProperty("showPoints")
    public VisibilityMode showPoints; 
    @JsonProperty("pointSize")
    public Double pointSize; 
    @JsonProperty("pointColor")
    public String pointColor; 
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
