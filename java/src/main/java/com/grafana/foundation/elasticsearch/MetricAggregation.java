// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = MetricAggregationDeserializer.class)
public class MetricAggregation {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Count count;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected MovingAverage movingAverage;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Derivative derivative;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected CumulativeSum cumulativeSum;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected BucketScript bucketScript;
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
    protected MetricAggregation() {}
    public static MetricAggregation createCount(Count count) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.count = count;
        return metricAggregation;
    }
    public static MetricAggregation createMovingAverage(MovingAverage movingAverage) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.movingAverage = movingAverage;
        return metricAggregation;
    }
    public static MetricAggregation createDerivative(Derivative derivative) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.derivative = derivative;
        return metricAggregation;
    }
    public static MetricAggregation createCumulativeSum(CumulativeSum cumulativeSum) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.cumulativeSum = cumulativeSum;
        return metricAggregation;
    }
    public static MetricAggregation createBucketScript(BucketScript bucketScript) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.bucketScript = bucketScript;
        return metricAggregation;
    }
    public static MetricAggregation createSerialDiff(SerialDiff serialDiff) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.serialDiff = serialDiff;
        return metricAggregation;
    }
    public static MetricAggregation createRawData(RawData rawData) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.rawData = rawData;
        return metricAggregation;
    }
    public static MetricAggregation createRawDocument(RawDocument rawDocument) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.rawDocument = rawDocument;
        return metricAggregation;
    }
    public static MetricAggregation createUniqueCount(UniqueCount uniqueCount) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.uniqueCount = uniqueCount;
        return metricAggregation;
    }
    public static MetricAggregation createPercentiles(Percentiles percentiles) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.percentiles = percentiles;
        return metricAggregation;
    }
    public static MetricAggregation createExtendedStats(ExtendedStats extendedStats) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.extendedStats = extendedStats;
        return metricAggregation;
    }
    public static MetricAggregation createMin(Min min) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.min = min;
        return metricAggregation;
    }
    public static MetricAggregation createMax(Max max) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.max = max;
        return metricAggregation;
    }
    public static MetricAggregation createSum(Sum sum) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.sum = sum;
        return metricAggregation;
    }
    public static MetricAggregation createAverage(Average average) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.average = average;
        return metricAggregation;
    }
    public static MetricAggregation createMovingFunction(MovingFunction movingFunction) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.movingFunction = movingFunction;
        return metricAggregation;
    }
    public static MetricAggregation createLogs(Logs logs) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.logs = logs;
        return metricAggregation;
    }
    public static MetricAggregation createRate(Rate rate) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.rate = rate;
        return metricAggregation;
    }
    public static MetricAggregation createTopMetrics(TopMetrics topMetrics) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.topMetrics = topMetrics;
        return metricAggregation;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
