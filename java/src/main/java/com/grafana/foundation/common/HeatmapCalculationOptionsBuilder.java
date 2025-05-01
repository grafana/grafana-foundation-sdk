// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class HeatmapCalculationOptionsBuilder implements com.grafana.foundation.cog.Builder<HeatmapCalculationOptions> {
    protected final HeatmapCalculationOptions internal;
    
    public HeatmapCalculationOptionsBuilder() {
        this.internal = new HeatmapCalculationOptions();
    }
    public HeatmapCalculationOptionsBuilder xBuckets(com.grafana.foundation.cog.Builder<HeatmapCalculationBucketConfig> xBuckets) {
    HeatmapCalculationBucketConfig xBucketsResource = xBuckets.build();
        this.internal.xBuckets = xBucketsResource;
        return this;
    }
    
    public HeatmapCalculationOptionsBuilder yBuckets(com.grafana.foundation.cog.Builder<HeatmapCalculationBucketConfig> yBuckets) {
    HeatmapCalculationBucketConfig yBucketsResource = yBuckets.build();
        this.internal.yBuckets = yBucketsResource;
        return this;
    }
    public HeatmapCalculationOptions build() {
        return this.internal;
    }
}
