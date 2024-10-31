// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ElasticsearchInlineScript {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("inline")
    public String inline;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchInlineScript> {
        protected final ElasticsearchInlineScript internal;
        
        public Builder() {
            this.internal = new ElasticsearchInlineScript();
        }
    public Builder inline(String inline) {
    this.internal.inline = inline;
        return this;
    }
    public ElasticsearchInlineScript build() {
            return this.internal;
        }
    }
}
