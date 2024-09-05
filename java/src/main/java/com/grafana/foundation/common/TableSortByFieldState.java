// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Sort by field state
public class TableSortByFieldState {
    // Sets the display name of the field to sort by
    @JsonProperty("displayName")
    public String displayName;
    // Flag used to indicate descending sort order
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("desc")
    public Boolean desc;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TableSortByFieldState> {
        private final TableSortByFieldState internal;
        
        public Builder() {
            this.internal = new TableSortByFieldState();
        }
    public Builder displayName(String displayName) {
    this.internal.displayName = displayName;
        return this;
    }
    
    public Builder desc(Boolean desc) {
    this.internal.desc = desc;
        return this;
    }
    public TableSortByFieldState build() {
            return this.internal;
        }
    }
}
