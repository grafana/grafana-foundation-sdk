// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class VariableValueSingleSerializer extends JsonSerializer<VariableValueSingle> {

    @Override
    public void serialize(VariableValueSingle value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.string != null) {
            gen.writeObject(value.string);
        }
        else if (value.bool != null) {
            gen.writeObject(value.bool);
        }
        else if (value.float64 != null) {
            gen.writeObject(value.float64);
        }
        else if (value.customVariableValue != null) {
            gen.writeObject(value.customVariableValue);
        }
    }
}
