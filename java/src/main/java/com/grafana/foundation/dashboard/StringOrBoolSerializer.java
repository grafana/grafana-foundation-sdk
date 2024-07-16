// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class StringOrBoolSerializer extends JsonSerializer<StringOrBool> {

    @Override
    public void serialize(StringOrBool value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
         if (value.string != null) {
            gen.writeObject(value.string);
        }
         else  if (value.bool != null) {
            gen.writeObject(value.bool);
        }
    }
}
