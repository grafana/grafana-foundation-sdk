// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class InlineScriptSerializer extends JsonSerializer<InlineScript> {

    @Override
    public void serialize(InlineScript value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.string != null) {
            gen.writeObject(value.string);
        }
        else if (value.elasticsearchInlineScript != null) {
            gen.writeObject(value.elasticsearchInlineScript);
        }
    }
}
