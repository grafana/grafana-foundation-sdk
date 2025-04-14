// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class USAQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mode")
    public String mode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("period")
    public String period;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fields")
    public List<String> fields;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("states")
    public List<String> states;
    public USAQuery() {
    }
    public USAQuery(String mode,String period,List<String> fields,List<String> states) {
        this.mode = mode;
        this.period = period;
        this.fields = fields;
        this.states = states;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
