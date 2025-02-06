// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class HeatmapCalculationOptionsBuilder implements com.grafana.foundation.cog.Builder<HeatmapCalculationOptions> {
    protected final HeatmapCalculationOptions internal;
    
    public HeatmapCalculationOptionsBuilder() {
        this.internal = new HeatmapCalculationOptions();
    }
    public HeatmapCalculationOptionsBuilder xBuckets(com.grafana.foundation.cog.Builder<HeatmapCalculationBucketConfig> xBuckets) {
        this.internal.xBuckets = xBuckets.build();
        return this;
    }
    
    public HeatmapCalculationOptionsBuilder yBuckets(com.grafana.foundation.cog.Builder<HeatmapCalculationBucketConfig> yBuckets) {
        this.internal.yBuckets = yBuckets.build();
        return this;
    }
    public HeatmapCalculationOptions build() {
        return this.internal;
    }
}
