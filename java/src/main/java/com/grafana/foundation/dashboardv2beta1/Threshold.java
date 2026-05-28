// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Threshold {
    // Value null means -Infinity
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("value")
    public Double value;
    @JsonProperty("color")
    public String color;
    public Threshold() {
        this.color = "";
    }
    public Threshold(Double value,String color) {
        this.value = value;
        this.color = color;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
