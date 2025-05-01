// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

// Thresholds configuration for the panel
public class ThresholdsConfig {
    // Thresholds mode.
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public ThresholdsMode mode;
    // Must be sorted by 'value', first value is always -Infinity
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
