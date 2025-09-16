// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class QueryEditorExpressionSerializer extends JsonSerializer<QueryEditorExpression> {

    @Override
    public void serialize(QueryEditorExpression value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.queryEditorArrayExpression != null) {
            gen.writeObject(value.queryEditorArrayExpression);
        }
        else if (value.queryEditorPropertyExpression != null) {
            gen.writeObject(value.queryEditorPropertyExpression);
        }
        else if (value.queryEditorGroupByExpression != null) {
            gen.writeObject(value.queryEditorGroupByExpression);
        }
        else if (value.queryEditorFunctionExpression != null) {
            gen.writeObject(value.queryEditorFunctionExpression);
        }
        else if (value.queryEditorFunctionParameterExpression != null) {
            gen.writeObject(value.queryEditorFunctionParameterExpression);
        }
        else if (value.queryEditorOperatorExpression != null) {
            gen.writeObject(value.queryEditorOperatorExpression);
        }
    }
}
