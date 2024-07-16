// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.histogram;

import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.GraphGradientMode;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class FieldConfig {
    // Controls line width of the bars. 
    @JsonProperty("lineWidth")
    public Integer lineWidth;
    // Controls the fill opacity of the bars. 
    @JsonProperty("fillOpacity")
    public Integer fillOpacity; 
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
    @JsonProperty("hideFrom")
    public HideSeriesConfig hideFrom;
    // Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
    // Gradient appearance is influenced by the Fill opacity setting. 
    @JsonProperty("gradientMode")
    public GraphGradientMode gradientMode; 
    @JsonProperty("axisBorderShow")
    public Boolean axisBorderShow;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
