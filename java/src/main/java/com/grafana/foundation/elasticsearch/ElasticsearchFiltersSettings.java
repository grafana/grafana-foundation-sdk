// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class ElasticsearchFiltersSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filters")
    public List<Filter> filters;
    public ElasticsearchFiltersSettings() {
    }
    public ElasticsearchFiltersSettings(List<Filter> filters) {
        this.filters = filters;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
