// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Controls various color options
public class HeatmapColorOptions {
    // Sets the color mode 
    @JsonProperty("mode")
    public HeatmapColorMode mode;
    // Controls the color scheme used 
    @JsonProperty("scheme")
    public String scheme;
    // Controls the color fill when in opacity mode 
    @JsonProperty("fill")
    public String fill;
    // Controls the color scale 
    @JsonProperty("scale")
    public HeatmapColorScale scale;
    // Controls the exponent when scale is set to exponential 
    @JsonProperty("exponent")
    public Float exponent;
    // Controls the number of color steps 
    @JsonProperty("steps")
    public Long steps;
    // Reverses the color scheme 
    @JsonProperty("reverse")
    public Boolean reverse;
    // Sets the minimum value for the color scale 
    @JsonProperty("min")
    public Float min;
    // Sets the maximum value for the color scale 
    @JsonProperty("max")
    public Float max;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
