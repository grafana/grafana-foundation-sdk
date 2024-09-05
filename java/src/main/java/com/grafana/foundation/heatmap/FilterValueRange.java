// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Controls the value filter range
public class FilterValueRange {
    // Sets the filter range to values less than or equal to the given value
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("le")
    public Float le;
    // Sets the filter range to values greater than or equal to the given value
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("ge")
    public Float ge;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
