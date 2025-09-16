// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = QueryEditorPropertyExpressionOrQueryEditorFunctionExpressionDeserializer.class)
public class QueryEditorPropertyExpressionOrQueryEditorFunctionExpression {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorPropertyExpression queryEditorPropertyExpression;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorFunctionExpression queryEditorFunctionExpression;
    protected QueryEditorPropertyExpressionOrQueryEditorFunctionExpression() {}
    public static QueryEditorPropertyExpressionOrQueryEditorFunctionExpression createQueryEditorPropertyExpression(QueryEditorPropertyExpression queryEditorPropertyExpression) {
        QueryEditorPropertyExpressionOrQueryEditorFunctionExpression queryEditorPropertyExpressionOrQueryEditorFunctionExpression = new QueryEditorPropertyExpressionOrQueryEditorFunctionExpression();
        queryEditorPropertyExpressionOrQueryEditorFunctionExpression.queryEditorPropertyExpression = queryEditorPropertyExpression;
        return queryEditorPropertyExpressionOrQueryEditorFunctionExpression;
    }
    public static QueryEditorPropertyExpressionOrQueryEditorFunctionExpression createQueryEditorFunctionExpression(QueryEditorFunctionExpression queryEditorFunctionExpression) {
        QueryEditorPropertyExpressionOrQueryEditorFunctionExpression queryEditorPropertyExpressionOrQueryEditorFunctionExpression = new QueryEditorPropertyExpressionOrQueryEditorFunctionExpression();
        queryEditorPropertyExpressionOrQueryEditorFunctionExpression.queryEditorFunctionExpression = queryEditorFunctionExpression;
        return queryEditorPropertyExpressionOrQueryEditorFunctionExpression;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
