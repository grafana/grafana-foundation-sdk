// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class BaseBucketAggregationBuilder implements com.grafana.foundation.cog.Builder<BaseBucketAggregation> {
    protected final BaseBucketAggregation internal;
    
    public BaseBucketAggregationBuilder() {
        this.internal = new BaseBucketAggregation();
    }
    public BaseBucketAggregationBuilder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public BaseBucketAggregationBuilder type(BucketAggregationType type) {
    this.internal.type = type;
        return this;
    }
    
    public BaseBucketAggregationBuilder settings(Object settings) {
    this.internal.settings = settings;
        return this;
    }
    public BaseBucketAggregation build() {
        return this.internal;
    }
}
