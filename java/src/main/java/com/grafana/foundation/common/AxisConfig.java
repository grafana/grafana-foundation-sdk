// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// TODO docs
public class AxisConfig {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("axisPlacement")
    public AxisPlacement axisPlacement;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("axisColorMode")
    public AxisColorMode axisColorMode;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("axisLabel")
    public String axisLabel;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("axisWidth")
    public Double axisWidth;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("axisSoftMin")
    public Double axisSoftMin;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("axisSoftMax")
    public Double axisSoftMax;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("axisGridShow")
    public Boolean axisGridShow;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("scaleDistribution")
    public ScaleDistributionConfig scaleDistribution;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("axisCenteredZero")
    public Boolean axisCenteredZero;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AxisConfig> {
        private final AxisConfig internal;
        
        public Builder() {
            this.internal = new AxisConfig();
        }
    public Builder axisPlacement(AxisPlacement axisPlacement) {
    this.internal.axisPlacement = axisPlacement;
        return this;
    }
    
    public Builder axisColorMode(AxisColorMode axisColorMode) {
    this.internal.axisColorMode = axisColorMode;
        return this;
    }
    
    public Builder axisLabel(String axisLabel) {
    this.internal.axisLabel = axisLabel;
        return this;
    }
    
    public Builder axisWidth(Double axisWidth) {
    this.internal.axisWidth = axisWidth;
        return this;
    }
    
    public Builder axisSoftMin(Double axisSoftMin) {
    this.internal.axisSoftMin = axisSoftMin;
        return this;
    }
    
    public Builder axisSoftMax(Double axisSoftMax) {
    this.internal.axisSoftMax = axisSoftMax;
        return this;
    }
    
    public Builder axisGridShow(Boolean axisGridShow) {
    this.internal.axisGridShow = axisGridShow;
        return this;
    }
    
    public Builder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
    this.internal.scaleDistribution = scaleDistribution.build();
        return this;
    }
    
    public Builder axisCenteredZero(Boolean axisCenteredZero) {
    this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    public AxisConfig build() {
            return this.internal;
        }
    }
}
