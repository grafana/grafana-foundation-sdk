// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class BucketAggregationWithField { 
    @JsonProperty("field")
    public String field; 
    @JsonProperty("id")
    public String id; 
    @JsonProperty("type")
    public BucketAggregationType type; 
    @JsonProperty("settings")
    public Object settings;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<BucketAggregationWithField> {
        private final BucketAggregationWithField internal;
        
        public Builder() {
            this.internal = new BucketAggregationWithField();
        }
    public Builder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder type(BucketAggregationType type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder settings(Object settings) {
    this.internal.settings = settings;
        return this;
    }
    public BucketAggregationWithField build() {
            return this.internal;
        }
    }
}
