// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class DashboardFieldConfigSourceOverrides {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("matcher")
    public MatcherConfig matcher;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("properties")
    public List<DynamicConfigValue> properties;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides> {
        protected final DashboardFieldConfigSourceOverrides internal;
        
        public Builder() {
            this.internal = new DashboardFieldConfigSourceOverrides();
        }
    public Builder matcher(MatcherConfig matcher) {
    this.internal.matcher = matcher;
        return this;
    }
    
    public Builder properties(List<DynamicConfigValue> properties) {
    this.internal.properties = properties;
        return this;
    }
    public DashboardFieldConfigSourceOverrides build() {
            return this.internal;
        }
    }
}
