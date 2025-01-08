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
    public BaseBucketAggregation() {
    }
    
    public BaseBucketAggregation(String id,BucketAggregationType type,Object settings) {
        this.id = id;
        this.type = type;
        this.settings = settings;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
