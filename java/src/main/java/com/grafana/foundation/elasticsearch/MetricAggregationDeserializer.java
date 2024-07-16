// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class MetricAggregationDeserializer extends JsonDeserializer<MetricAggregation> {

    @Override
    public MetricAggregation deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        MetricAggregation metricAggregation = new MetricAggregation();
        if (!root.has("type")) {
            throw new IOException("Cannot find discriminator for MetricAggregation");
        }
        String discriminator = root.get("type").asText();  
        
        switch (discriminator) {
        case "avg":
            metricAggregation.average = mapper.convertValue(root, Average.class);
            break;
        case "bucket_script":
            metricAggregation.bucketScript = mapper.convertValue(root, BucketScript.class);
            break;
        case "cardinality":
            metricAggregation.uniqueCount = mapper.convertValue(root, UniqueCount.class);
            break;
        case "count":
            metricAggregation.count = mapper.convertValue(root, Count.class);
            break;
        case "cumulative_sum":
            metricAggregation.cumulativeSum = mapper.convertValue(root, CumulativeSum.class);
            break;
        case "derivative":
            metricAggregation.derivative = mapper.convertValue(root, Derivative.class);
            break;
        case "extended_stats":
            metricAggregation.extendedStats = mapper.convertValue(root, ExtendedStats.class);
            break;
        case "logs":
            metricAggregation.logs = mapper.convertValue(root, Logs.class);
            break;
        case "max":
            metricAggregation.max = mapper.convertValue(root, Max.class);
            break;
        case "min":
            metricAggregation.min = mapper.convertValue(root, Min.class);
            break;
        case "moving_avg":
            metricAggregation.movingAverage = mapper.convertValue(root, MovingAverage.class);
            break;
        case "moving_fn":
            metricAggregation.movingFunction = mapper.convertValue(root, MovingFunction.class);
            break;
        case "percentiles":
            metricAggregation.percentiles = mapper.convertValue(root, Percentiles.class);
            break;
        case "rate":
            metricAggregation.rate = mapper.convertValue(root, Rate.class);
            break;
        case "raw_data":
            metricAggregation.rawData = mapper.convertValue(root, RawData.class);
            break;
        case "raw_document":
            metricAggregation.rawDocument = mapper.convertValue(root, RawDocument.class);
            break;
        case "serial_diff":
            metricAggregation.serialDiff = mapper.convertValue(root, SerialDiff.class);
            break;
        case "sum":
            metricAggregation.sum = mapper.convertValue(root, Sum.class);
            break;
        case "top_metrics":
            metricAggregation.topMetrics = mapper.convertValue(root, TopMetrics.class);
            break;
        }
        
        return metricAggregation;
    }
}
