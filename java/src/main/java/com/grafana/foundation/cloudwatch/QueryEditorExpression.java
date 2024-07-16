// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = QueryEditorExpressionDeserializer.class)
public class QueryEditorExpression { 
    @JsonUnwrapped
    public QueryEditorArrayExpression queryEditorArrayExpression; 
    @JsonUnwrapped
    public QueryEditorPropertyExpression queryEditorPropertyExpression; 
    @JsonUnwrapped
    public QueryEditorGroupByExpression queryEditorGroupByExpression; 
    @JsonUnwrapped
    public QueryEditorFunctionExpression queryEditorFunctionExpression; 
    @JsonUnwrapped
    public QueryEditorFunctionParameterExpression queryEditorFunctionParameterExpression; 
    @JsonUnwrapped
    public QueryEditorOperatorExpression queryEditorOperatorExpression;
    
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
