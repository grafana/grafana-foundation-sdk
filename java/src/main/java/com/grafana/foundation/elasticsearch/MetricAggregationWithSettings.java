// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = MetricAggregationWithSettingsDeserializer.class)
public class MetricAggregationWithSettings { 
    @JsonUnwrapped
    public BucketScript bucketScript; 
    @JsonUnwrapped
    public CumulativeSum cumulativeSum; 
    @JsonUnwrapped
    public Derivative derivative; 
    @JsonUnwrapped
    public SerialDiff serialDiff; 
    @JsonUnwrapped
    public RawData rawData; 
    @JsonUnwrapped
    public RawDocument rawDocument; 
    @JsonUnwrapped
    public UniqueCount uniqueCount; 
    @JsonUnwrapped
    public Percentiles percentiles; 
    @JsonUnwrapped
    public ExtendedStats extendedStats; 
    @JsonUnwrapped
    public Min min; 
    @JsonUnwrapped
    public Max max; 
    @JsonUnwrapped
    public Sum sum; 
    @JsonUnwrapped
    public Average average; 
    @JsonUnwrapped
    public MovingAverage movingAverage; 
    @JsonUnwrapped
    public MovingFunction movingFunction; 
    @JsonUnwrapped
    public Logs logs; 
    @JsonUnwrapped
    public Rate rate; 
    @JsonUnwrapped
    public TopMetrics topMetrics;
    
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
