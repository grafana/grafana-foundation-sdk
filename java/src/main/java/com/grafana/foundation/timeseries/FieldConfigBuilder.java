// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.timeseries;

import com.grafana.foundation.common.GraphDrawStyle;
import com.grafana.foundation.common.GraphGradientMode;
import com.grafana.foundation.common.GraphThresholdsStyleConfig;
import com.grafana.foundation.common.LineInterpolation;
import com.grafana.foundation.common.LineStyle;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.BarAlignment;
import com.grafana.foundation.common.StackingConfig;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.GraphTransform;
import com.grafana.foundation.common.BoolOrFloat64;
import com.grafana.foundation.common.BoolOrUint32;

public class FieldConfigBuilder implements com.grafana.foundation.cog.Builder<FieldConfig> {
    protected final FieldConfig internal;
    
    public FieldConfigBuilder() {
        this.internal = new FieldConfig();
    }
    public FieldConfigBuilder drawStyle(GraphDrawStyle drawStyle) {
        this.internal.drawStyle = drawStyle;
        return this;
    }
    
    public FieldConfigBuilder gradientMode(GraphGradientMode gradientMode) {
        this.internal.gradientMode = gradientMode;
        return this;
    }
    
    public FieldConfigBuilder thresholdsStyle(com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> thresholdsStyle) {
    GraphThresholdsStyleConfig thresholdsStyleResource = thresholdsStyle.build();
        this.internal.thresholdsStyle = thresholdsStyleResource;
        return this;
    }
    
    public FieldConfigBuilder lineColor(String lineColor) {
        this.internal.lineColor = lineColor;
        return this;
    }
    
    public FieldConfigBuilder lineWidth(Double lineWidth) {
        this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public FieldConfigBuilder lineInterpolation(LineInterpolation lineInterpolation) {
        this.internal.lineInterpolation = lineInterpolation;
        return this;
    }
    
    public FieldConfigBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
    LineStyle lineStyleResource = lineStyle.build();
        this.internal.lineStyle = lineStyleResource;
        return this;
    }
    
    public FieldConfigBuilder fillColor(String fillColor) {
        this.internal.fillColor = fillColor;
        return this;
    }
    
    public FieldConfigBuilder fillOpacity(Double fillOpacity) {
        this.internal.fillOpacity = fillOpacity;
        return this;
    }
    
    public FieldConfigBuilder showPoints(VisibilityMode showPoints) {
        this.internal.showPoints = showPoints;
        return this;
    }
    
    public FieldConfigBuilder pointSize(Double pointSize) {
        this.internal.pointSize = pointSize;
        return this;
    }
    
    public FieldConfigBuilder pointColor(String pointColor) {
        this.internal.pointColor = pointColor;
        return this;
    }
    
    public FieldConfigBuilder axisPlacement(AxisPlacement axisPlacement) {
        this.internal.axisPlacement = axisPlacement;
        return this;
    }
    
    public FieldConfigBuilder axisColorMode(AxisColorMode axisColorMode) {
        this.internal.axisColorMode = axisColorMode;
        return this;
    }
    
    public FieldConfigBuilder axisLabel(String axisLabel) {
        this.internal.axisLabel = axisLabel;
        return this;
    }
    
    public FieldConfigBuilder axisWidth(Double axisWidth) {
        this.internal.axisWidth = axisWidth;
        return this;
    }
    
    public FieldConfigBuilder axisSoftMin(Double axisSoftMin) {
        this.internal.axisSoftMin = axisSoftMin;
        return this;
    }
    
    public FieldConfigBuilder axisSoftMax(Double axisSoftMax) {
        this.internal.axisSoftMax = axisSoftMax;
        return this;
    }
    
    public FieldConfigBuilder axisGridShow(Boolean axisGridShow) {
        this.internal.axisGridShow = axisGridShow;
        return this;
    }
    
    public FieldConfigBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
    ScaleDistributionConfig scaleDistributionResource = scaleDistribution.build();
        this.internal.scaleDistribution = scaleDistributionResource;
        return this;
    }
    
    public FieldConfigBuilder axisCenteredZero(Boolean axisCenteredZero) {
        this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    
    public FieldConfigBuilder barAlignment(BarAlignment barAlignment) {
        this.internal.barAlignment = barAlignment;
        return this;
    }
    
    public FieldConfigBuilder barWidthFactor(Double barWidthFactor) {
        this.internal.barWidthFactor = barWidthFactor;
        return this;
    }
    
    public FieldConfigBuilder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking) {
    StackingConfig stackingResource = stacking.build();
        this.internal.stacking = stackingResource;
        return this;
    }
    
    public FieldConfigBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    HideSeriesConfig hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
        return this;
    }
    
    public FieldConfigBuilder transform(GraphTransform transform) {
        this.internal.transform = transform;
        return this;
    }
    
    public FieldConfigBuilder spanNulls(BoolOrFloat64 spanNulls) {
        this.internal.spanNulls = spanNulls;
        return this;
    }
    
    public FieldConfigBuilder fillBelowTo(String fillBelowTo) {
        this.internal.fillBelowTo = fillBelowTo;
        return this;
    }
    
    public FieldConfigBuilder pointSymbol(String pointSymbol) {
        this.internal.pointSymbol = pointSymbol;
        return this;
    }
    
    public FieldConfigBuilder axisBorderShow(Boolean axisBorderShow) {
        this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
    
    public FieldConfigBuilder barMaxWidth(Double barMaxWidth) {
        this.internal.barMaxWidth = barMaxWidth;
        return this;
    }
    
    public FieldConfigBuilder insertNulls(BoolOrUint32 insertNulls) {
        this.internal.insertNulls = insertNulls;
        return this;
    }
    public FieldConfig build() {
        return this.internal;
    }
}
