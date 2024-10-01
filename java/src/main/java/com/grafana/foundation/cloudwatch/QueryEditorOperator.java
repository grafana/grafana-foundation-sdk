// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
public class QueryEditorOperator {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("value")
    public QueryEditorOperatorValueType value;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<QueryEditorOperator> {
        protected final QueryEditorOperator internal;
        
        public Builder() {
            this.internal = new QueryEditorOperator();
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder value(QueryEditorOperatorValueType value) {
    this.internal.value = value;
        return this;
    }
    public QueryEditorOperator build() {
            return this.internal;
        }
    }
}
