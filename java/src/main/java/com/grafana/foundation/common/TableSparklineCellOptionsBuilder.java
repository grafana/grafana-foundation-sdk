// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class TableSparklineCellOptionsBuilder implements com.grafana.foundation.cog.Builder<TableSparklineCellOptions> {
    protected final TableSparklineCellOptions internal;
    
    public TableSparklineCellOptionsBuilder() {
        this.internal = new TableSparklineCellOptions();
    this.internal.type = "sparkline";
    }
    public TableSparklineCellOptionsBuilder drawStyle(GraphDrawStyle drawStyle) {
    this.internal.drawStyle = drawStyle;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder gradientMode(GraphGradientMode gradientMode) {
    this.internal.gradientMode = gradientMode;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle) {
    this.internal.thresholdsStyle = thresholdsStyle.build();
        return this;
    }
    
    public TableSparklineCellOptionsBuilder lineColor(String lineColor) {
    this.internal.lineColor = lineColor;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder lineWidth(Double lineWidth) {
    this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder lineInterpolation(LineInterpolation lineInterpolation) {
    this.internal.lineInterpolation = lineInterpolation;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
    this.internal.lineStyle = lineStyle.build();
        return this;
    }
    
    public TableSparklineCellOptionsBuilder fillColor(String fillColor) {
    this.internal.fillColor = fillColor;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder fillOpacity(Double fillOpacity) {
    this.internal.fillOpacity = fillOpacity;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder showPoints(VisibilityMode showPoints) {
    this.internal.showPoints = showPoints;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder pointSize(Double pointSize) {
    this.internal.pointSize = pointSize;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder pointColor(String pointColor) {
    this.internal.pointColor = pointColor;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder axisPlacement(AxisPlacement axisPlacement) {
    this.internal.axisPlacement = axisPlacement;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder axisColorMode(AxisColorMode axisColorMode) {
    this.internal.axisColorMode = axisColorMode;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder axisLabel(String axisLabel) {
    this.internal.axisLabel = axisLabel;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder axisWidth(Double axisWidth) {
    this.internal.axisWidth = axisWidth;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder axisSoftMin(Double axisSoftMin) {
    this.internal.axisSoftMin = axisSoftMin;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder axisSoftMax(Double axisSoftMax) {
    this.internal.axisSoftMax = axisSoftMax;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder axisGridShow(Boolean axisGridShow) {
    this.internal.axisGridShow = axisGridShow;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
    this.internal.scaleDistribution = scaleDistribution.build();
        return this;
    }
    
    public TableSparklineCellOptionsBuilder axisCenteredZero(Boolean axisCenteredZero) {
    this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder barAlignment(BarAlignment barAlignment) {
    this.internal.barAlignment = barAlignment;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder barWidthFactor(Double barWidthFactor) {
    this.internal.barWidthFactor = barWidthFactor;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking) {
    this.internal.stacking = stacking.build();
        return this;
    }
    
    public TableSparklineCellOptionsBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    this.internal.hideFrom = hideFrom.build();
        return this;
    }
    
    public TableSparklineCellOptionsBuilder hideValue(Boolean hideValue) {
    this.internal.hideValue = hideValue;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder transform(GraphTransform transform) {
    this.internal.transform = transform;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder spanNulls(BoolOrFloat64 spanNulls) {
    this.internal.spanNulls = spanNulls;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder fillBelowTo(String fillBelowTo) {
    this.internal.fillBelowTo = fillBelowTo;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder pointSymbol(String pointSymbol) {
    this.internal.pointSymbol = pointSymbol;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder axisBorderShow(Boolean axisBorderShow) {
    this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
    
    public TableSparklineCellOptionsBuilder barMaxWidth(Double barMaxWidth) {
    this.internal.barMaxWidth = barMaxWidth;
        return this;
    }
    public TableSparklineCellOptions build() {
        return this.internal;
    }
}
