// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class SimulationQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("config")
    public Object config;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("key")
    public Key key;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("last")
    public Boolean last;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("stream")
    public Boolean stream;
    public SimulationQuery() {
        this.key = new com.grafana.foundation.testdata.KeyBuilder().build();
    }
    public SimulationQuery(Object config,Key key,Boolean last,Boolean stream) {
        this.config = config;
        this.key = key;
        this.last = last;
        this.stream = stream;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
