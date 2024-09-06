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
    public static PipelineMetricAggregation createMovingAverage(com.grafana.foundation.cog.Builder<MovingAverage> movingAverage) {
        PipelineMetricAggregation pipelineMetricAggregation = new PipelineMetricAggregation();
        pipelineMetricAggregation.movingAverage = movingAverage.build();
        return pipelineMetricAggregation;
    }
    public static PipelineMetricAggregation createDerivative(com.grafana.foundation.cog.Builder<Derivative> derivative) {
        PipelineMetricAggregation pipelineMetricAggregation = new PipelineMetricAggregation();
        pipelineMetricAggregation.derivative = derivative.build();
        return pipelineMetricAggregation;
    }
    public static PipelineMetricAggregation createCumulativeSum(com.grafana.foundation.cog.Builder<CumulativeSum> cumulativeSum) {
        PipelineMetricAggregation pipelineMetricAggregation = new PipelineMetricAggregation();
        pipelineMetricAggregation.cumulativeSum = cumulativeSum.build();
        return pipelineMetricAggregation;
    }
    public static PipelineMetricAggregation createBucketScript(com.grafana.foundation.cog.Builder<BucketScript> bucketScript) {
        PipelineMetricAggregation pipelineMetricAggregation = new PipelineMetricAggregation();
        pipelineMetricAggregation.bucketScript = bucketScript.build();
        return pipelineMetricAggregation;
    }
    
    public String toJSON() throws JsonProcessingException {
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
        
        return null;
    }

}
