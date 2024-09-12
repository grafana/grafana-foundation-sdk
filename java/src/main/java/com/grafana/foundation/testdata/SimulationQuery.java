// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class SimulationQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("config")
    public Object config;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("key")
    public Key key;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("last")
    public Boolean last;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("stream")
    public Boolean stream;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<SimulationQuery> {
        private final SimulationQuery internal;
        
        public Builder() {
            this.internal = new SimulationQuery();
        }
    public Builder config(Object config) {
    this.internal.config = config;
        return this;
    }
    
    public Builder key(com.grafana.foundation.cog.Builder<Key> key) {
    this.internal.key = key.build();
        return this;
    }
    
    public Builder last(Boolean last) {
    this.internal.last = last;
        return this;
    }
    
    public Builder stream(Boolean stream) {
    this.internal.stream = stream;
        return this;
    }
    public SimulationQuery build() {
            return this.internal;
        }
    }
}
