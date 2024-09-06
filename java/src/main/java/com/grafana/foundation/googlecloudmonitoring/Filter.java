// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Query filter representation.
public class Filter {
    // Filter key.
    @JsonProperty("key")
    public String key;
    // Filter operator.
    @JsonProperty("operator")
    public String operator;
    // Filter value.
    @JsonProperty("value")
    public String value;
    // Filter condition.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("condition")
    public String condition;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Filter> {
        private final Filter internal;
        
        public Builder() {
            this.internal = new Filter();
        }
    public Builder key(String key) {
    this.internal.key = key;
        return this;
    }
    
    public Builder operator(String operator) {
    this.internal.operator = operator;
        return this;
    }
    
    public Builder value(String value) {
    this.internal.value = value;
        return this;
    }
    
    public Builder condition(String condition) {
    this.internal.condition = condition;
        return this;
    }
    public Filter build() {
            return this.internal;
        }
    }
}
