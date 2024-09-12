// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryEditorFunctionParameterExpression {
    @JsonProperty("type")
    public String type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public String name;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<QueryEditorFunctionParameterExpression> {
        private final QueryEditorFunctionParameterExpression internal;
        
        public Builder() {
            this.internal = new QueryEditorFunctionParameterExpression();
    this.internal.type = "functionParameter";
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    public QueryEditorFunctionParameterExpression build() {
            return this.internal;
        }
    }
}
