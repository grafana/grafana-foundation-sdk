// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class StringOrInt64Serializer extends JsonSerializer<StringOrInt64> {

    @Override
    public void serialize(StringOrInt64 value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.string != null) {
            gen.writeObject(value.string);
        }
        else if (value.int64 != null) {
            gen.writeObject(value.int64);
        }
    }
}
