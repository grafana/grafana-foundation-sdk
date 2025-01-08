// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.barchart;

import com.grafana.foundation.common.GraphGradientMode;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.GraphThresholdsStyleConfig;

public class FieldConfig {
    // Controls line width of the bars.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lineWidth")
    public Integer lineWidth;
    // Controls the fill opacity of the bars.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillOpacity")
    public Integer fillOpacity;
    // Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
    // Gradient appearance is influenced by the Fill opacity setting.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("gradientMode")
    public GraphGradientMode gradientMode;
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
    @JsonProperty("hideFrom")
    public HideSeriesConfig hideFrom;
    // Threshold rendering
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("thresholdsStyle")
    public GraphThresholdsStyleConfig thresholdsStyle;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("axisBorderShow")
    public Boolean axisBorderShow;
    public FieldConfig() {
        this.lineWidth = 1;
        this.fillOpacity = 80;
        this.gradientMode = GraphGradientMode.NONE;
    }
    
    public FieldConfig(Integer lineWidth,Integer fillOpacity,GraphGradientMode gradientMode,AxisPlacement axisPlacement,AxisColorMode axisColorMode,String axisLabel,Double axisWidth,Double axisSoftMin,Double axisSoftMax,Boolean axisGridShow,ScaleDistributionConfig scaleDistribution,Boolean axisCenteredZero,HideSeriesConfig hideFrom,GraphThresholdsStyleConfig thresholdsStyle,Boolean axisBorderShow) {
        this.lineWidth = lineWidth;
        this.fillOpacity = fillOpacity;
        this.gradientMode = gradientMode;
        this.axisPlacement = axisPlacement;
        this.axisColorMode = axisColorMode;
        this.axisLabel = axisLabel;
        this.axisWidth = axisWidth;
        this.axisSoftMin = axisSoftMin;
        this.axisSoftMax = axisSoftMax;
        this.axisGridShow = axisGridShow;
        this.scaleDistribution = scaleDistribution;
        this.axisCenteredZero = axisCenteredZero;
        this.hideFrom = hideFrom;
        this.thresholdsStyle = thresholdsStyle;
        this.axisBorderShow = axisBorderShow;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
