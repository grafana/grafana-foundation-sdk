// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class BoolOrFloat64Serializer extends JsonSerializer<BoolOrFloat64> {

    @Override
    public void serialize(BoolOrFloat64 value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.bool != null) {
            gen.writeObject(value.bool);
        }
        else if (value.float64 != null) {
            gen.writeObject(value.float64);
        }
    }
}
