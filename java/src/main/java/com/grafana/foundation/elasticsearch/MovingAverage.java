// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.Map;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// #MovingAverage's settings are overridden in types.ts
public class MovingAverage { 
    @JsonProperty("pipelineAgg")
    public String pipelineAgg; 
    @JsonProperty("field")
    public String field; 
    @JsonProperty("type")
    public String type; 
    @JsonProperty("id")
    public String id; 
    @JsonProperty("settings")
    public Map<String, Object> settings; 
    @JsonProperty("hide")
    public Boolean hide;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<MovingAverage> {
        private final MovingAverage internal;
        
        public Builder() {
            this.internal = new MovingAverage();
    this.internal.type = "moving_avg";
        }
    public Builder pipelineAgg(String pipelineAgg) {
    this.internal.pipelineAgg = pipelineAgg;
        return this;
    }
    
    public Builder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder settings(Map<String, Object> settings) {
    this.internal.settings = settings;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public MovingAverage build() {
            return this.internal;
        }
    }
}
