// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class ThresholdsConfig {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public ThresholdsMode mode;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("steps")
    public List<Threshold> steps;
    public ThresholdsConfig() {
        this.mode = ThresholdsMode.ABSOLUTE;
        this.steps = new LinkedList<>();
    }
    public ThresholdsConfig(ThresholdsMode mode,List<Threshold> steps) {
        this.mode = mode;
        this.steps = steps;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
