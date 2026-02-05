// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class GeoHashGrid {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("field")
    public String field;
    @JsonProperty("id")
    public String id;
    @JsonProperty("type")
    public BucketAggregationType type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("settings")
    public ElasticsearchGeoHashGridSettings settings;
    public GeoHashGrid() {
        this.id = "";
        this.type = BucketAggregationType.GEOHASH_GRID;
    }
    public GeoHashGrid(String field,String id,ElasticsearchGeoHashGridSettings settings) {
        this.field = field;
        this.id = id;
        this.type = BucketAggregationType.GEOHASH_GRID;
        this.settings = settings;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
