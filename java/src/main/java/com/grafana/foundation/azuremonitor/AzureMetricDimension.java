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
    public AzureMetricDimension() {
    }
    public AzureMetricDimension(String dimension,String operator,List<String> filters,String filter) {
        this.dimension = dimension;
        this.operator = operator;
        this.filters = filters;
        this.filter = filter;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
