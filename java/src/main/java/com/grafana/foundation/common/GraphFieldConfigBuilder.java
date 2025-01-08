// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class GraphFieldConfigBuilder implements com.grafana.foundation.cog.Builder<GraphFieldConfig> {
    protected final GraphFieldConfig internal;
    
    public GraphFieldConfigBuilder() {
        this.internal = new GraphFieldConfig();
    }
    public GraphFieldConfigBuilder drawStyle(GraphDrawStyle drawStyle) {
    this.internal.drawStyle = drawStyle;
        return this;
    }
    
    public GraphFieldConfigBuilder gradientMode(GraphGradientMode gradientMode) {
    this.internal.gradientMode = gradientMode;
        return this;
    }
    
    public GraphFieldConfigBuilder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle) {
    this.internal.thresholdsStyle = thresholdsStyle.build();
        return this;
    }
    
    public GraphFieldConfigBuilder transform(GraphTransform transform) {
    this.internal.transform = transform;
        return this;
    }
    
    public GraphFieldConfigBuilder lineColor(String lineColor) {
    this.internal.lineColor = lineColor;
        return this;
    }
    
    public GraphFieldConfigBuilder lineWidth(Double lineWidth) {
    this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public GraphFieldConfigBuilder lineInterpolation(LineInterpolation lineInterpolation) {
    this.internal.lineInterpolation = lineInterpolation;
        return this;
    }
    
    public GraphFieldConfigBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
    this.internal.lineStyle = lineStyle.build();
        return this;
    }
    
    public GraphFieldConfigBuilder fillColor(String fillColor) {
    this.internal.fillColor = fillColor;
        return this;
    }
    
    public GraphFieldConfigBuilder fillOpacity(Double fillOpacity) {
    this.internal.fillOpacity = fillOpacity;
        return this;
    }
    
    public GraphFieldConfigBuilder showPoints(VisibilityMode showPoints) {
    this.internal.showPoints = showPoints;
        return this;
    }
    
    public GraphFieldConfigBuilder pointSize(Double pointSize) {
    this.internal.pointSize = pointSize;
        return this;
    }
    
    public GraphFieldConfigBuilder pointColor(String pointColor) {
    this.internal.pointColor = pointColor;
        return this;
    }
    
    public GraphFieldConfigBuilder axisPlacement(AxisPlacement axisPlacement) {
    this.internal.axisPlacement = axisPlacement;
        return this;
    }
    
    public GraphFieldConfigBuilder axisColorMode(AxisColorMode axisColorMode) {
    this.internal.axisColorMode = axisColorMode;
        return this;
    }
    
    public GraphFieldConfigBuilder axisLabel(String axisLabel) {
    this.internal.axisLabel = axisLabel;
        return this;
    }
    
    public GraphFieldConfigBuilder axisWidth(Double axisWidth) {
    this.internal.axisWidth = axisWidth;
        return this;
    }
    
    public GraphFieldConfigBuilder axisSoftMin(Double axisSoftMin) {
    this.internal.axisSoftMin = axisSoftMin;
        return this;
    }
    
    public GraphFieldConfigBuilder axisSoftMax(Double axisSoftMax) {
    this.internal.axisSoftMax = axisSoftMax;
        return this;
    }
    
    public GraphFieldConfigBuilder axisGridShow(Boolean axisGridShow) {
    this.internal.axisGridShow = axisGridShow;
        return this;
    }
    
    public GraphFieldConfigBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
    this.internal.scaleDistribution = scaleDistribution.build();
        return this;
    }
    
    public GraphFieldConfigBuilder axisCenteredZero(Boolean axisCenteredZero) {
    this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    
    public GraphFieldConfigBuilder barAlignment(BarAlignment barAlignment) {
    this.internal.barAlignment = barAlignment;
        return this;
    }
    
    public GraphFieldConfigBuilder barWidthFactor(Double barWidthFactor) {
    this.internal.barWidthFactor = barWidthFactor;
        return this;
    }
    
    public GraphFieldConfigBuilder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking) {
    this.internal.stacking = stacking.build();
        return this;
    }
    
    public GraphFieldConfigBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    this.internal.hideFrom = hideFrom.build();
        return this;
    }
    
    public GraphFieldConfigBuilder insertNulls(BoolOrFloat64 insertNulls) {
    this.internal.insertNulls = insertNulls;
        return this;
    }
    
    public GraphFieldConfigBuilder spanNulls(BoolOrFloat64 spanNulls) {
    this.internal.spanNulls = spanNulls;
        return this;
    }
    
    public GraphFieldConfigBuilder fillBelowTo(String fillBelowTo) {
    this.internal.fillBelowTo = fillBelowTo;
        return this;
    }
    
    public GraphFieldConfigBuilder pointSymbol(String pointSymbol) {
    this.internal.pointSymbol = pointSymbol;
        return this;
    }
    
    public GraphFieldConfigBuilder axisBorderShow(Boolean axisBorderShow) {
    this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
    
    public GraphFieldConfigBuilder barMaxWidth(Double barMaxWidth) {
    this.internal.barMaxWidth = barMaxWidth;
        return this;
    }
    public GraphFieldConfig build() {
        return this.internal;
    }
}
