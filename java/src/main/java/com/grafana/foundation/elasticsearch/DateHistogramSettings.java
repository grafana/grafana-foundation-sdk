// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class DateHistogramSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("interval")
    public String interval;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("min_doc_count")
    public String minDocCount;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("trimEdges")
    public String trimEdges;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("offset")
    public String offset;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeZone")
    public String timeZone;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DateHistogramSettings> {
        protected final DateHistogramSettings internal;
        
        public Builder() {
            this.internal = new DateHistogramSettings();
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
    public DateHistogramSettings build() {
            return this.internal;
        }
    }
}
