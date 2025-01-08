// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.LineStyle;

public class FieldConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("show")
    public XYShowMode show;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointSize")
    public XychartFieldConfigPointSize pointSize;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointShape")
    public PointShape pointShape;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointStrokeWidth")
    public Integer pointStrokeWidth;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillOpacity")
    public Integer fillOpacity;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lineWidth")
    public Integer lineWidth;
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
    @JsonProperty("lineStyle")
    public LineStyle lineStyle;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisBorderShow")
    public Boolean axisBorderShow;
    public FieldConfig() {
        this.show = XYShowMode.POINTS;
        this.fillOpacity = 50;
    }
    
    public FieldConfig(XYShowMode show,XychartFieldConfigPointSize pointSize,PointShape pointShape,Integer pointStrokeWidth,Integer fillOpacity,Integer lineWidth,HideSeriesConfig hideFrom,AxisPlacement axisPlacement,AxisColorMode axisColorMode,String axisLabel,Double axisWidth,Double axisSoftMin,Double axisSoftMax,Boolean axisGridShow,ScaleDistributionConfig scaleDistribution,Boolean axisCenteredZero,LineStyle lineStyle,Boolean axisBorderShow) {
        this.show = show;
        this.pointSize = pointSize;
        this.pointShape = pointShape;
        this.pointStrokeWidth = pointStrokeWidth;
        this.fillOpacity = fillOpacity;
        this.lineWidth = lineWidth;
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
        this.lineStyle = lineStyle;
        this.axisBorderShow = axisBorderShow;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
