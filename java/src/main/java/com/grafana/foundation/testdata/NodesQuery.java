// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class NodesQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("count")
    public Long count;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("seed")
    public Long seed;
    // Possible enum values:
    //  - `"random"` 
    //  - `"random edges"` 
    //  - `"response_medium"` 
    //  - `"response_small"` 
    //  - `"feature_showcase"` 
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("type")
    public NodesQueryType type;
    public NodesQuery() {
    }
    
    public NodesQuery(Long count,Long seed,NodesQueryType type) {
        this.count = count;
        this.seed = seed;
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
