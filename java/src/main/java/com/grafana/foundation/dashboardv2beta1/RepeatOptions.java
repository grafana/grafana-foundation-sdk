// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class RepeatOptions {
    @JsonProperty("mode")
    public RepeatMode mode;
    @JsonProperty("value")
    public String value;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("direction")
    public RepeatOptionsDirection direction;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxPerRow")
    public Long maxPerRow;
    public RepeatOptions() {
        this.mode = RepeatMode.VARIABLE;
        this.value = "";
    }
    public RepeatOptions(String value,RepeatOptionsDirection direction,Long maxPerRow) {
        this.mode = RepeatMode.VARIABLE;
        this.value = value;
        this.direction = direction;
        this.maxPerRow = maxPerRow;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
