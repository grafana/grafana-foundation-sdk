// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class CanvasElementOptions {
    @JsonProperty("name")
    public String name;
    @JsonProperty("type")
    public String type;
    // TODO: figure out how to define this (element config(s))
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("config")
    public Object config;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("constraint")
    public Constraint constraint;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("placement")
    public Placement placement;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("background")
    public BackgroundConfig background;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("border")
    public LineConfig border;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("connections")
    public List<CanvasConnection> connections;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
