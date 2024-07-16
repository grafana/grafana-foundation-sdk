// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class MapViewConfig { 
    @JsonProperty("id")
    public String id; 
    @JsonProperty("lat")
    public Long lat; 
    @JsonProperty("lon")
    public Long lon; 
    @JsonProperty("zoom")
    public Long zoom; 
    @JsonProperty("minZoom")
    public Long minZoom; 
    @JsonProperty("maxZoom")
    public Long maxZoom; 
    @JsonProperty("padding")
    public Long padding; 
    @JsonProperty("allLayers")
    public Boolean allLayers; 
    @JsonProperty("lastOnly")
    public Boolean lastOnly; 
    @JsonProperty("layer")
    public String layer; 
    @JsonProperty("shared")
    public Boolean shared;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
