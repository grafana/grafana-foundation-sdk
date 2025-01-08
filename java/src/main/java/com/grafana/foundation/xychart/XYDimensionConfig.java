// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// Configuration for the Table/Auto mode
public class XYDimensionConfig {
    @JsonProperty("frame")
    public Integer frame;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("x")
    public String x;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("exclude")
    public List<String> exclude;
    public XYDimensionConfig() {
    }
    
    public XYDimensionConfig(Integer frame,String x,List<String> exclude) {
        this.frame = frame;
        this.x = x;
        this.exclude = exclude;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
