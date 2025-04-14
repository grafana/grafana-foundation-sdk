// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class StreamingQuery {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public StreamingQueryType type;
    @JsonProperty("speed")
    public Integer speed;
    @JsonProperty("spread")
    public Integer spread;
    @JsonProperty("noise")
    public Integer noise;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("bands")
    public Integer bands;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("url")
    public String url;
    public StreamingQuery() {
    }
    public StreamingQuery(StreamingQueryType type,Integer speed,Integer spread,Integer noise,Integer bands,String url) {
        this.type = type;
        this.speed = speed;
        this.spread = spread;
        this.noise = noise;
        this.bands = bands;
        this.url = url;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
