// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.nodegraph;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ArcOption {
    // Field from which to get the value. Values should be less than 1, representing fraction of a circle. 
    @JsonProperty("field")
    public String field;
    // The color of the arc. 
    @JsonProperty("color")
    public String color;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}