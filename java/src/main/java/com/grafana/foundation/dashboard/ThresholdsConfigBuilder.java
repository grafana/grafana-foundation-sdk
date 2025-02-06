// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;

public class ThresholdsConfigBuilder implements com.grafana.foundation.cog.Builder<ThresholdsConfig> {
    protected final ThresholdsConfig internal;
    
    public ThresholdsConfigBuilder() {
        this.internal = new ThresholdsConfig();
    }
    public ThresholdsConfigBuilder mode(ThresholdsMode mode) {
        this.internal.mode = mode;
        return this;
    }
    
    public ThresholdsConfigBuilder steps(List<Threshold> steps) {
        this.internal.steps = steps;
        return this;
    }
    public ThresholdsConfig build() {
        return this.internal;
    }
}
