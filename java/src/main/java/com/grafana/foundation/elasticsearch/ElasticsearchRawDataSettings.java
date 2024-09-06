// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ElasticsearchRawDataSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("size")
    public String size;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchRawDataSettings> {
        private final ElasticsearchRawDataSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchRawDataSettings();
        }
    public Builder size(String size) {
    this.internal.size = size;
        return this;
    }
    public ElasticsearchRawDataSettings build() {
            return this.internal;
        }
    }
}
