// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class ValueMappingSerializer extends JsonSerializer<ValueMapping> {

    @Override
    public void serialize(ValueMapping value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.valueMap != null) {
            gen.writeObject(value.valueMap);
        }
        else if (value.rangeMap != null) {
            gen.writeObject(value.rangeMap);
        }
        else if (value.regexMap != null) {
            gen.writeObject(value.regexMap);
        }
        else if (value.specialValueMap != null) {
            gen.writeObject(value.specialValueMap);
        }
    }
}
