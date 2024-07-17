// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Maps numerical ranges to a display text and color.
// For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
public class RangeMap { 
    @JsonProperty("type")
    public String type;
    // Range to match against and the result to apply when the value is within the range 
    @JsonProperty("options")
    public DashboardRangeMapOptions options;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<RangeMap> {
        private final RangeMap internal;
        
        public Builder() {
            this.internal = new RangeMap();
    this.internal.type = "range";
        }
    public Builder options(com.grafana.foundation.cog.Builder<DashboardRangeMapOptions> options) {
    this.internal.options = options.build();
        return this;
    }
    public RangeMap build() {
            return this.internal;
        }
    }
}
