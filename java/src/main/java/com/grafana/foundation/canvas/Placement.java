// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Placement {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("top")
    public Double top;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("left")
    public Double left;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("right")
    public Double right;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("bottom")
    public Double bottom;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("width")
    public Double width;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("height")
    public Double height;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
