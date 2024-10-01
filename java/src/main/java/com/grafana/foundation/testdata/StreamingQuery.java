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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<StreamingQuery> {
        protected final StreamingQuery internal;
        
        public Builder() {
            this.internal = new StreamingQuery();
        }
    public Builder type(StreamingQueryType type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder speed(Integer speed) {
    this.internal.speed = speed;
        return this;
    }
    
    public Builder spread(Integer spread) {
    this.internal.spread = spread;
        return this;
    }
    
    public Builder noise(Integer noise) {
    this.internal.noise = noise;
        return this;
    }
    
    public Builder bands(Integer bands) {
    this.internal.bands = bands;
        return this;
    }
    
    public Builder url(String url) {
    this.internal.url = url;
        return this;
    }
    public StreamingQuery build() {
            return this.internal;
        }
    }
}
