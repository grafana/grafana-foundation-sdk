// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// TODO docs
public class LineStyle {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fill")
    public LineStyleFill fill;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dash")
    public List<Double> dash;
    public LineStyle() {
    }
    
    public LineStyle(LineStyleFill fill,List<Double> dash) {
        this.fill = fill;
        this.dash = dash;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
