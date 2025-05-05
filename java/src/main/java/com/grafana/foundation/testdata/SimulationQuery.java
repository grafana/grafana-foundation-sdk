// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;

public class SimulationQuery {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("key")
    public Key key;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("config")
    public Map<String, Object> config;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("stream")
    public Boolean stream;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("last")
    public Boolean last;
    public SimulationQuery() {
        this.key = new com.grafana.foundation.testdata.Key();
    }
    public SimulationQuery(Key key,Map<String, Object> config,Boolean stream,Boolean last) {
        this.key = key;
        this.config = config;
        this.stream = stream;
        this.last = last;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
