// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class MetricAggregationWithSettingsDeserializer extends JsonDeserializer<MetricAggregationWithSettings> {

    @Override
    public MetricAggregationWithSettings deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        if (!root.has("type")) {
            throw new IOException("Cannot find discriminator for MetricAggregationWithSettings");
        }
        String discriminator = root.get("type").asText();  
        
        switch (discriminator) {
        case "avg":
            metricAggregationWithSettings.average = mapper.convertValue(root, Average.class);
            break;
        case "bucket_script":
            metricAggregationWithSettings.bucketScript = mapper.convertValue(root, BucketScript.class);
            break;
        case "cardinality":
            metricAggregationWithSettings.uniqueCount = mapper.convertValue(root, UniqueCount.class);
            break;
        case "cumulative_sum":
            metricAggregationWithSettings.cumulativeSum = mapper.convertValue(root, CumulativeSum.class);
            break;
        case "derivative":
            metricAggregationWithSettings.derivative = mapper.convertValue(root, Derivative.class);
            break;
        case "extended_stats":
            metricAggregationWithSettings.extendedStats = mapper.convertValue(root, ExtendedStats.class);
            break;
        case "logs":
            metricAggregationWithSettings.logs = mapper.convertValue(root, Logs.class);
            break;
        case "max":
            metricAggregationWithSettings.max = mapper.convertValue(root, Max.class);
            break;
        case "min":
            metricAggregationWithSettings.min = mapper.convertValue(root, Min.class);
            break;
        case "moving_avg":
            metricAggregationWithSettings.movingAverage = mapper.convertValue(root, MovingAverage.class);
            break;
        case "moving_fn":
            metricAggregationWithSettings.movingFunction = mapper.convertValue(root, MovingFunction.class);
            break;
        case "percentiles":
            metricAggregationWithSettings.percentiles = mapper.convertValue(root, Percentiles.class);
            break;
        case "rate":
            metricAggregationWithSettings.rate = mapper.convertValue(root, Rate.class);
            break;
        case "raw_data":
            metricAggregationWithSettings.rawData = mapper.convertValue(root, RawData.class);
            break;
        case "raw_document":
            metricAggregationWithSettings.rawDocument = mapper.convertValue(root, RawDocument.class);
            break;
        case "serial_diff":
            metricAggregationWithSettings.serialDiff = mapper.convertValue(root, SerialDiff.class);
            break;
        case "sum":
            metricAggregationWithSettings.sum = mapper.convertValue(root, Sum.class);
            break;
        case "top_metrics":
            metricAggregationWithSettings.topMetrics = mapper.convertValue(root, TopMetrics.class);
            break;
        }
        
        return metricAggregationWithSettings;
    }
}
