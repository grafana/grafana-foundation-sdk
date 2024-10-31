// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class BucketAggregationDeserializer extends JsonDeserializer<BucketAggregation> {

    @Override
    public BucketAggregation deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        BucketAggregation bucketAggregation = new BucketAggregation();
        if (!root.has("type")) {
            throw new IOException("Cannot find discriminator for BucketAggregation");
        }
        String discriminator = root.get("type").asText();  
        
        switch (discriminator) {
        case "date_histogram":
            bucketAggregation.dateHistogram = mapper.convertValue(root, DateHistogram.class);
            break;
        case "filters":
            bucketAggregation.filters = mapper.convertValue(root, Filters.class);
            break;
        case "geohash_grid":
            bucketAggregation.geoHashGrid = mapper.convertValue(root, GeoHashGrid.class);
            break;
        case "histogram":
            bucketAggregation.histogram = mapper.convertValue(root, Histogram.class);
            break;
        case "nested":
            bucketAggregation.nested = mapper.convertValue(root, Nested.class);
            break;
        case "terms":
            bucketAggregation.terms = mapper.convertValue(root, Terms.class);
            break;
        }
        
        return bucketAggregation;
    }
}
