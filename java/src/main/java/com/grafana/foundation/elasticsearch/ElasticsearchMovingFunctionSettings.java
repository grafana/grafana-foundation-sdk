// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class ElasticsearchMovingFunctionSettings {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("window")
    public String window;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("script")
    public InlineScript script;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("shift")
    public String shift;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchMovingFunctionSettings> {
        private final ElasticsearchMovingFunctionSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchMovingFunctionSettings();
        }
    public Builder window(String window) {
    this.internal.window = window;
        return this;
    }
    
    public Builder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    
    public Builder shift(String shift) {
    this.internal.shift = shift;
        return this;
    }
    public ElasticsearchMovingFunctionSettings build() {
            return this.internal;
        }
    }
}
