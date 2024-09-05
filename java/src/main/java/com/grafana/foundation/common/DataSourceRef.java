// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class DataSourceRef {
    // The plugin type-id
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("type")
    public String type;
    // Specific datasource instance
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("uid")
    public String uid;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DataSourceRef> {
        private final DataSourceRef internal;
        
        public Builder() {
            this.internal = new DataSourceRef();
        }
    public Builder type(String type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder uid(String uid) {
    this.internal.uid = uid;
        return this;
    }
    public DataSourceRef build() {
            return this.internal;
        }
    }
}
