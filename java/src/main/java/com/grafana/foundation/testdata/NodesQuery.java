// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class NodesQuery { 
    @JsonProperty("type")
    public NodesQueryType type; 
    @JsonProperty("count")
    public Long count; 
    @JsonProperty("seed")
    public Long seed;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<NodesQuery> {
        private final NodesQuery internal;
        
        public Builder() {
            this.internal = new NodesQuery();
        }
    public Builder type(NodesQueryType type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder count(Long count) {
    this.internal.count = count;
        return this;
    }
    
    public Builder seed(Long seed) {
    this.internal.seed = seed;
        return this;
    }
    public NodesQuery build() {
            return this.internal;
        }
    }
}
