// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchDateHistogramSettings { 
    @JsonProperty("interval")
    public String interval; 
    @JsonProperty("min_doc_count")
    public String minDocCount; 
    @JsonProperty("trimEdges")
    public String trimEdges; 
    @JsonProperty("offset")
    public String offset; 
    @JsonProperty("timeZone")
    public String timeZone;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchDateHistogramSettings> {
        private final ElasticsearchDateHistogramSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchDateHistogramSettings();
        }
    public Builder interval(String interval) {
    this.internal.interval = interval;
        return this;
    }
    
    public Builder minDocCount(String minDocCount) {
    this.internal.minDocCount = minDocCount;
        return this;
    }
    
    public Builder trimEdges(String trimEdges) {
    this.internal.trimEdges = trimEdges;
        return this;
    }
    
    public Builder offset(String offset) {
    this.internal.offset = offset;
        return this;
    }
    
    public Builder timeZone(String timeZone) {
    this.internal.timeZone = timeZone;
        return this;
    }
    public ElasticsearchDateHistogramSettings build() {
            return this.internal;
        }
    }
}
