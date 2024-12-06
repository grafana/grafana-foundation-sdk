// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class HistogramSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("interval")
    public String interval;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("min_doc_count")
    public String minDocCount;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<HistogramSettings> {
        protected final HistogramSettings internal;
        
        public Builder() {
            this.internal = new HistogramSettings();
        }
    public Builder interval(String interval) {
    this.internal.interval = interval;
        return this;
    }
    
    public Builder minDocCount(String minDocCount) {
    this.internal.minDocCount = minDocCount;
        return this;
    }
    public HistogramSettings build() {
            return this.internal;
        }
    }
}
