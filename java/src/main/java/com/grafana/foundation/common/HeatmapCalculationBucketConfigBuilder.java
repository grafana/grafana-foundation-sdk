// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class HeatmapCalculationBucketConfigBuilder implements com.grafana.foundation.cog.Builder<HeatmapCalculationBucketConfig> {
    protected final HeatmapCalculationBucketConfig internal;
    
    public HeatmapCalculationBucketConfigBuilder() {
        this.internal = new HeatmapCalculationBucketConfig();
    }
    public HeatmapCalculationBucketConfigBuilder mode(HeatmapCalculationMode mode) {
        this.internal.mode = mode;
        return this;
    }
    
    public HeatmapCalculationBucketConfigBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    
    public HeatmapCalculationBucketConfigBuilder scale(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scale) {
        this.internal.scale = scale.build();
        return this;
    }
    public HeatmapCalculationBucketConfig build() {
        return this.internal;
    }
}
