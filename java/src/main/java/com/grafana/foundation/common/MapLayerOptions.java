// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MapLayerOptions {
    @JsonProperty("type")
    public String type;
    // configured unique display name
    @JsonProperty("name")
    public String name;
    // Custom options depending on the type
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("config")
    public Object config;
    // Common method to define geometry fields
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("location")
    public FrameGeometrySource location;
    // Defines a frame MatcherConfig that may filter data for the given layer
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filterData")
    public Object filterData;
    // Common properties:
    // https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
    // Layer opacity (0-1)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("opacity")
    public Long opacity;
    // Check tooltip (defaults to true)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("tooltip")
    public Boolean tooltip;
    public MapLayerOptions() {
    }
    
    public MapLayerOptions(String type,String name,Object config,FrameGeometrySource location,Object filterData,Long opacity,Boolean tooltip) {
        this.type = type;
        this.name = name;
        this.config = config;
        this.location = location;
        this.filterData = filterData;
        this.opacity = opacity;
        this.tooltip = tooltip;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
