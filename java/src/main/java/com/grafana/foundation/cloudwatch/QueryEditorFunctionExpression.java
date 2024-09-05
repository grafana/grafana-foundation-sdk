// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import java.util.List;

public class QueryEditorFunctionExpression {
    @JsonProperty("type")
    public String type;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("name")
    public String name;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("parameters")
    public List<QueryEditorFunctionParameterExpression> parameters;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> {
        private final QueryEditorFunctionExpression internal;
        
        public Builder() {
            this.internal = new QueryEditorFunctionExpression();
    this.internal.type = "function";
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder parameters(com.grafana.foundation.cog.Builder<List<QueryEditorFunctionParameterExpression>> parameters) {
    this.internal.parameters = parameters.build();
        return this;
    }
    public QueryEditorFunctionExpression build() {
            return this.internal;
        }
    }
}
