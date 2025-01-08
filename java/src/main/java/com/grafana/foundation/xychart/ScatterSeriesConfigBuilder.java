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

public class ScatterSeriesConfigBuilder implements com.grafana.foundation.cog.Builder<ScatterSeriesConfig> {
    protected final ScatterSeriesConfig internal;
    
    public ScatterSeriesConfigBuilder() {
        this.internal = new ScatterSeriesConfig();
    }
    public ScatterSeriesConfigBuilder x(String x) {
    this.internal.x = x;
        return this;
    }
    
    public ScatterSeriesConfigBuilder y(String y) {
    this.internal.y = y;
        return this;
    }
    
    public ScatterSeriesConfigBuilder show(ScatterShow show) {
    this.internal.show = show;
        return this;
    }
    
    public ScatterSeriesConfigBuilder pointSize(com.grafana.foundation.cog.Builder<ScaleDimensionConfig> pointSize) {
    this.internal.pointSize = pointSize.build();
        return this;
    }
    
    public ScatterSeriesConfigBuilder pointColor(com.grafana.foundation.cog.Builder<ColorDimensionConfig> pointColor) {
    this.internal.pointColor = pointColor.build();
        return this;
    }
    
    public ScatterSeriesConfigBuilder lineColor(com.grafana.foundation.cog.Builder<ColorDimensionConfig> lineColor) {
    this.internal.lineColor = lineColor.build();
        return this;
    }
    
    public ScatterSeriesConfigBuilder lineWidth(Integer lineWidth) {
        if (!(lineWidth >= 0)) {
            throw new IllegalArgumentException("lineWidth must be >= 0");
        }
    this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public ScatterSeriesConfigBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
    this.internal.lineStyle = lineStyle.build();
        return this;
    }
    
    public ScatterSeriesConfigBuilder label(VisibilityMode label) {
    this.internal.label = label;
        return this;
    }
    
    public ScatterSeriesConfigBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    this.internal.hideFrom = hideFrom.build();
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisPlacement(AxisPlacement axisPlacement) {
    this.internal.axisPlacement = axisPlacement;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisColorMode(AxisColorMode axisColorMode) {
    this.internal.axisColorMode = axisColorMode;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisLabel(String axisLabel) {
    this.internal.axisLabel = axisLabel;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisWidth(Double axisWidth) {
    this.internal.axisWidth = axisWidth;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisSoftMin(Double axisSoftMin) {
    this.internal.axisSoftMin = axisSoftMin;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisSoftMax(Double axisSoftMax) {
    this.internal.axisSoftMax = axisSoftMax;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisGridShow(Boolean axisGridShow) {
    this.internal.axisGridShow = axisGridShow;
        return this;
    }
    
    public ScatterSeriesConfigBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
    this.internal.scaleDistribution = scaleDistribution.build();
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisCenteredZero(Boolean axisCenteredZero) {
    this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    
    public ScatterSeriesConfigBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public ScatterSeriesConfigBuilder labelValue(com.grafana.foundation.cog.Builder<TextDimensionConfig> labelValue) {
    this.internal.labelValue = labelValue.build();
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisBorderShow(Boolean axisBorderShow) {
    this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
    public ScatterSeriesConfig build() {
        return this.internal;
    }
}
