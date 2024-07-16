// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class InlineScript { 
    @JsonProperty("String")
    public String string; 
    @JsonProperty("ElasticsearchInlineScript")
    public ElasticsearchInlineScript elasticsearchInlineScript;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<InlineScript> {
        private final InlineScript internal;
        
        public Builder() {
            this.internal = new InlineScript();
        }
    public Builder string(String string) {
    this.internal.string = string;
        return this;
    }
    
    public Builder elasticsearchInlineScript(com.grafana.foundation.cog.Builder<ElasticsearchInlineScript> elasticsearchInlineScript) {
    this.internal.elasticsearchInlineScript = elasticsearchInlineScript.build();
        return this;
    }
    public InlineScript build() {
            return this.internal;
        }
    }
}
