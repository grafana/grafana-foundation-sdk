// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchPercentilesSettings { 
    @JsonProperty("script")
    public InlineScript script; 
    @JsonProperty("missing")
    public String missing; 
    @JsonProperty("percents")
    public List<String> percents;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchPercentilesSettings> {
        private final ElasticsearchPercentilesSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchPercentilesSettings();
        }
    public Builder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    
    public Builder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    
    public Builder percents(List<String> percents) {
    this.internal.percents = percents;
        return this;
    }
    public ElasticsearchPercentilesSettings build() {
            return this.internal;
        }
    }
}
