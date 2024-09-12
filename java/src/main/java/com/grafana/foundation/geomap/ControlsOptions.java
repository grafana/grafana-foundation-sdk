// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ControlsOptions {
    // Zoom (upper left)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showZoom")
    public Boolean showZoom;
    // let the mouse wheel zoom
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mouseWheelZoom")
    public Boolean mouseWheelZoom;
    // Lower right
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showAttribution")
    public Boolean showAttribution;
    // Scale options
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showScale")
    public Boolean showScale;
    // Show debug
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showDebug")
    public Boolean showDebug;
    // Show measure
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showMeasure")
    public Boolean showMeasure;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
