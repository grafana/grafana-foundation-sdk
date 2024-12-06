// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Controls cell value options
public class CellValues {
    // Controls the cell value unit
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("unit")
    public String unit;
    // Controls the number of decimals for cell values
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("decimals")
    public Float decimals;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CellValues> {
        protected final CellValues internal;
        
        public Builder() {
            this.internal = new CellValues();
        }
    public Builder unit(String unit) {
    this.internal.unit = unit;
        return this;
    }
    
    public Builder decimals(Float decimals) {
    this.internal.decimals = decimals;
        return this;
    }
    public CellValues build() {
            return this.internal;
        }
    }
}
