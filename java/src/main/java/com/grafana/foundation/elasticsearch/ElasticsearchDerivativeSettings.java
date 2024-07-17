// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchDerivativeSettings { 
    @JsonProperty("unit")
    public String unit;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchDerivativeSettings> {
        private final ElasticsearchDerivativeSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchDerivativeSettings();
        }
    public Builder unit(String unit) {
    this.internal.unit = unit;
        return this;
    }
    public ElasticsearchDerivativeSettings build() {
            return this.internal;
        }
    }
}
