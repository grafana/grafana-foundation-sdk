// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class PipelineVariable {
    @JsonProperty("name")
    public String name;
    @JsonProperty("pipelineAgg")
    public String pipelineAgg;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<PipelineVariable> {
        private final PipelineVariable internal;
        
        public Builder() {
            this.internal = new PipelineVariable();
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder pipelineAgg(String pipelineAgg) {
    this.internal.pipelineAgg = pipelineAgg;
        return this;
    }
    public PipelineVariable build() {
            return this.internal;
        }
    }
}
