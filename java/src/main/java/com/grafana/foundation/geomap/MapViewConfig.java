// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MapViewConfig {
    @JsonProperty("id")
    public String id;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lat")
    public Long lat;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lon")
    public Long lon;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("zoom")
    public Long zoom;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("minZoom")
    public Long minZoom;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxZoom")
    public Long maxZoom;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("padding")
    public Long padding;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("allLayers")
    public Boolean allLayers;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lastOnly")
    public Boolean lastOnly;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("layer")
    public String layer;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("shared")
    public Boolean shared;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<MapViewConfig> {
        protected final MapViewConfig internal;
        
        public Builder() {
            this.internal = new MapViewConfig();
        this.id("zero");
        this.lat(0L);
        this.lon(0L);
        this.zoom(1L);
        this.allLayers(true);
        }
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder lat(Long lat) {
    this.internal.lat = lat;
        return this;
    }
    
    public Builder lon(Long lon) {
    this.internal.lon = lon;
        return this;
    }
    
    public Builder zoom(Long zoom) {
    this.internal.zoom = zoom;
        return this;
    }
    
    public Builder minZoom(Long minZoom) {
    this.internal.minZoom = minZoom;
        return this;
    }
    
    public Builder maxZoom(Long maxZoom) {
    this.internal.maxZoom = maxZoom;
        return this;
    }
    
    public Builder padding(Long padding) {
    this.internal.padding = padding;
        return this;
    }
    
    public Builder allLayers(Boolean allLayers) {
    this.internal.allLayers = allLayers;
        return this;
    }
    
    public Builder lastOnly(Boolean lastOnly) {
    this.internal.lastOnly = lastOnly;
        return this;
    }
    
    public Builder layer(String layer) {
    this.internal.layer = layer;
        return this;
    }
    
    public Builder shared(Boolean shared) {
    this.internal.shared = shared;
        return this;
    }
    public MapViewConfig build() {
            return this.internal;
        }
    }
}
