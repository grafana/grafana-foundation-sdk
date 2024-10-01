// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class FrameGeometrySource {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public FrameGeometrySourceMode mode;
    // Field mappings
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("geohash")
    public String geohash;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("latitude")
    public String latitude;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("longitude")
    public String longitude;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("wkt")
    public String wkt;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lookup")
    public String lookup;
    // Path to Gazetteer
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("gazetteer")
    public String gazetteer;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<FrameGeometrySource> {
        protected final FrameGeometrySource internal;
        
        public Builder() {
            this.internal = new FrameGeometrySource();
        }
    public Builder mode(FrameGeometrySourceMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder geohash(String geohash) {
    this.internal.geohash = geohash;
        return this;
    }
    
    public Builder latitude(String latitude) {
    this.internal.latitude = latitude;
        return this;
    }
    
    public Builder longitude(String longitude) {
    this.internal.longitude = longitude;
        return this;
    }
    
    public Builder wkt(String wkt) {
    this.internal.wkt = wkt;
        return this;
    }
    
    public Builder lookup(String lookup) {
    this.internal.lookup = lookup;
        return this;
    }
    
    public Builder gazetteer(String gazetteer) {
    this.internal.gazetteer = gazetteer;
        return this;
    }
    public FrameGeometrySource build() {
            return this.internal;
        }
    }
}
