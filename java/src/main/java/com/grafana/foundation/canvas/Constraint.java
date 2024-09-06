// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Constraint {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("horizontal")
    public HorizontalConstraint horizontal;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("vertical")
    public VerticalConstraint vertical;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
