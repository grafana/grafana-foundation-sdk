// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchGeoHashGridSettings { 
    @JsonProperty("precision")
    public String precision;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchGeoHashGridSettings> {
        private final ElasticsearchGeoHashGridSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchGeoHashGridSettings();
        }
    public Builder precision(String precision) {
    this.internal.precision = precision;
        return this;
    }
    public ElasticsearchGeoHashGridSettings build() {
            return this.internal;
        }
    }
}
