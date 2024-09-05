// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

@JsonDeserialize(using = QueryEditorExpressionDeserializer.class)
public class QueryEditorExpression {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorArrayExpression queryEditorArrayExpression;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorPropertyExpression queryEditorPropertyExpression;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorGroupByExpression queryEditorGroupByExpression;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorFunctionExpression queryEditorFunctionExpression;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorFunctionParameterExpression queryEditorFunctionParameterExpression;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected QueryEditorOperatorExpression queryEditorOperatorExpression;
    protected QueryEditorExpression() {}
    public static QueryEditorExpression createQueryEditorArrayExpression(com.grafana.foundation.cog.Builder<QueryEditorArrayExpression> queryEditorArrayExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorArrayExpression = queryEditorArrayExpression.build();
        return queryEditorExpression;
    }
    public static QueryEditorExpression createQueryEditorPropertyExpression(com.grafana.foundation.cog.Builder<QueryEditorPropertyExpression> queryEditorPropertyExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorPropertyExpression = queryEditorPropertyExpression.build();
        return queryEditorExpression;
    }
    public static QueryEditorExpression createQueryEditorGroupByExpression(com.grafana.foundation.cog.Builder<QueryEditorGroupByExpression> queryEditorGroupByExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorGroupByExpression = queryEditorGroupByExpression.build();
        return queryEditorExpression;
    }
    public static QueryEditorExpression createQueryEditorFunctionExpression(com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> queryEditorFunctionExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorFunctionExpression = queryEditorFunctionExpression.build();
        return queryEditorExpression;
    }
    public static QueryEditorExpression createQueryEditorFunctionParameterExpression(com.grafana.foundation.cog.Builder<QueryEditorFunctionParameterExpression> queryEditorFunctionParameterExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorFunctionParameterExpression = queryEditorFunctionParameterExpression.build();
        return queryEditorExpression;
    }
    public static QueryEditorExpression createQueryEditorOperatorExpression(com.grafana.foundation.cog.Builder<QueryEditorOperatorExpression> queryEditorOperatorExpression) {
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        queryEditorExpression.queryEditorOperatorExpression = queryEditorOperatorExpression.build();
        return queryEditorExpression;
    }
    
    public String toJSON() throws JsonProcessingException {
        if (queryEditorArrayExpression != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(queryEditorArrayExpression);
        }
        if (queryEditorPropertyExpression != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(queryEditorPropertyExpression);
        }
        if (queryEditorGroupByExpression != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(queryEditorGroupByExpression);
        }
        if (queryEditorFunctionExpression != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(queryEditorFunctionExpression);
        }
        if (queryEditorFunctionParameterExpression != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(queryEditorFunctionParameterExpression);
        }
        if (queryEditorOperatorExpression != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(queryEditorOperatorExpression);
        }
        
        return null;
    }

}
