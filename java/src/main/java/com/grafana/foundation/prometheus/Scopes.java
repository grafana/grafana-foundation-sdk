// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class Scopes {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("defaultPath")
    public List<String> defaultPath;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filters")
    public List<ScopesFilters> filters;
    @JsonProperty("title")
    public String title;
    public Scopes() {
        this.title = "";
    }
    public Scopes(List<String> defaultPath,List<ScopesFilters> filters,String title) {
        this.defaultPath = defaultPath;
        this.filters = filters;
        this.title = title;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
