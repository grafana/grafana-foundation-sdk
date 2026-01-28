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

public class FieldConfigBuilder implements com.grafana.foundation.cog.Builder<FieldConfig> {
    protected final FieldConfig internal;
    
    public FieldConfigBuilder() {
        this.internal = new FieldConfig();
    }
    public FieldConfigBuilder show(ScatterShow show) {
        this.internal.show = show;
        return this;
    }
    
    public FieldConfigBuilder pointSize(com.grafana.foundation.cog.Builder<ScaleDimensionConfig> pointSize) {
    ScaleDimensionConfig pointSizeResource = pointSize.build();
        this.internal.pointSize = pointSizeResource;
        return this;
    }
    
    public FieldConfigBuilder pointColor(com.grafana.foundation.cog.Builder<ColorDimensionConfig> pointColor) {
    ColorDimensionConfig pointColorResource = pointColor.build();
        this.internal.pointColor = pointColorResource;
        return this;
    }
    
    public FieldConfigBuilder lineColor(com.grafana.foundation.cog.Builder<ColorDimensionConfig> lineColor) {
    ColorDimensionConfig lineColorResource = lineColor.build();
        this.internal.lineColor = lineColorResource;
        return this;
    }
    
    public FieldConfigBuilder lineWidth(Integer lineWidth) {
        if (!(lineWidth >= 0)) {
            throw new IllegalArgumentException("lineWidth must be >= 0");
        }
        this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public FieldConfigBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
    LineStyle lineStyleResource = lineStyle.build();
        this.internal.lineStyle = lineStyleResource;
        return this;
    }
    
    public FieldConfigBuilder label(VisibilityMode label) {
        this.internal.label = label;
        return this;
    }
    
    public FieldConfigBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    HideSeriesConfig hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
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
    
    public FieldConfigBuilder labelValue(com.grafana.foundation.cog.Builder<TextDimensionConfig> labelValue) {
    TextDimensionConfig labelValueResource = labelValue.build();
        this.internal.labelValue = labelValueResource;
        return this;
    }
    
    public FieldConfigBuilder axisCenteredZero(Boolean axisCenteredZero) {
        this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    public FieldConfig build() {
        return this.internal;
    }
}
