// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class FiltersSettings { 
    @JsonProperty("filters")
    public List<Filter> filters;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<FiltersSettings> {
        private final FiltersSettings internal;
        
        public Builder() {
            this.internal = new FiltersSettings();
        }
    public Builder filters(com.grafana.foundation.cog.Builder<List<Filter>> filters) {
    this.internal.filters = filters.build();
        return this;
    }
    public FiltersSettings build() {
            return this.internal;
        }
    }
}
