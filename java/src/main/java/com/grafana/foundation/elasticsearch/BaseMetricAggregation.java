// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class BaseMetricAggregation { 
    @JsonProperty("type")
    public MetricAggregationType type; 
    @JsonProperty("id")
    public String id; 
    @JsonProperty("hide")
    public Boolean hide;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<BaseMetricAggregation> {
        private final BaseMetricAggregation internal;
        
        public Builder() {
            this.internal = new BaseMetricAggregation();
        }
    public Builder type(com.grafana.foundation.cog.Builder<MetricAggregationType> type) {
    this.internal.type = type.build();
        return this;
    }
    
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public BaseMetricAggregation build() {
            return this.internal;
        }
    }
}
