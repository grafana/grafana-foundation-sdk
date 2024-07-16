// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Sparkline cell options
public class TableSparklineCellOptions { 
    @JsonProperty("type")
    public String type; 
    @JsonProperty("drawStyle")
    public GraphDrawStyle drawStyle; 
    @JsonProperty("gradientMode")
    public GraphGradientMode gradientMode; 
    @JsonProperty("thresholdsStyle")
    public GraphThresholdsStyleConfig thresholdsStyle; 
    @JsonProperty("lineColor")
    public String lineColor; 
    @JsonProperty("lineWidth")
    public Double lineWidth; 
    @JsonProperty("lineInterpolation")
    public LineInterpolation lineInterpolation; 
    @JsonProperty("lineStyle")
    public LineStyle lineStyle; 
    @JsonProperty("fillColor")
    public String fillColor; 
    @JsonProperty("fillOpacity")
    public Double fillOpacity; 
    @JsonProperty("showPoints")
    public VisibilityMode showPoints; 
    @JsonProperty("pointSize")
    public Double pointSize; 
    @JsonProperty("pointColor")
    public String pointColor; 
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
    @JsonProperty("barAlignment")
    public BarAlignment barAlignment; 
    @JsonProperty("barWidthFactor")
    public Double barWidthFactor; 
    @JsonProperty("stacking")
    public StackingConfig stacking; 
    @JsonProperty("hideFrom")
    public HideSeriesConfig hideFrom; 
    @JsonProperty("hideValue")
    public Boolean hideValue; 
    @JsonProperty("transform")
    public GraphTransform transform;
    // Indicate if null values should be treated as gaps or connected.
    // When the value is a number, it represents the maximum delta in the
    // X axis that should be considered connected.  For timeseries, this is milliseconds 
    @JsonProperty("spanNulls")
    public BoolOrFloat64 spanNulls; 
    @JsonProperty("fillBelowTo")
    public String fillBelowTo; 
    @JsonProperty("pointSymbol")
    public String pointSymbol; 
    @JsonProperty("axisBorderShow")
    public Boolean axisBorderShow; 
    @JsonProperty("barMaxWidth")
    public Double barMaxWidth;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TableSparklineCellOptions> {
        private final TableSparklineCellOptions internal;
        
        public Builder() {
            this.internal = new TableSparklineCellOptions();
    this.internal.type = "sparkline";
        }
    public Builder drawStyle(GraphDrawStyle drawStyle) {
    this.internal.drawStyle = drawStyle;
        return this;
    }
    
    public Builder gradientMode(GraphGradientMode gradientMode) {
    this.internal.gradientMode = gradientMode;
        return this;
    }
    
    public Builder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle) {
    this.internal.thresholdsStyle = thresholdsStyle.build();
        return this;
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
    
    public Builder fillColor(String fillColor) {
    this.internal.fillColor = fillColor;
        return this;
    }
    
    public Builder fillOpacity(Double fillOpacity) {
    this.internal.fillOpacity = fillOpacity;
        return this;
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
    
    public Builder barAlignment(BarAlignment barAlignment) {
    this.internal.barAlignment = barAlignment;
        return this;
    }
    
    public Builder barWidthFactor(Double barWidthFactor) {
    this.internal.barWidthFactor = barWidthFactor;
        return this;
    }
    
    public Builder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking) {
    this.internal.stacking = stacking.build();
        return this;
    }
    
    public Builder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    this.internal.hideFrom = hideFrom.build();
        return this;
    }
    
    public Builder hideValue(Boolean hideValue) {
    this.internal.hideValue = hideValue;
        return this;
    }
    
    public Builder transform(GraphTransform transform) {
    this.internal.transform = transform;
        return this;
    }
    
    public Builder spanNulls(BoolOrFloat64 spanNulls) {
    this.internal.spanNulls = spanNulls;
        return this;
    }
    
    public Builder fillBelowTo(String fillBelowTo) {
    this.internal.fillBelowTo = fillBelowTo;
        return this;
    }
    
    public Builder pointSymbol(String pointSymbol) {
    this.internal.pointSymbol = pointSymbol;
        return this;
    }
    
    public Builder axisBorderShow(Boolean axisBorderShow) {
    this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
    
    public Builder barMaxWidth(Double barMaxWidth) {
    this.internal.barMaxWidth = barMaxWidth;
        return this;
    }
    public TableSparklineCellOptions build() {
            return this.internal;
        }
    }
}
