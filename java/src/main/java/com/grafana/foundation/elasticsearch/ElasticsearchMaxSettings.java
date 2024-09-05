// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class ElasticsearchMaxSettings {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("script")
    public InlineScript script;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("missing")
    public String missing;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchMaxSettings> {
        private final ElasticsearchMaxSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchMaxSettings();
        }
    public Builder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    
    public Builder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    public ElasticsearchMaxSettings build() {
            return this.internal;
        }
    }
}
