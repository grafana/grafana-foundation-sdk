// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class StringOrMapSerializer extends JsonSerializer<StringOrMap> {

    @Override
    public void serialize(StringOrMap value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.string != null) {
            gen.writeObject(value.string);
        }
        else if (value.map != null) {
            gen.writeObject(value.map);
        }
    }
}
