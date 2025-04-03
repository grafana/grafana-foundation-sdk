// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class QueryEditorOperatorTypeSerializer extends JsonSerializer<QueryEditorOperatorType> {

    @Override
    public void serialize(QueryEditorOperatorType value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
         if (value.string != null) {
            gen.writeObject(value.string);
        }
         else  if (value.bool != null) {
            gen.writeObject(value.bool);
        }
         else  if (value.int64 != null) {
            gen.writeObject(value.int64);
        }
    }
}
