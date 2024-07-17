// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchUniqueCountSettings { 
    @JsonProperty("precision_threshold")
    public String precisionThreshold; 
    @JsonProperty("missing")
    public String missing;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchUniqueCountSettings> {
        private final ElasticsearchUniqueCountSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchUniqueCountSettings();
        }
    public Builder precisionThreshold(String precisionThreshold) {
    this.internal.precisionThreshold = precisionThreshold;
        return this;
    }
    
    public Builder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    public ElasticsearchUniqueCountSettings build() {
            return this.internal;
        }
    }
}
