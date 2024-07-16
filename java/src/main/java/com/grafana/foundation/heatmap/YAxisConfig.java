// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Configuration options for the yAxis
public class YAxisConfig {
    // Sets the yAxis unit 
    @JsonProperty("unit")
    public String unit;
    // Reverses the yAxis 
    @JsonProperty("reverse")
    public Boolean reverse;
    // Controls the number of decimals for yAxis values 
    @JsonProperty("decimals")
    public Float decimals;
    // Sets the minimum value for the yAxis 
    @JsonProperty("min")
    public Float min; 
    @JsonProperty("axisPlacement")
    public AxisPlacement axisPlacement; 
    @JsonProperty("axisColorMode")
    public AxisColorMode axisColorMode; 
    @JsonProperty("axisLabel")
    public String axisLabel; 
    @JsonProperty("axisWidth")
    public Double axisWidth; 
    @JsonProperty("axisSoftMin")
    public Double axisSoftMin; 
    @JsonProperty("axisSoftMax")
    public Double axisSoftMax; 
    @JsonProperty("axisGridShow")
    public Boolean axisGridShow; 
    @JsonProperty("scaleDistribution")
    public ScaleDistributionConfig scaleDistribution; 
    @JsonProperty("axisCenteredZero")
    public Boolean axisCenteredZero;
    // Sets the maximum value for the yAxis 
    @JsonProperty("max")
    public Float max; 
    @JsonProperty("axisBorderShow")
    public Boolean axisBorderShow;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
