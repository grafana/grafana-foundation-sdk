// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MovingAverageModelOption {
    @JsonProperty("label")
    public String label;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("value")
    public MovingAverageModel value;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<MovingAverageModelOption> {
        private final MovingAverageModelOption internal;
        
        public Builder() {
            this.internal = new MovingAverageModelOption();
        }
    public Builder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public Builder value(MovingAverageModel value) {
    this.internal.value = value;
        return this;
    }
    public MovingAverageModelOption build() {
            return this.internal;
        }
    }
}
