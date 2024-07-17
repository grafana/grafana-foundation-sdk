// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Filter { 
    @JsonProperty("query")
    public String query; 
    @JsonProperty("label")
    public String label;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Filter> {
        private final Filter internal;
        
        public Builder() {
            this.internal = new Filter();
        }
    public Builder query(String query) {
    this.internal.query = query;
        return this;
    }
    
    public Builder label(String label) {
    this.internal.label = label;
        return this;
    }
    public Filter build() {
            return this.internal;
        }
    }
}
