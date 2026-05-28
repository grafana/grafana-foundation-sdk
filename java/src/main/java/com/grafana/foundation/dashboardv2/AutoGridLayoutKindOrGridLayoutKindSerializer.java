// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class AutoGridLayoutKindOrGridLayoutKindSerializer extends JsonSerializer<AutoGridLayoutKindOrGridLayoutKind> {

    @Override
    public void serialize(AutoGridLayoutKindOrGridLayoutKind value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.autoGridLayoutKind != null) {
            gen.writeObject(value.autoGridLayoutKind);
        }
        else if (value.gridLayoutKind != null) {
            gen.writeObject(value.gridLayoutKind);
        }
    }
}
