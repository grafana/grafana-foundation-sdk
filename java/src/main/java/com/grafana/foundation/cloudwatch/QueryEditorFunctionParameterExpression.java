// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryEditorFunctionParameterExpression {
    @JsonProperty("type")
    public QueryEditorExpressionType type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public String name;
    public QueryEditorFunctionParameterExpression() {
        this.type = QueryEditorExpressionType.FUNCTION_PARAMETER;
    }
    public QueryEditorFunctionParameterExpression(String name) {
        this.type = QueryEditorExpressionType.FUNCTION_PARAMETER;
        this.name = name;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
