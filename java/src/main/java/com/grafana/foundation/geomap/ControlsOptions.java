// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class ControlsOptions {
    // Zoom (upper left)
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("showZoom")
    public Boolean showZoom;
    // let the mouse wheel zoom
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("mouseWheelZoom")
    public Boolean mouseWheelZoom;
    // Lower right
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("showAttribution")
    public Boolean showAttribution;
    // Scale options
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("showScale")
    public Boolean showScale;
    // Show debug
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("showDebug")
    public Boolean showDebug;
    // Show measure
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("showMeasure")
    public Boolean showMeasure;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
