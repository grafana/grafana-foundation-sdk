// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BaseBucketAggregation {
    @JsonProperty("id")
    public String id;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public BucketAggregationType type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("settings")
    public Object settings;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<BaseBucketAggregation> {
        protected final BaseBucketAggregation internal;
        
        public Builder() {
            this.internal = new BaseBucketAggregation();
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
    public BaseBucketAggregation build() {
            return this.internal;
        }
    }
}
