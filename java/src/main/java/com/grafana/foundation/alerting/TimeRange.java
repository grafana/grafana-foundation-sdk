// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Redefining this to avoid an import cycle
public class TimeRange {
    // Redefining this to avoid an import cycle
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("from")
    public String from;
    // Redefining this to avoid an import cycle
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("to")
    public String to;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TimeRange> {
        private final TimeRange internal;
        
        public Builder() {
            this.internal = new TimeRange();
        }
    public Builder from(String from) {
    this.internal.from = from;
        return this;
    }
    
    public Builder to(String to) {
    this.internal.to = to;
        return this;
    }
    public TimeRange build() {
            return this.internal;
        }
    }
}
