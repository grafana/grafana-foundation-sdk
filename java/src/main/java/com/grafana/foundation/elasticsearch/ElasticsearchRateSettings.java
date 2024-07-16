// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchRateSettings { 
    @JsonProperty("unit")
    public String unit; 
    @JsonProperty("mode")
    public String mode;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchRateSettings> {
        private final ElasticsearchRateSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchRateSettings();
        }
    public Builder unit(String unit) {
    this.internal.unit = unit;
        return this;
    }
    
    public Builder mode(String mode) {
    this.internal.mode = mode;
        return this;
    }
    public ElasticsearchRateSettings build() {
            return this.internal;
        }
    }
}