// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class QueryEditorFunctionExpression {
    @JsonProperty("type")
    public QueryEditorExpressionType type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("parameters")
    public List<QueryEditorFunctionParameterExpression> parameters;
    public QueryEditorFunctionExpression() {
        this.type = QueryEditorExpressionType.FUNCTION;
    }
    public QueryEditorFunctionExpression(String name,List<QueryEditorFunctionParameterExpression> parameters) {
        this.type = QueryEditorExpressionType.FUNCTION;
        this.name = name;
        this.parameters = parameters;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
