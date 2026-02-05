// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Nested {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("field")
    public String field;
    @JsonProperty("id")
    public String id;
    @JsonProperty("type")
    public BucketAggregationType type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("settings")
    public Object settings;
    public Nested() {
        this.id = "";
        this.type = BucketAggregationType.NESTED;
    }
    public Nested(String field,String id,Object settings) {
        this.field = field;
        this.id = id;
        this.type = BucketAggregationType.NESTED;
        this.settings = settings;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
