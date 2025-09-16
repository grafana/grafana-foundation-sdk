// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class BucketAggregationSerializer extends JsonSerializer<BucketAggregation> {

    @Override
    public void serialize(BucketAggregation value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.dateHistogram != null) {
            gen.writeObject(value.dateHistogram);
        }
        else if (value.histogram != null) {
            gen.writeObject(value.histogram);
        }
        else if (value.terms != null) {
            gen.writeObject(value.terms);
        }
        else if (value.filters != null) {
            gen.writeObject(value.filters);
        }
        else if (value.geoHashGrid != null) {
            gen.writeObject(value.geoHashGrid);
        }
        else if (value.nested != null) {
            gen.writeObject(value.nested);
        }
    }
}
