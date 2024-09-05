// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

@JsonDeserialize(using = BucketAggregationDeserializer.class)
public class BucketAggregation {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected DateHistogram dateHistogram;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Histogram histogram;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Terms terms;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Filters filters;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected GeoHashGrid geoHashGrid;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonUnwrapped
    protected Nested nested;
    protected BucketAggregation() {}
    public static BucketAggregation createDateHistogram(com.grafana.foundation.cog.Builder<DateHistogram> dateHistogram) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.dateHistogram = dateHistogram.build();
        return bucketAggregation;
    }
    public static BucketAggregation createHistogram(com.grafana.foundation.cog.Builder<Histogram> histogram) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.histogram = histogram.build();
        return bucketAggregation;
    }
    public static BucketAggregation createTerms(com.grafana.foundation.cog.Builder<Terms> terms) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.terms = terms.build();
        return bucketAggregation;
    }
    public static BucketAggregation createFilters(com.grafana.foundation.cog.Builder<Filters> filters) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.filters = filters.build();
        return bucketAggregation;
    }
    public static BucketAggregation createGeoHashGrid(com.grafana.foundation.cog.Builder<GeoHashGrid> geoHashGrid) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.geoHashGrid = geoHashGrid.build();
        return bucketAggregation;
    }
    public static BucketAggregation createNested(com.grafana.foundation.cog.Builder<Nested> nested) {
        BucketAggregation bucketAggregation = new BucketAggregation();
        bucketAggregation.nested = nested.build();
        return bucketAggregation;
    }
    
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
