// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class BuiltinRoleRefOrCustomRoleRefSerializer extends JsonSerializer<BuiltinRoleRefOrCustomRoleRef> {

    @Override
    public void serialize(BuiltinRoleRefOrCustomRoleRef value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.builtinRoleRef != null) {
            gen.writeObject(value.builtinRoleRef);
        }
        else if (value.customRoleRef != null) {
            gen.writeObject(value.customRoleRef);
        }
    }
}
