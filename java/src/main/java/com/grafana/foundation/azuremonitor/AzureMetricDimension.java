// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class AzureMetricDimension {
    // Name of Dimension to be filtered on.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dimension")
    public String dimension;
    // String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("operator")
    public String operator;
    // Values to match with the filter.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filters")
    public List<String> filters;
    // @deprecated filter is deprecated in favour of filters to support multiselect.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filter")
    public String filter;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AzureMetricDimension> {
        protected final AzureMetricDimension internal;
        
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
