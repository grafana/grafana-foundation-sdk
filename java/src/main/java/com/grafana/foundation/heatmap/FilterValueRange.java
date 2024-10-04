// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Controls the value filter range
public class FilterValueRange {
    // Sets the filter range to values less than or equal to the given value
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("le")
    public Float le;
    // Sets the filter range to values greater than or equal to the given value
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("ge")
    public Float ge;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<FilterValueRange> {
        protected final FilterValueRange internal;
        
        public Builder() {
            this.internal = new FilterValueRange();
        }
    public Builder le(Float le) {
    this.internal.le = le;
        return this;
    }
    
    public Builder ge(Float ge) {
    this.internal.ge = ge;
        return this;
    }
    public FilterValueRange build() {
            return this.internal;
        }
    }
}
