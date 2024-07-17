// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// TODO docs
public class ReduceDataOptions {
    // If true show each row value 
    @JsonProperty("values")
    public Boolean values;
    // if showing all values limit 
    @JsonProperty("limit")
    public Double limit;
    // When !values, pick one value for the whole field 
    @JsonProperty("calcs")
    public List<String> calcs;
    // Which fields to show.  By default this is only numeric fields 
    @JsonProperty("fields")
    public String fields;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ReduceDataOptions> {
        private final ReduceDataOptions internal;
        
        public Builder() {
            this.internal = new ReduceDataOptions();
        }
    public Builder values(Boolean values) {
    this.internal.values = values;
        return this;
    }
    
    public Builder limit(Double limit) {
    this.internal.limit = limit;
        return this;
    }
    
    public Builder calcs(List<String> calcs) {
    this.internal.calcs = calcs;
        return this;
    }
    
    public Builder fields(String fields) {
    this.internal.fields = fields;
        return this;
    }
    public ReduceDataOptions build() {
            return this.internal;
        }
    }
}
