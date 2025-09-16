// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class QueryEditorPropertyExpressionOrQueryEditorFunctionExpressionSerializer extends JsonSerializer<QueryEditorPropertyExpressionOrQueryEditorFunctionExpression> {

    @Override
    public void serialize(QueryEditorPropertyExpressionOrQueryEditorFunctionExpression value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.queryEditorPropertyExpression != null) {
            gen.writeObject(value.queryEditorPropertyExpression);
        }
        else if (value.queryEditorFunctionExpression != null) {
            gen.writeObject(value.queryEditorFunctionExpression);
        }
    }
}
