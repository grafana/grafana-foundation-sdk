// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class BuilderQueryEditorOperatorTypeSerializer extends JsonSerializer<BuilderQueryEditorOperatorType> {

    @Override
    public void serialize(BuilderQueryEditorOperatorType value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.string != null) {
            gen.writeObject(value.string);
        }
        else if (value.bool != null) {
            gen.writeObject(value.bool);
        }
        else if (value.float64 != null) {
            gen.writeObject(value.float64);
        }
        else if (value.selectableValue != null) {
            gen.writeObject(value.selectableValue);
        }
    }
}
