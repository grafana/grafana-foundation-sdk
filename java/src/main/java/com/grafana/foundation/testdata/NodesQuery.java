// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class NodesQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("type")
    public NodesQueryType type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("count")
    public Long count;
    public NodesQuery() {
    }
    
    public NodesQuery(NodesQueryType type,Long count) {
        this.type = type;
        this.count = count;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
