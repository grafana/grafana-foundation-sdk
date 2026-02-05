// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ScalarDimensionConfig {
    @JsonProperty("min")
    public Double min;
    @JsonProperty("max")
    public Double max;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fixed")
    public Double fixed;
    // fixed: T -- will be added by each element
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("field")
    public String field;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mode")
    public ScalarDimensionMode mode;
    public ScalarDimensionConfig() {
        this.min = 0.0;
        this.max = 0.0;
    }
    public ScalarDimensionConfig(Double min,Double max,Double fixed,String field,ScalarDimensionMode mode) {
        this.min = min;
        this.max = max;
        this.fixed = fixed;
        this.field = field;
        this.mode = mode;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
