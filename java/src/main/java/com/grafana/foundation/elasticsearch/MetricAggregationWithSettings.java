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
    public static MetricAggregationWithSettings createBucketScript(BucketScript bucketScript) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.bucketScript = bucketScript;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createCumulativeSum(CumulativeSum cumulativeSum) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.cumulativeSum = cumulativeSum;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createDerivative(Derivative derivative) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.derivative = derivative;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createSerialDiff(SerialDiff serialDiff) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.serialDiff = serialDiff;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createRawData(RawData rawData) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.rawData = rawData;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createRawDocument(RawDocument rawDocument) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.rawDocument = rawDocument;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createUniqueCount(UniqueCount uniqueCount) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.uniqueCount = uniqueCount;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createPercentiles(Percentiles percentiles) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.percentiles = percentiles;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createExtendedStats(ExtendedStats extendedStats) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.extendedStats = extendedStats;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createMin(Min min) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.min = min;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createMax(Max max) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.max = max;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createSum(Sum sum) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.sum = sum;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createAverage(Average average) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.average = average;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createMovingAverage(MovingAverage movingAverage) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.movingAverage = movingAverage;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createMovingFunction(MovingFunction movingFunction) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.movingFunction = movingFunction;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createLogs(Logs logs) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.logs = logs;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createRate(Rate rate) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.rate = rate;
        return metricAggregationWithSettings;
    }
    public static MetricAggregationWithSettings createTopMetrics(TopMetrics topMetrics) {
        MetricAggregationWithSettings metricAggregationWithSettings = new MetricAggregationWithSettings();
        metricAggregationWithSettings.topMetrics = topMetrics;
        return metricAggregationWithSettings;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
