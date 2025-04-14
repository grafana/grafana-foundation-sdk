// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class StreamingQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("bands")
    public Long bands;
    @JsonProperty("noise")
    public Double noise;
    @JsonProperty("speed")
    public Double speed;
    @JsonProperty("spread")
    public Double spread;
    // Possible enum values:
    //  - `"fetch"` 
    //  - `"logs"` 
    //  - `"signal"` 
    //  - `"traces"` 
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public StreamingQueryType type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("url")
    public String url;
    public StreamingQuery() {
    }
    public StreamingQuery(Long bands,Double noise,Double speed,Double spread,StreamingQueryType type,String url) {
        this.bands = bands;
        this.noise = noise;
        this.speed = speed;
        this.spread = spread;
        this.type = type;
        this.url = url;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
