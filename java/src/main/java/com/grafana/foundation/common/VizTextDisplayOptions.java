// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class VizTextDisplayOptions {
    // Explicit title text size
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("titleSize")
    public Double titleSize;
    // Explicit value text size
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("valueSize")
    public Double valueSize;
    public VizTextDisplayOptions() {
    }
    
    public VizTextDisplayOptions(Double titleSize,Double valueSize) {
        this.titleSize = titleSize;
        this.valueSize = valueSize;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
