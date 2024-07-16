// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class BaseMovingAverageModelSettings { 
    @JsonProperty("model")
    public MovingAverageModel model; 
    @JsonProperty("window")
    public String window; 
    @JsonProperty("predict")
    public String predict;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<BaseMovingAverageModelSettings> {
        private final BaseMovingAverageModelSettings internal;
        
        public Builder() {
            this.internal = new BaseMovingAverageModelSettings();
        }
    public Builder model(MovingAverageModel model) {
    this.internal.model = model;
        return this;
    }
    
    public Builder window(String window) {
    this.internal.window = window;
        return this;
    }
    
    public Builder predict(String predict) {
    this.internal.predict = predict;
        return this;
    }
    public BaseMovingAverageModelSettings build() {
            return this.internal;
        }
    }
}
