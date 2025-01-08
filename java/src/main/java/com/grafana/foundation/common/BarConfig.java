// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class BarConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("barAlignment")
    public BarAlignment barAlignment;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("barWidthFactor")
    public Double barWidthFactor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("barMaxWidth")
    public Double barMaxWidth;
    public BarConfig() {
    }
    
    public BarConfig(BarAlignment barAlignment,Double barWidthFactor,Double barMaxWidth) {
        this.barAlignment = barAlignment;
        this.barWidthFactor = barWidthFactor;
        this.barMaxWidth = barMaxWidth;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
