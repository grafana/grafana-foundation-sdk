// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class AzureTracesFilter {
    // Property name, auto-populated based on available traces.
    @JsonProperty("property")
    public String property;
    // Comparison operator to use. Either equals or not equals.
    @JsonProperty("operation")
    public String operation;
    // Values to filter by.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("filters")
    public List<String> filters;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AzureTracesFilter> {
        private final AzureTracesFilter internal;
        
        public Builder() {
            this.internal = new AzureTracesFilter();
        }
    public Builder property(String property) {
    this.internal.property = property;
        return this;
    }
    
    public Builder operation(String operation) {
    this.internal.operation = operation;
        return this;
    }
    
    public Builder filters(List<String> filters) {
    this.internal.filters = filters;
        return this;
    }
    public AzureTracesFilter build() {
            return this.internal;
        }
    }
}
