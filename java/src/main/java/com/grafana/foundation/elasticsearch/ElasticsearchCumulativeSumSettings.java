// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchCumulativeSumSettings { 
    @JsonProperty("format")
    public String format;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchCumulativeSumSettings> {
        private final ElasticsearchCumulativeSumSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchCumulativeSumSettings();
        }
    public Builder format(String format) {
    this.internal.format = format;
        return this;
    }
    public ElasticsearchCumulativeSumSettings build() {
            return this.internal;
        }
    }
}
