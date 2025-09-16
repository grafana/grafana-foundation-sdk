// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class PanelOrRowPanelSerializer extends JsonSerializer<PanelOrRowPanel> {

    @Override
    public void serialize(PanelOrRowPanel value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.panel != null) {
            gen.writeObject(value.panel);
        }
        else if (value.rowPanel != null) {
            gen.writeObject(value.rowPanel);
        }
    }
}
