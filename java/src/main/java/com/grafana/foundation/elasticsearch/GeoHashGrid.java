// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class GeoHashGrid { 
    @JsonProperty("field")
    public String field; 
    @JsonProperty("id")
    public String id; 
    @JsonProperty("type")
    public String type; 
    @JsonProperty("settings")
    public ElasticsearchGeoHashGridSettings settings;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<GeoHashGrid> {
        private final GeoHashGrid internal;
        
        public Builder() {
            this.internal = new GeoHashGrid();
    this.internal.type = "geohash_grid";
        }
    public Builder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder settings(com.grafana.foundation.cog.Builder<ElasticsearchGeoHashGridSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    public GeoHashGrid build() {
            return this.internal;
        }
    }
}
