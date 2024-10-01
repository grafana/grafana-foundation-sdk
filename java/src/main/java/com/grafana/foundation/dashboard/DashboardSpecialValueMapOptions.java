// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class DashboardSpecialValueMapOptions {
    // Special value to match against
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("match")
    public SpecialValueMatch match;
    // Config to apply when the value matches the special value
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("result")
    public ValueMappingResult result;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DashboardSpecialValueMapOptions> {
        protected final DashboardSpecialValueMapOptions internal;
        
        public Builder() {
            this.internal = new DashboardSpecialValueMapOptions();
        }
    public Builder match(SpecialValueMatch match) {
    this.internal.match = match;
        return this;
    }
    
    public Builder result(com.grafana.foundation.cog.Builder<ValueMappingResult> result) {
    this.internal.result = result.build();
        return this;
    }
    public DashboardSpecialValueMapOptions build() {
            return this.internal;
        }
    }
}
