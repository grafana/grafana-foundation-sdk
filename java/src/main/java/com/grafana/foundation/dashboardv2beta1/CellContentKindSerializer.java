// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class CellContentKindSerializer extends JsonSerializer<CellContentKind> {

    @Override
    public void serialize(CellContentKind value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.markdownCellContentKind != null) {
            gen.writeObject(value.markdownCellContentKind);
        }
        else if (value.codeCellContentKind != null) {
            gen.writeObject(value.codeCellContentKind);
        }
    }
}
