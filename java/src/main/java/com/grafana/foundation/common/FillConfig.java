// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class FillConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillColor")
    public String fillColor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillOpacity")
    public Double fillOpacity;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillBelowTo")
    public String fillBelowTo;
    public FillConfig() {
    }
    public FillConfig(String fillColor,Double fillOpacity,String fillBelowTo) {
        this.fillColor = fillColor;
        this.fillOpacity = fillOpacity;
        this.fillBelowTo = fillBelowTo;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
