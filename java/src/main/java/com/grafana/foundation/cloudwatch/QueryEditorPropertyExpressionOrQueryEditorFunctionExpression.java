// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

@JsonDeserialize(using = QueryEditorPropertyExpressionOrQueryEditorFunctionExpressionDeserializer.class)
public class QueryEditorPropertyExpressionOrQueryEditorFunctionExpression {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorPropertyExpression queryEditorPropertyExpression;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorFunctionExpression queryEditorFunctionExpression;
    protected QueryEditorPropertyExpressionOrQueryEditorFunctionExpression() {}
    public static QueryEditorPropertyExpressionOrQueryEditorFunctionExpression createQueryEditorPropertyExpression(com.grafana.foundation.cog.Builder<QueryEditorPropertyExpression> queryEditorPropertyExpression) {
        QueryEditorPropertyExpressionOrQueryEditorFunctionExpression queryEditorPropertyExpressionOrQueryEditorFunctionExpression = new QueryEditorPropertyExpressionOrQueryEditorFunctionExpression();
        queryEditorPropertyExpressionOrQueryEditorFunctionExpression.queryEditorPropertyExpression = queryEditorPropertyExpression.build();
        return queryEditorPropertyExpressionOrQueryEditorFunctionExpression;
    }
    public static QueryEditorPropertyExpressionOrQueryEditorFunctionExpression createQueryEditorFunctionExpression(com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> queryEditorFunctionExpression) {
        QueryEditorPropertyExpressionOrQueryEditorFunctionExpression queryEditorPropertyExpressionOrQueryEditorFunctionExpression = new QueryEditorPropertyExpressionOrQueryEditorFunctionExpression();
        queryEditorPropertyExpressionOrQueryEditorFunctionExpression.queryEditorFunctionExpression = queryEditorFunctionExpression.build();
        return queryEditorPropertyExpressionOrQueryEditorFunctionExpression;
    }
    
    public String toJSON() throws JsonProcessingException {
        if (queryEditorPropertyExpression != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(queryEditorPropertyExpression);
        }
        if (queryEditorFunctionExpression != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(queryEditorFunctionExpression);
        }
        
        return null;
    }

}
