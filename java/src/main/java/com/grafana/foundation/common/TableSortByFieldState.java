// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Sort by field state
public class TableSortByFieldState {
    // Sets the display name of the field to sort by
    @JsonProperty("displayName")
    public String displayName;
    // Flag used to indicate descending sort order
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("desc")
    public Boolean desc;
    public TableSortByFieldState() {
    }
    public TableSortByFieldState(String displayName,Boolean desc) {
        this.displayName = displayName;
        this.desc = desc;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
