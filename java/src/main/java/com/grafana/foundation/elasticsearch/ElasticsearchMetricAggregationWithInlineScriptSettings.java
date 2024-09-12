// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ElasticsearchMetricAggregationWithInlineScriptSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("script")
    public InlineScript script;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchMetricAggregationWithInlineScriptSettings> {
        private final ElasticsearchMetricAggregationWithInlineScriptSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchMetricAggregationWithInlineScriptSettings();
        }
    public Builder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    public ElasticsearchMetricAggregationWithInlineScriptSettings build() {
            return this.internal;
        }
    }
}
