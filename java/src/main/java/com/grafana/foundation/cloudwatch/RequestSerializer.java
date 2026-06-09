// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class RequestSerializer extends JsonSerializer<Request> {

    @Override
    public void serialize(Request value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.metricsQuery != null) {
            gen.writeObject(value.metricsQuery);
        }
        else if (value.logsQuery != null) {
            gen.writeObject(value.logsQuery);
        }
        else if (value.annotationQuery != null) {
            gen.writeObject(value.annotationQuery);
        }
    }
}
