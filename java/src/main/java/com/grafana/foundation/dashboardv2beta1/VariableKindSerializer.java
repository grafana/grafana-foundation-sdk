// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class VariableKindSerializer extends JsonSerializer<VariableKind> {

    @Override
    public void serialize(VariableKind value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.queryVariableKind != null) {
            gen.writeObject(value.queryVariableKind);
        }
        else if (value.textVariableKind != null) {
            gen.writeObject(value.textVariableKind);
        }
        else if (value.constantVariableKind != null) {
            gen.writeObject(value.constantVariableKind);
        }
        else if (value.datasourceVariableKind != null) {
            gen.writeObject(value.datasourceVariableKind);
        }
        else if (value.intervalVariableKind != null) {
            gen.writeObject(value.intervalVariableKind);
        }
        else if (value.customVariableKind != null) {
            gen.writeObject(value.customVariableKind);
        }
        else if (value.groupByVariableKind != null) {
            gen.writeObject(value.groupByVariableKind);
        }
        else if (value.adhocVariableKind != null) {
            gen.writeObject(value.adhocVariableKind);
        }
        else if (value.switchVariableKind != null) {
            gen.writeObject(value.switchVariableKind);
        }
    }
}
