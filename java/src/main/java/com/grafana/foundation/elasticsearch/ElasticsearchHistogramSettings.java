// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchHistogramSettings { 
    @JsonProperty("interval")
    public String interval; 
    @JsonProperty("min_doc_count")
    public String minDocCount;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchHistogramSettings> {
        private final ElasticsearchHistogramSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchHistogramSettings();
        }
    public Builder interval(String interval) {
    this.internal.interval = interval;
        return this;
    }
    
    public Builder minDocCount(String minDocCount) {
    this.internal.minDocCount = minDocCount;
        return this;
    }
    public ElasticsearchHistogramSettings build() {
            return this.internal;
        }
    }
}
