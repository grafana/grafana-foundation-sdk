// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class ElasticsearchRawDocumentSettings {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("size")
    public String size;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchRawDocumentSettings> {
        private final ElasticsearchRawDocumentSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchRawDocumentSettings();
        }
    public Builder size(String size) {
    this.internal.size = size;
        return this;
    }
    public ElasticsearchRawDocumentSettings build() {
            return this.internal;
        }
    }
}
