// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class DashboardRegexMapOptions {
    // Regular expression to match against
    @JsonProperty("pattern")
    public String pattern;
    // Config to apply when the value matches the regex
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("result")
    public ValueMappingResult result;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DashboardRegexMapOptions> {
        protected final DashboardRegexMapOptions internal;
        
        public Builder() {
            this.internal = new DashboardRegexMapOptions();
        }
    public Builder pattern(String pattern) {
    this.internal.pattern = pattern;
        return this;
    }
    
    public Builder result(com.grafana.foundation.cog.Builder<ValueMappingResult> result) {
    this.internal.result = result.build();
        return this;
    }
    public DashboardRegexMapOptions build() {
            return this.internal;
        }
    }
}
