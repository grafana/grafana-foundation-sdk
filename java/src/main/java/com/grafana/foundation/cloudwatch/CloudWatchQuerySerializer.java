// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class CloudWatchQuerySerializer extends JsonSerializer<CloudWatchQuery> {

    @Override
    public void serialize(CloudWatchQuery value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.cloudWatchMetricsQuery != null) {
            gen.writeObject(value.cloudWatchMetricsQuery);
        }
        else if (value.cloudWatchLogsQuery != null) {
            gen.writeObject(value.cloudWatchLogsQuery);
        }
        else if (value.cloudWatchAnnotationQuery != null) {
            gen.writeObject(value.cloudWatchAnnotationQuery);
        }
    }
}
