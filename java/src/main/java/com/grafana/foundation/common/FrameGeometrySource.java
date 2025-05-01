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
    public FrameGeometrySource() {
        this.mode = FrameGeometrySourceMode.AUTO;
    }
    public FrameGeometrySource(FrameGeometrySourceMode mode,String geohash,String latitude,String longitude,String wkt,String lookup,String gazetteer) {
        this.mode = mode;
        this.geohash = geohash;
        this.latitude = latitude;
        this.longitude = longitude;
        this.wkt = wkt;
        this.lookup = lookup;
        this.gazetteer = gazetteer;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
