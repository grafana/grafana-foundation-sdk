// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class BoolOrUint32Serializer extends JsonSerializer<BoolOrUint32> {

    @Override
    public void serialize(BoolOrUint32 value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.bool != null) {
            gen.writeObject(value.bool);
        }
        else if (value.uint32 != null) {
            gen.writeObject(value.uint32);
        }
    }
}
