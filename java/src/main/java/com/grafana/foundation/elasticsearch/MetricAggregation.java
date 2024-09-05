// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

@JsonDeserialize(using = MetricAggregationDeserializer.class)
public class MetricAggregation {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Count count;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected MovingAverage movingAverage;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Derivative derivative;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected CumulativeSum cumulativeSum;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected BucketScript bucketScript;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected SerialDiff serialDiff;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected RawData rawData;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected RawDocument rawDocument;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected UniqueCount uniqueCount;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Percentiles percentiles;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected ExtendedStats extendedStats;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Min min;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Max max;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Sum sum;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Average average;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected MovingFunction movingFunction;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Logs logs;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Rate rate;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected TopMetrics topMetrics;
    protected MetricAggregation() {}
    public static MetricAggregation createCount(com.grafana.foundation.cog.Builder<Count> count) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.count = count.build();
        return metricAggregation;
    }
    public static MetricAggregation createMovingAverage(com.grafana.foundation.cog.Builder<MovingAverage> movingAverage) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.movingAverage = movingAverage.build();
        return metricAggregation;
    }
    public static MetricAggregation createDerivative(com.grafana.foundation.cog.Builder<Derivative> derivative) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.derivative = derivative.build();
        return metricAggregation;
    }
    public static MetricAggregation createCumulativeSum(com.grafana.foundation.cog.Builder<CumulativeSum> cumulativeSum) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.cumulativeSum = cumulativeSum.build();
        return metricAggregation;
    }
    public static MetricAggregation createBucketScript(com.grafana.foundation.cog.Builder<BucketScript> bucketScript) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.bucketScript = bucketScript.build();
        return metricAggregation;
    }
    public static MetricAggregation createSerialDiff(com.grafana.foundation.cog.Builder<SerialDiff> serialDiff) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.serialDiff = serialDiff.build();
        return metricAggregation;
    }
    public static MetricAggregation createRawData(com.grafana.foundation.cog.Builder<RawData> rawData) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.rawData = rawData.build();
        return metricAggregation;
    }
    public static MetricAggregation createRawDocument(com.grafana.foundation.cog.Builder<RawDocument> rawDocument) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.rawDocument = rawDocument.build();
        return metricAggregation;
    }
    public static MetricAggregation createUniqueCount(com.grafana.foundation.cog.Builder<UniqueCount> uniqueCount) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.uniqueCount = uniqueCount.build();
        return metricAggregation;
    }
    public static MetricAggregation createPercentiles(com.grafana.foundation.cog.Builder<Percentiles> percentiles) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.percentiles = percentiles.build();
        return metricAggregation;
    }
    public static MetricAggregation createExtendedStats(com.grafana.foundation.cog.Builder<ExtendedStats> extendedStats) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.extendedStats = extendedStats.build();
        return metricAggregation;
    }
    public static MetricAggregation createMin(com.grafana.foundation.cog.Builder<Min> min) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.min = min.build();
        return metricAggregation;
    }
    public static MetricAggregation createMax(com.grafana.foundation.cog.Builder<Max> max) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.max = max.build();
        return metricAggregation;
    }
    public static MetricAggregation createSum(com.grafana.foundation.cog.Builder<Sum> sum) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.sum = sum.build();
        return metricAggregation;
    }
    public static MetricAggregation createAverage(com.grafana.foundation.cog.Builder<Average> average) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.average = average.build();
        return metricAggregation;
    }
    public static MetricAggregation createMovingFunction(com.grafana.foundation.cog.Builder<MovingFunction> movingFunction) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.movingFunction = movingFunction.build();
        return metricAggregation;
    }
    public static MetricAggregation createLogs(com.grafana.foundation.cog.Builder<Logs> logs) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.logs = logs.build();
        return metricAggregation;
    }
    public static MetricAggregation createRate(com.grafana.foundation.cog.Builder<Rate> rate) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.rate = rate.build();
        return metricAggregation;
    }
    public static MetricAggregation createTopMetrics(com.grafana.foundation.cog.Builder<TopMetrics> topMetrics) {
        MetricAggregation metricAggregation = new MetricAggregation();
        metricAggregation.topMetrics = topMetrics.build();
        return metricAggregation;
    }
    
    public String toJSON() throws JsonProcessingException {
        if (count != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(count);
        }
        if (movingAverage != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(movingAverage);
        }
        if (derivative != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(derivative);
        }
        if (cumulativeSum != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(cumulativeSum);
        }
        if (bucketScript != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(bucketScript);
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
