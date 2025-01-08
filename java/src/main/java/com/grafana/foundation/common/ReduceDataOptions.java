// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

// TODO docs
public class ReduceDataOptions {
    // If true show each row value
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("values")
    public Boolean values;
    // if showing all values limit
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("limit")
    public Double limit;
    // When !values, pick one value for the whole field
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("calcs")
    public List<String> calcs;
    // Which fields to show.  By default this is only numeric fields
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fields")
    public String fields;
    public ReduceDataOptions() {
    }
    
    public ReduceDataOptions(Boolean values,Double limit,List<String> calcs,String fields) {
        this.values = values;
        this.limit = limit;
        this.calcs = calcs;
        this.fields = fields;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
