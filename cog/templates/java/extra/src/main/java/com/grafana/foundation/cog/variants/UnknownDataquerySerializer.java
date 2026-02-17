package com.grafana.foundation.cog.variants;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class UnknownDataquerySerializer extends JsonSerializer<UnknownDataquery> {
    @Override
    public void serialize(UnknownDataquery unknownDataquery, JsonGenerator jsonGenerator, SerializerProvider serializerProvider) throws IOException {
        jsonGenerator.writeObject(unknownDataquery.genericFields);
    }
}
