// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class MetricAggregationSerializer extends JsonSerializer<MetricAggregation> {

    @Override
    public void serialize(MetricAggregation value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.count != null) {
            gen.writeObject(value.count);
        }
        else if (value.movingAverage != null) {
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
        else if (value.serialDiff != null) {
            gen.writeObject(value.serialDiff);
        }
        else if (value.rawData != null) {
            gen.writeObject(value.rawData);
        }
        else if (value.rawDocument != null) {
            gen.writeObject(value.rawDocument);
        }
        else if (value.uniqueCount != null) {
            gen.writeObject(value.uniqueCount);
        }
        else if (value.percentiles != null) {
            gen.writeObject(value.percentiles);
        }
        else if (value.extendedStats != null) {
            gen.writeObject(value.extendedStats);
        }
        else if (value.min != null) {
            gen.writeObject(value.min);
        }
        else if (value.max != null) {
            gen.writeObject(value.max);
        }
        else if (value.sum != null) {
            gen.writeObject(value.sum);
        }
        else if (value.average != null) {
            gen.writeObject(value.average);
        }
        else if (value.movingFunction != null) {
            gen.writeObject(value.movingFunction);
        }
        else if (value.logs != null) {
            gen.writeObject(value.logs);
        }
        else if (value.rate != null) {
            gen.writeObject(value.rate);
        }
        else if (value.topMetrics != null) {
            gen.writeObject(value.topMetrics);
        }
    }
}
