// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class PipelineMetricAggregationDeserializer extends JsonDeserializer<PipelineMetricAggregation> {

    @Override
    public PipelineMetricAggregation deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        PipelineMetricAggregation pipelineMetricAggregation = new PipelineMetricAggregation();
        if (!root.has("type")) {
            throw new IOException("Cannot find discriminator for PipelineMetricAggregation");
        }
        String discriminator = root.get("type").asText();  
        
        switch (discriminator) {
        case "bucket_script":
            pipelineMetricAggregation.bucketScript = mapper.convertValue(root, BucketScript.class);
            break;
        case "cumulative_sum":
            pipelineMetricAggregation.cumulativeSum = mapper.convertValue(root, CumulativeSum.class);
            break;
        case "derivative":
            pipelineMetricAggregation.derivative = mapper.convertValue(root, Derivative.class);
            break;
        case "moving_avg":
            pipelineMetricAggregation.movingAverage = mapper.convertValue(root, MovingAverage.class);
            break;
        }
        
        return pipelineMetricAggregation;
    }
}
