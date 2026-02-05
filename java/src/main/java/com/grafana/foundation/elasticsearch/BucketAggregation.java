// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = BucketAggregationDeserializer.class)
public class BucketAggregation {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected DateHistogram dateHistogram;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Histogram histogram;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Terms terms;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Filters filters;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected GeoHashGrid geoHashGrid;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected Nested nested;
    protected BucketAggregation() {}
    public static BucketAggregation createDateHistogram(DateHistogram dateHistogram) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.dateHistogram = dateHistogram;
        return bucketAggregation;
    }
    public static BucketAggregation createHistogram(Histogram histogram) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.histogram = histogram;
        return bucketAggregation;
    }
    public static BucketAggregation createTerms(Terms terms) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.terms = terms;
        return bucketAggregation;
    }
    public static BucketAggregation createFilters(Filters filters) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.filters = filters;
        return bucketAggregation;
    }
    public static BucketAggregation createGeoHashGrid(GeoHashGrid geoHashGrid) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.geoHashGrid = geoHashGrid;
        return bucketAggregation;
    }
    public static BucketAggregation createNested(Nested nested) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.nested = nested;
        return bucketAggregation;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
