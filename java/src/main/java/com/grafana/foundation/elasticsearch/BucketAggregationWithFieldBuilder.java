// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class BucketAggregationWithFieldBuilder implements com.grafana.foundation.cog.Builder<BucketAggregationWithField> {
    protected final BucketAggregationWithField internal;
    
    public BucketAggregationWithFieldBuilder() {
        this.internal = new BucketAggregationWithField();
    }
    public BucketAggregationWithFieldBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public BucketAggregationWithFieldBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public BucketAggregationWithFieldBuilder type(BucketAggregationType type) {
        this.internal.type = type;
        return this;
    }
    
    public BucketAggregationWithFieldBuilder settings(Object settings) {
        this.internal.settings = settings;
        return this;
    }
    public BucketAggregationWithField build() {
        return this.internal;
    }
}
