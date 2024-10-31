// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class GeoHashGridSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("precision")
    public String precision;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<GeoHashGridSettings> {
        protected final GeoHashGridSettings internal;
        
        public Builder() {
            this.internal = new GeoHashGridSettings();
        }
    public Builder precision(String precision) {
    this.internal.precision = precision;
        return this;
    }
    public GeoHashGridSettings build() {
            return this.internal;
        }
    }
}
