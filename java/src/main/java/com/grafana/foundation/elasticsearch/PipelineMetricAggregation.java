// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = PipelineMetricAggregationDeserializer.class)
public class PipelineMetricAggregation {
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
    protected PipelineMetricAggregation() {}
    public static PipelineMetricAggregation createMovingAverage(MovingAverage movingAverage) {
        PipelineMetricAggregation pipelineMetricAggregation = new PipelineMetricAggregation();
        pipelineMetricAggregation.movingAverage = movingAverage;
        return pipelineMetricAggregation;
    }
    public static PipelineMetricAggregation createDerivative(Derivative derivative) {
        PipelineMetricAggregation pipelineMetricAggregation = new PipelineMetricAggregation();
        pipelineMetricAggregation.derivative = derivative;
        return pipelineMetricAggregation;
    }
    public static PipelineMetricAggregation createCumulativeSum(CumulativeSum cumulativeSum) {
        PipelineMetricAggregation pipelineMetricAggregation = new PipelineMetricAggregation();
        pipelineMetricAggregation.cumulativeSum = cumulativeSum;
        return pipelineMetricAggregation;
    }
    public static PipelineMetricAggregation createBucketScript(BucketScript bucketScript) {
        PipelineMetricAggregation pipelineMetricAggregation = new PipelineMetricAggregation();
        pipelineMetricAggregation.bucketScript = bucketScript;
        return pipelineMetricAggregation;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
