// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class StringOrFloat64Serializer extends JsonSerializer<StringOrFloat64> {

    @Override
    public void serialize(StringOrFloat64 value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.string != null) {
            gen.writeObject(value.string);
        }
        else if (value.float64 != null) {
            gen.writeObject(value.float64);
        }
    }
}
