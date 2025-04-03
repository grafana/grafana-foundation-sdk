// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ColorDimensionConfig {
    // color value
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fixed")
    public String fixed;
    // fixed: T -- will be added by each element
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("field")
    public String field;
    public ColorDimensionConfig() {
    }
    public ColorDimensionConfig(String fixed,String field) {
        this.fixed = fixed;
        this.field = field;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
