// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = QueryEditorPropertyExpressionOrQueryEditorFunctionExpressionDeserializer.class)
public class QueryEditorPropertyExpressionOrQueryEditorFunctionExpression { 
    @JsonUnwrapped
    public QueryEditorPropertyExpression queryEditorPropertyExpression; 
    @JsonUnwrapped
    public QueryEditorFunctionExpression queryEditorFunctionExpression;
    
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
