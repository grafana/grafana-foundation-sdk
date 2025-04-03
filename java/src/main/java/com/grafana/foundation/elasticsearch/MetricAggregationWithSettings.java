// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = MetricAggregationWithSettingsDeserializer.class)
public class MetricAggregationWithSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected BucketScript bucketScript;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected CumulativeSum cumulativeSum;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Derivative derivative;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected SerialDiff serialDiff;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected RawData rawData;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected RawDocument rawDocument;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected UniqueCount uniqueCount;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Percentiles percentiles;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected ExtendedStats extendedStats;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Min min;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Max max;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Sum sum;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Average average;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected MovingAverage movingAverage;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected MovingFunction movingFunction;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Logs logs;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Rate rate;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TopMetrics topMetrics;
    protected MetricAggregationWithSettings() {}
    public static MetricAggregationWithSettings createBucketScript(com.grafana.foundation.cog.Builder<BucketScript> bucketScript) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.bucketScript = bucketScript.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createCumulativeSum(com.grafana.foundation.cog.Builder<CumulativeSum> cumulativeSum) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.cumulativeSum = cumulativeSum.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createDerivative(com.grafana.foundation.cog.Builder<Derivative> derivative) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.derivative = derivative.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createSerialDiff(com.grafana.foundation.cog.Builder<SerialDiff> serialDiff) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.serialDiff = serialDiff.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createRawData(com.grafana.foundation.cog.Builder<RawData> rawData) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.rawData = rawData.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createRawDocument(com.grafana.foundation.cog.Builder<RawDocument> rawDocument) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.rawDocument = rawDocument.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createUniqueCount(com.grafana.foundation.cog.Builder<UniqueCount> uniqueCount) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.uniqueCount = uniqueCount.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createPercentiles(com.grafana.foundation.cog.Builder<Percentiles> percentiles) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.percentiles = percentiles.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createExtendedStats(com.grafana.foundation.cog.Builder<ExtendedStats> extendedStats) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.extendedStats = extendedStats.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createMin(com.grafana.foundation.cog.Builder<Min> min) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.min = min.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createMax(com.grafana.foundation.cog.Builder<Max> max) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.max = max.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createSum(com.grafana.foundation.cog.Builder<Sum> sum) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.sum = sum.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createAverage(com.grafana.foundation.cog.Builder<Average> average) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.average = average.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createMovingAverage(com.grafana.foundation.cog.Builder<MovingAverage> movingAverage) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.movingAverage = movingAverage.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createMovingFunction(com.grafana.foundation.cog.Builder<MovingFunction> movingFunction) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.movingFunction = movingFunction.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createLogs(com.grafana.foundation.cog.Builder<Logs> logs) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.logs = logs.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createRate(com.grafana.foundation.cog.Builder<Rate> rate) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.rate = rate.build();
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createTopMetrics(com.grafana.foundation.cog.Builder<TopMetrics> topMetrics) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.topMetrics = topMetrics.build();
        return metricAggregationWithSettings;
    }
    
    public String toJSON() throws JsonProcessingException {
        if (bucketScript != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(bucketScript);
        }
        if (cumulativeSum != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(cumulativeSum);
        }
        if (derivative != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(derivative);
        }
        if (serialDiff != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(serialDiff);
        }
        if (rawData != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(rawData);
        }
        if (rawDocument != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(rawDocument);
        }
        if (uniqueCount != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(uniqueCount);
        }
        if (percentiles != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(percentiles);
        }
        if (extendedStats != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(extendedStats);
        }
        if (min != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(min);
        }
        if (max != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(max);
        }
        if (sum != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(sum);
        }
        if (average != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(average);
        }
        if (movingAverage != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(movingAverage);
        }
        if (movingFunction != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(movingFunction);
        }
        if (logs != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(logs);
        }
        if (rate != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(rate);
        }
        if (topMetrics != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(topMetrics);
        }
        
        return null;
    }

}
