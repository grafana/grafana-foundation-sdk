// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = PipelineMetricAggregationDeserializer.class)
public class PipelineMetricAggregation { 
    @JsonUnwrapped
    public MovingAverage movingAverage; 
    @JsonUnwrapped
    public Derivative derivative; 
    @JsonUnwrapped
    public CumulativeSum cumulativeSum; 
    @JsonUnwrapped
    public BucketScript bucketScript;
    
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
