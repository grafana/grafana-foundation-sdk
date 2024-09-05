// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class ElasticsearchSerialDiffSettings {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("lag")
    public String lag;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchSerialDiffSettings> {
        private final ElasticsearchSerialDiffSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchSerialDiffSettings();
        }
    public Builder lag(String lag) {
    this.internal.lag = lag;
        return this;
    }
    public ElasticsearchSerialDiffSettings build() {
            return this.internal;
        }
    }
}
