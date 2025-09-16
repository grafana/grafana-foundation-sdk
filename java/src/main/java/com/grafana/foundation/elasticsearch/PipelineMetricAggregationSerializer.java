// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class PipelineMetricAggregationSerializer extends JsonSerializer<PipelineMetricAggregation> {

    @Override
    public void serialize(PipelineMetricAggregation value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.movingAverage != null) {
            gen.writeObject(value.movingAverage);
        }
        else if (value.derivative != null) {
            gen.writeObject(value.derivative);
        }
        else if (value.cumulativeSum != null) {
            gen.writeObject(value.cumulativeSum);
        }
        else if (value.bucketScript != null) {
            gen.writeObject(value.bucketScript);
        }
    }
}
