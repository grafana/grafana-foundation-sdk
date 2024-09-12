// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// User-defined value for a metric that triggers visual changes in a panel when this value is met or exceeded
// They are used to conditionally style and color visualizations based on query results , and can be applied to most visualizations.
public class Threshold {
    // Value represents a specified metric for the threshold, which triggers a visual change in the dashboard when this value is met or exceeded.
    // Nulls currently appear here when serializing -Infinity to JSON.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("value")
    public Double value;
    // Color represents the color of the visual change that will occur in the dashboard when the threshold value is met or exceeded.
    @JsonProperty("color")
    public String color;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
