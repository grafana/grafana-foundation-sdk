// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class GraphFieldConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("drawStyle")
    public GraphDrawStyle drawStyle;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("gradientMode")
    public GraphGradientMode gradientMode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("thresholdsStyle")
    public GraphThresholdsStyleConfig thresholdsStyle;
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
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillColor")
    public String fillColor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillOpacity")
    public Double fillOpacity;
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
    @JsonProperty("barAlignment")
    public BarAlignment barAlignment;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("barWidthFactor")
    public Double barWidthFactor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("stacking")
    public StackingConfig stacking;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hideFrom")
    public HideSeriesConfig hideFrom;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("transform")
    public GraphTransform transform;
    // Indicate if null values should be treated as gaps or connected.
    // When the value is a number, it represents the maximum delta in the
    // X axis that should be considered connected.  For timeseries, this is milliseconds
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("spanNulls")
    public BoolOrFloat64 spanNulls;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillBelowTo")
    public String fillBelowTo;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointSymbol")
    public String pointSymbol;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisCenteredZero")
    public Boolean axisCenteredZero;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("barMaxWidth")
    public Double barMaxWidth;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("insertNulls")
    public BoolOrUint32 insertNulls;
    public GraphFieldConfig() {
    }
    public GraphFieldConfig(GraphDrawStyle drawStyle,GraphGradientMode gradientMode,GraphThresholdsStyleConfig thresholdsStyle,String lineColor,Double lineWidth,LineInterpolation lineInterpolation,LineStyle lineStyle,String fillColor,Double fillOpacity,VisibilityMode showPoints,Double pointSize,String pointColor,AxisPlacement axisPlacement,AxisColorMode axisColorMode,String axisLabel,Double axisWidth,Double axisSoftMin,Double axisSoftMax,Boolean axisGridShow,ScaleDistributionConfig scaleDistribution,BarAlignment barAlignment,Double barWidthFactor,StackingConfig stacking,HideSeriesConfig hideFrom,GraphTransform transform,BoolOrFloat64 spanNulls,String fillBelowTo,String pointSymbol,Boolean axisCenteredZero,Double barMaxWidth,BoolOrUint32 insertNulls) {
        this.drawStyle = drawStyle;
        this.gradientMode = gradientMode;
        this.thresholdsStyle = thresholdsStyle;
        this.lineColor = lineColor;
        this.lineWidth = lineWidth;
        this.lineInterpolation = lineInterpolation;
        this.lineStyle = lineStyle;
        this.fillColor = fillColor;
        this.fillOpacity = fillOpacity;
        this.showPoints = showPoints;
        this.pointSize = pointSize;
        this.pointColor = pointColor;
        this.axisPlacement = axisPlacement;
        this.axisColorMode = axisColorMode;
        this.axisLabel = axisLabel;
        this.axisWidth = axisWidth;
        this.axisSoftMin = axisSoftMin;
        this.axisSoftMax = axisSoftMax;
        this.axisGridShow = axisGridShow;
        this.scaleDistribution = scaleDistribution;
        this.barAlignment = barAlignment;
        this.barWidthFactor = barWidthFactor;
        this.stacking = stacking;
        this.hideFrom = hideFrom;
        this.transform = transform;
        this.spanNulls = spanNulls;
        this.fillBelowTo = fillBelowTo;
        this.pointSymbol = pointSymbol;
        this.axisCenteredZero = axisCenteredZero;
        this.barMaxWidth = barMaxWidth;
        this.insertNulls = insertNulls;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
