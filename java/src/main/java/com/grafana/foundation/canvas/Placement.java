// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Placement { 
    @JsonProperty("top")
    public Double top; 
    @JsonProperty("left")
    public Double left; 
    @JsonProperty("right")
    public Double right; 
    @JsonProperty("bottom")
    public Double bottom; 
    @JsonProperty("width")
    public Double width; 
    @JsonProperty("height")
    public Double height; 
    @JsonProperty("rotation")
    public Double rotation;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
