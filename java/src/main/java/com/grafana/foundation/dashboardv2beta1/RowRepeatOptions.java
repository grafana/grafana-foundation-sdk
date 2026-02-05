// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class RowRepeatOptions {
    @JsonProperty("mode")
    public RepeatMode mode;
    @JsonProperty("value")
    public String value;
    public RowRepeatOptions() {
        this.mode = RepeatMode.VARIABLE;
        this.value = "";
    }
    public RowRepeatOptions(String value) {
        this.mode = RepeatMode.VARIABLE;
        this.value = value;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
