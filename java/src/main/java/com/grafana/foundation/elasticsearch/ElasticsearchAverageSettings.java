// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchAverageSettings { 
    @JsonProperty("script")
    public InlineScript script; 
    @JsonProperty("missing")
    public String missing;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchAverageSettings> {
        private final ElasticsearchAverageSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchAverageSettings();
        }
    public Builder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    
    public Builder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    public ElasticsearchAverageSettings build() {
            return this.internal;
        }
    }
}
