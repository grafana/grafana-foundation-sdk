// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class CanvasOptionsRoot {
    // Name of the root element 
    @JsonProperty("name")
    public String name;
    // Type of root element (frame) 
    @JsonProperty("type")
    public String type;
    // The list of canvas elements attached to the root element 
    @JsonProperty("elements")
    public List<CanvasElementOptions> elements;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
