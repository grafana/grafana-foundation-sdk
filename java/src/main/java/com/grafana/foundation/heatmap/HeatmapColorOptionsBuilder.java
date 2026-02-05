// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;


public class HeatmapColorOptionsBuilder implements com.grafana.foundation.cog.Builder<HeatmapColorOptions> {
    protected final HeatmapColorOptions internal;
    
    public HeatmapColorOptionsBuilder() {
        this.internal = new HeatmapColorOptions();
    }
    public HeatmapColorOptionsBuilder mode(HeatmapColorMode mode) {
        this.internal.mode = mode;
        return this;
    }
    
    public HeatmapColorOptionsBuilder scheme(String scheme) {
        this.internal.scheme = scheme;
        return this;
    }
    
    public HeatmapColorOptionsBuilder fill(String fill) {
        this.internal.fill = fill;
        return this;
    }
    
    public HeatmapColorOptionsBuilder scale(HeatmapColorScale scale) {
        this.internal.scale = scale;
        return this;
    }
    
    public HeatmapColorOptionsBuilder exponent(Float exponent) {
        this.internal.exponent = exponent;
        return this;
    }
    
    public HeatmapColorOptionsBuilder steps(Long steps) {
        if (!(steps >= 2)) {
            throw new IllegalArgumentException("steps must be >= 2");
        }
        if (!(steps <= 128)) {
            throw new IllegalArgumentException("steps must be <= 128");
        }
        this.internal.steps = steps;
        return this;
    }
    
    public HeatmapColorOptionsBuilder reverse(Boolean reverse) {
        this.internal.reverse = reverse;
        return this;
    }
    
    public HeatmapColorOptionsBuilder min(Float min) {
        this.internal.min = min;
        return this;
    }
    
    public HeatmapColorOptionsBuilder max(Float max) {
        this.internal.max = max;
        return this;
    }
    public HeatmapColorOptions build() {
        return this.internal;
    }
}
