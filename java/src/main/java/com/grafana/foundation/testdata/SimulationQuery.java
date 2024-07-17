// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import java.util.Map;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class SimulationQuery { 
    @JsonProperty("key")
    public Key key; 
    @JsonProperty("config")
    public Map<String, Object> config; 
    @JsonProperty("stream")
    public Boolean stream; 
    @JsonProperty("last")
    public Boolean last;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<SimulationQuery> {
        private final SimulationQuery internal;
        
        public Builder() {
            this.internal = new SimulationQuery();
        }
    public Builder key(com.grafana.foundation.cog.Builder<Key> key) {
    this.internal.key = key.build();
        return this;
    }
    
    public Builder config(Map<String, Object> config) {
    this.internal.config = config;
        return this;
    }
    
    public Builder stream(Boolean stream) {
    this.internal.stream = stream;
        return this;
    }
    
    public Builder last(Boolean last) {
    this.internal.last = last;
        return this;
    }
    public SimulationQuery build() {
            return this.internal;
        }
    }
}
