// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Dashboardv2beta1ActionStyle {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("backgroundColor")
    public String backgroundColor;
    public Dashboardv2beta1ActionStyle() {
    }
    public Dashboardv2beta1ActionStyle(String backgroundColor) {
        this.backgroundColor = backgroundColor;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
