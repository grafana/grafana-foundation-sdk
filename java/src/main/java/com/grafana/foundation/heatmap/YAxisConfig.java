// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;

// Configuration options for the yAxis
public class YAxisConfig {
    // Sets the yAxis unit
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("unit")
    public String unit;
    // Reverses the yAxis
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("reverse")
    public Boolean reverse;
    // Controls the number of decimals for yAxis values
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("decimals")
    public Float decimals;
    // Sets the minimum value for the yAxis
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("min")
    public Float min;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisPlacement")
    public AxisPlacement axisPlacement;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisColorMode")
    public AxisColorMode axisColorMode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisLabel")
    public String axisLabel;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisWidth")
    public Double axisWidth;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisSoftMin")
    public Double axisSoftMin;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisSoftMax")
    public Double axisSoftMax;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisGridShow")
    public Boolean axisGridShow;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("scaleDistribution")
    public ScaleDistributionConfig scaleDistribution;
    // Sets the maximum value for the yAxis
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("max")
    public Float max;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisCenteredZero")
    public Boolean axisCenteredZero;
    public YAxisConfig() {
    }
    public YAxisConfig(String unit,Boolean reverse,Float decimals,Float min,AxisPlacement axisPlacement,AxisColorMode axisColorMode,String axisLabel,Double axisWidth,Double axisSoftMin,Double axisSoftMax,Boolean axisGridShow,ScaleDistributionConfig scaleDistribution,Float max,Boolean axisCenteredZero) {
        this.unit = unit;
        this.reverse = reverse;
        this.decimals = decimals;
        this.min = min;
        this.axisPlacement = axisPlacement;
        this.axisColorMode = axisColorMode;
        this.axisLabel = axisLabel;
        this.axisWidth = axisWidth;
        this.axisSoftMin = axisSoftMin;
        this.axisSoftMax = axisSoftMax;
        this.axisGridShow = axisGridShow;
        this.scaleDistribution = scaleDistribution;
        this.max = max;
        this.axisCenteredZero = axisCenteredZero;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
