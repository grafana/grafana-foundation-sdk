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
    public AzureTracesFilter() {
    }
    public AzureTracesFilter(String property,String operation,List<String> filters) {
        this.property = property;
        this.operation = operation;
        this.filters = filters;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
