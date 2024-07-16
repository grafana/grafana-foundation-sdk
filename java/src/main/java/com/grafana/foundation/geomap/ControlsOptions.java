// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ControlsOptions {
    // Zoom (upper left) 
    @JsonProperty("showZoom")
    public Boolean showZoom;
    // let the mouse wheel zoom 
    @JsonProperty("mouseWheelZoom")
    public Boolean mouseWheelZoom;
    // Lower right 
    @JsonProperty("showAttribution")
    public Boolean showAttribution;
    // Scale options 
    @JsonProperty("showScale")
    public Boolean showScale;
    // Show debug 
    @JsonProperty("showDebug")
    public Boolean showDebug;
    // Show measure 
    @JsonProperty("showMeasure")
    public Boolean showMeasure;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
