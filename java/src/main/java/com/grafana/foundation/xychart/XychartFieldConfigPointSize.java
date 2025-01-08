// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class XychartFieldConfigPointSize {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fixed")
    public Integer fixed;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("min")
    public Integer min;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("max")
    public Integer max;
    public XychartFieldConfigPointSize() {
    }
    
    public XychartFieldConfigPointSize(Integer fixed,Integer min,Integer max) {
        this.fixed = fixed;
        this.min = min;
        this.max = max;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
