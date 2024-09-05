// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class MapLayerOptions {
    @JsonProperty("type")
    public String type;
    // configured unique display name
    @JsonProperty("name")
    public String name;
    // Custom options depending on the type
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("config")
    public Object config;
    // Common method to define geometry fields
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("location")
    public FrameGeometrySource location;
    // Defines a frame MatcherConfig that may filter data for the given layer
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("filterData")
    public Object filterData;
    // Common properties:
    // https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
    // Layer opacity (0-1)
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("opacity")
    public Long opacity;
    // Check tooltip (defaults to true)
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("tooltip")
    public Boolean tooltip;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<MapLayerOptions> {
        private final MapLayerOptions internal;
        
        public Builder() {
            this.internal = new MapLayerOptions();
        }
    public Builder type(String type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder config(Object config) {
    this.internal.config = config;
        return this;
    }
    
    public Builder location(com.grafana.foundation.cog.Builder<FrameGeometrySource> location) {
    this.internal.location = location.build();
        return this;
    }
    
    public Builder filterData(Object filterData) {
    this.internal.filterData = filterData;
        return this;
    }
    
    public Builder opacity(Long opacity) {
    this.internal.opacity = opacity;
        return this;
    }
    
    public Builder tooltip(Boolean tooltip) {
    this.internal.tooltip = tooltip;
        return this;
    }
    public MapLayerOptions build() {
            return this.internal;
        }
    }
}
