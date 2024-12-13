// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class DashboardRangeMapOptions {
    // Min value of the range. It can be null which means -Infinity
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("from")
    public Double from;
    // Max value of the range. It can be null which means +Infinity
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("to")
    public Double to;
    // Config to apply when the value is within the range
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("result")
    public ValueMappingResult result;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DashboardRangeMapOptions> {
        protected final DashboardRangeMapOptions internal;
        
        public Builder() {
            this.internal = new DashboardRangeMapOptions();
        }
    public Builder from(Double from) {
    this.internal.from = from;
        return this;
    }
    
    public Builder to(Double to) {
    this.internal.to = to;
        return this;
    }
    
    public Builder result(ValueMappingResult result) {
    this.internal.result = result;
        return this;
    }
    public DashboardRangeMapOptions build() {
            return this.internal;
        }
    }
}
