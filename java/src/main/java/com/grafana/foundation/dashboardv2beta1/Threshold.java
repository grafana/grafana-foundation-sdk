// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Threshold {
    // Value null means -Infinity
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("value")
    public Double value;
    // Optional dashboard-variable expression (e.g. `$myVar`) resolved at render time; `value` is the numeric fallback when the expression cannot be resolved to a single finite number.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("valueExpr")
    public String valueExpr;
    @JsonProperty("color")
    public String color;
    public Threshold() {
        this.color = "";
    }
    public Threshold(Double value,String valueExpr,String color) {
        this.value = value;
        this.valueExpr = valueExpr;
        this.color = color;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
