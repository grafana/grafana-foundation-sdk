// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryEditorProperty {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public QueryEditorPropertyType type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public String name;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<QueryEditorProperty> {
        protected final QueryEditorProperty internal;
        
        public Builder() {
            this.internal = new QueryEditorProperty();
        }
    public Builder type(QueryEditorPropertyType type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    public QueryEditorProperty build() {
            return this.internal;
        }
    }
}
