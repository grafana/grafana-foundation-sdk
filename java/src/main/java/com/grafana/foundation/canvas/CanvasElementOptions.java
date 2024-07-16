// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class CanvasElementOptions { 
    @JsonProperty("name")
    public String name; 
    @JsonProperty("type")
    public String type;
    // TODO: figure out how to define this (element config(s)) 
    @JsonProperty("config")
    public Object config; 
    @JsonProperty("constraint")
    public Constraint constraint; 
    @JsonProperty("placement")
    public Placement placement; 
    @JsonProperty("background")
    public BackgroundConfig background; 
    @JsonProperty("border")
    public LineConfig border; 
    @JsonProperty("connections")
    public List<CanvasConnection> connections;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
