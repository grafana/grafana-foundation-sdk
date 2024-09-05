// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import java.util.List;

public class AzureMetricDimension {
    // Name of Dimension to be filtered on.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("dimension")
    public String dimension;
    // String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("operator")
    public String operator;
    // Values to match with the filter.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("filters")
    public List<String> filters;
    // @deprecated filter is deprecated in favour of filters to support multiselect.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("filter")
    public String filter;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AzureMetricDimension> {
        private final AzureMetricDimension internal;
        
        public Builder() {
            this.internal = new AzureMetricDimension();
        }
    public Builder dimension(String dimension) {
    this.internal.dimension = dimension;
        return this;
    }
    
    public Builder operator(String operator) {
    this.internal.operator = operator;
        return this;
    }
    
    public Builder filters(List<String> filters) {
    this.internal.filters = filters;
        return this;
    }
    
    public Builder filter(String filter) {
    this.internal.filter = filter;
        return this;
    }
    public AzureMetricDimension build() {
            return this.internal;
        }
    }
}
