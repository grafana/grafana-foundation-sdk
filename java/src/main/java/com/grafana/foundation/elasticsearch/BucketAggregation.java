// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = BucketAggregationDeserializer.class)
public class BucketAggregation { 
    @JsonUnwrapped
    public DateHistogram dateHistogram; 
    @JsonUnwrapped
    public Histogram histogram; 
    @JsonUnwrapped
    public Terms terms; 
    @JsonUnwrapped
    public Filters filters; 
    @JsonUnwrapped
    public GeoHashGrid geoHashGrid; 
    @JsonUnwrapped
    public Nested nested;
    
    public String toJSON() throws JsonProcessingException {
        if (dateHistogram != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(dateHistogram);
        }
        if (histogram != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(histogram);
        }
        if (terms != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(terms);
        }
        if (filters != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(filters);
        }
        if (geoHashGrid != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(geoHashGrid);
        }
        if (nested != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(nested);
        }
        
        return null;
    }

}
