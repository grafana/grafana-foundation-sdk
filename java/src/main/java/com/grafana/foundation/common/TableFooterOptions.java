// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonInclude;

// Footer options
public class TableFooterOptions {
    @JsonProperty("show")
    public Boolean show;
    // actually 1 value
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("reducer")
    public List<String> reducer;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fields")
    public List<String> fields;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("enablePagination")
    public Boolean enablePagination;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("countRows")
    public Boolean countRows;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TableFooterOptions> {
        protected final TableFooterOptions internal;
        
        public Builder() {
            this.internal = new TableFooterOptions();
        }
    public Builder show(Boolean show) {
    this.internal.show = show;
        return this;
    }
    
    public Builder reducer(List<String> reducer) {
    this.internal.reducer = reducer;
        return this;
    }
    
    public Builder fields(List<String> fields) {
    this.internal.fields = fields;
        return this;
    }
    
    public Builder enablePagination(Boolean enablePagination) {
    this.internal.enablePagination = enablePagination;
        return this;
    }
    
    public Builder countRows(Boolean countRows) {
    this.internal.countRows = countRows;
        return this;
    }
    public TableFooterOptions build() {
            return this.internal;
        }
    }
}
