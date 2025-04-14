// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.VisibilityMode;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.ScaleDimensionConfig;
import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.LineStyle;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.TextDimensionConfig;

public class ScatterSeriesConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("x")
    public String x;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("y")
    public String y;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("show")
    public ScatterShow show;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointSize")
    public ScaleDimensionConfig pointSize;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointColor")
    public ColorDimensionConfig pointColor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lineColor")
    public ColorDimensionConfig lineColor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lineWidth")
    public Integer lineWidth;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lineStyle")
    public LineStyle lineStyle;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("label")
    public VisibilityMode label;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hideFrom")
    public HideSeriesConfig hideFrom;
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
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisCenteredZero")
    public Boolean axisCenteredZero;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("frame")
    public Double frame;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labelValue")
    public TextDimensionConfig labelValue;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisBorderShow")
    public Boolean axisBorderShow;
    public ScatterSeriesConfig() {
        this.show = ScatterShow.POINTS;
        this.label = VisibilityMode.AUTO;
    }
    public ScatterSeriesConfig(String x,String y,String name,ScatterShow show,ScaleDimensionConfig pointSize,ColorDimensionConfig pointColor,ColorDimensionConfig lineColor,Integer lineWidth,LineStyle lineStyle,VisibilityMode label,HideSeriesConfig hideFrom,AxisPlacement axisPlacement,AxisColorMode axisColorMode,String axisLabel,Double axisWidth,Double axisSoftMin,Double axisSoftMax,Boolean axisGridShow,ScaleDistributionConfig scaleDistribution,Boolean axisCenteredZero,Double frame,TextDimensionConfig labelValue,Boolean axisBorderShow) {
        this.x = x;
        this.y = y;
        this.name = name;
        this.show = show;
        this.pointSize = pointSize;
        this.pointColor = pointColor;
        this.lineColor = lineColor;
        this.lineWidth = lineWidth;
        this.lineStyle = lineStyle;
        this.label = label;
        this.hideFrom = hideFrom;
        this.axisPlacement = axisPlacement;
        this.axisColorMode = axisColorMode;
        this.axisLabel = axisLabel;
        this.axisWidth = axisWidth;
        this.axisSoftMin = axisSoftMin;
        this.axisSoftMax = axisSoftMax;
        this.axisGridShow = axisGridShow;
        this.scaleDistribution = scaleDistribution;
        this.axisCenteredZero = axisCenteredZero;
        this.frame = frame;
        this.labelValue = labelValue;
        this.axisBorderShow = axisBorderShow;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
