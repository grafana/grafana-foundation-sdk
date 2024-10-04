// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.ScaleDimensionConfig;
import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.LineStyle;
import com.grafana.foundation.common.VisibilityMode;
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
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labelValue")
    public TextDimensionConfig labelValue;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisCenteredZero")
    public Boolean axisCenteredZero;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ScatterSeriesConfig> {
        protected final ScatterSeriesConfig internal;
        
        public Builder() {
            this.internal = new ScatterSeriesConfig();
        this.show(ScatterShow.POINTS);
        this.label(VisibilityMode.AUTO);
        }
    public Builder x(String x) {
    this.internal.x = x;
        return this;
    }
    
    public Builder y(String y) {
    this.internal.y = y;
        return this;
    }
    
    public Builder show(ScatterShow show) {
    this.internal.show = show;
        return this;
    }
    
    public Builder pointSize(com.grafana.foundation.cog.Builder<ScaleDimensionConfig> pointSize) {
    this.internal.pointSize = pointSize.build();
        return this;
    }
    
    public Builder pointColor(com.grafana.foundation.cog.Builder<ColorDimensionConfig> pointColor) {
    this.internal.pointColor = pointColor.build();
        return this;
    }
    
    public Builder lineColor(com.grafana.foundation.cog.Builder<ColorDimensionConfig> lineColor) {
    this.internal.lineColor = lineColor.build();
        return this;
    }
    
    public Builder lineWidth(Integer lineWidth) {
        if (!(lineWidth >= 0)) {
            throw new IllegalArgumentException("lineWidth must be >= 0");
        }
    this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public Builder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
    this.internal.lineStyle = lineStyle.build();
        return this;
    }
    
    public Builder label(VisibilityMode label) {
    this.internal.label = label;
        return this;
    }
    
    public Builder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    this.internal.hideFrom = hideFrom.build();
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
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder labelValue(com.grafana.foundation.cog.Builder<TextDimensionConfig> labelValue) {
    this.internal.labelValue = labelValue.build();
        return this;
    }
    
    public Builder axisCenteredZero(Boolean axisCenteredZero) {
    this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    public ScatterSeriesConfig build() {
            return this.internal;
        }
    }
}
