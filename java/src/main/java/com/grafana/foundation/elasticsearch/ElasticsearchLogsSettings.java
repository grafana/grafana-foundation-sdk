// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ElasticsearchLogsSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("limit")
    public String limit;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchLogsSettings> {
        private final ElasticsearchLogsSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchLogsSettings();
        }
    public Builder limit(String limit) {
    this.internal.limit = limit;
        return this;
    }
    public ElasticsearchLogsSettings build() {
            return this.internal;
        }
    }
}
