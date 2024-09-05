// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;

// Configuration options for the yAxis
public class YAxisConfig {
    // Sets the yAxis unit
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("unit")
    public String unit;
    // Reverses the yAxis
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("reverse")
    public Boolean reverse;
    // Controls the number of decimals for yAxis values
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("decimals")
    public Float decimals;
    // Sets the minimum value for the yAxis
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("min")
    public Float min;
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
    // Sets the maximum value for the yAxis
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("max")
    public Float max;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("axisCenteredZero")
    public Boolean axisCenteredZero;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
