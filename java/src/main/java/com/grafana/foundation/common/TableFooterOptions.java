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
    public TableFooterOptions() {
    }
    public TableFooterOptions(Boolean show,List<String> reducer,List<String> fields,Boolean enablePagination,Boolean countRows) {
        this.show = show;
        this.reducer = reducer;
        this.fields = fields;
        this.enablePagination = enablePagination;
        this.countRows = countRows;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
