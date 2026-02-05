// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = QueryEditorExpressionDeserializer.class)
public class QueryEditorExpression {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorArrayExpression queryEditorArrayExpression;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorPropertyExpression queryEditorPropertyExpression;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorGroupByExpression queryEditorGroupByExpression;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorFunctionExpression queryEditorFunctionExpression;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorFunctionParameterExpression queryEditorFunctionParameterExpression;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorOperatorExpression queryEditorOperatorExpression;
    protected QueryEditorExpression() {}
    public static QueryEditorExpression createQueryEditorArrayExpression(QueryEditorArrayExpression queryEditorArrayExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorArrayExpression = queryEditorArrayExpression;
        return queryEditorExpression;
    }
    public static QueryEditorExpression createQueryEditorPropertyExpression(QueryEditorPropertyExpression queryEditorPropertyExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorPropertyExpression = queryEditorPropertyExpression;
        return queryEditorExpression;
    }
    public static QueryEditorExpression createQueryEditorGroupByExpression(QueryEditorGroupByExpression queryEditorGroupByExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorGroupByExpression = queryEditorGroupByExpression;
        return queryEditorExpression;
    }
    public static QueryEditorExpression createQueryEditorFunctionExpression(QueryEditorFunctionExpression queryEditorFunctionExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorFunctionExpression = queryEditorFunctionExpression;
        return queryEditorExpression;
    }
    public static QueryEditorExpression createQueryEditorFunctionParameterExpression(QueryEditorFunctionParameterExpression queryEditorFunctionParameterExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorFunctionParameterExpression = queryEditorFunctionParameterExpression;
        return queryEditorExpression;
    }
    public static QueryEditorExpression createQueryEditorOperatorExpression(QueryEditorOperatorExpression queryEditorOperatorExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorOperatorExpression = queryEditorOperatorExpression;
        return queryEditorExpression;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
