// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class ElementSerializer extends JsonSerializer<Element> {

    @Override
    public void serialize(Element value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.panelKind != null) {
            gen.writeObject(value.panelKind);
        }
        else if (value.libraryPanelKind != null) {
            gen.writeObject(value.libraryPanelKind);
        }
    }
}
