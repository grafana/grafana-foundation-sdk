// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.grafana.foundation.common.ScaleDimensionConfig;
import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.LineStyle;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.TextDimensionConfig;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ScatterSeriesConfig { 
    @JsonProperty("x")
    public String x; 
    @JsonProperty("y")
    public String y; 
    @JsonProperty("show")
    public ScatterShow show; 
    @JsonProperty("pointSize")
    public ScaleDimensionConfig pointSize; 
    @JsonProperty("pointColor")
    public ColorDimensionConfig pointColor; 
    @JsonProperty("lineColor")
    public ColorDimensionConfig lineColor; 
    @JsonProperty("lineWidth")
    public Integer lineWidth; 
    @JsonProperty("lineStyle")
    public LineStyle lineStyle; 
    @JsonProperty("label")
    public VisibilityMode label; 
    @JsonProperty("hideFrom")
    public HideSeriesConfig hideFrom; 
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
    @JsonProperty("name")
    public String name; 
    @JsonProperty("labelValue")
    public TextDimensionConfig labelValue; 
    @JsonProperty("axisCenteredZero")
    public Boolean axisCenteredZero;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
