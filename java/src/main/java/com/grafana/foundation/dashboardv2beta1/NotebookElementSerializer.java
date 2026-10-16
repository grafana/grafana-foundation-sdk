// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class NotebookElementSerializer extends JsonSerializer<NotebookElement> {

    @Override
    public void serialize(NotebookElement value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.cellKind != null) {
            gen.writeObject(value.cellKind);
        }
        else if (value.v2PanelKind != null) {
            gen.writeObject(value.v2PanelKind);
        }
        else if (value.libraryPanelKind != null) {
            gen.writeObject(value.libraryPanelKind);
        }
    }
}
