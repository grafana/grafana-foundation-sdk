// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class ScaleDistributionConfigBuilder implements com.grafana.foundation.cog.Builder<ScaleDistributionConfig> {
    protected final ScaleDistributionConfig internal;
    
    public ScaleDistributionConfigBuilder() {
        this.internal = new ScaleDistributionConfig();
    }
    public ScaleDistributionConfigBuilder type(ScaleDistribution type) {
        this.internal.type = type;
        return this;
    }
    
    public ScaleDistributionConfigBuilder log(Double log) {
        this.internal.log = log;
        return this;
    }
    
    public ScaleDistributionConfigBuilder linearThreshold(Double linearThreshold) {
        this.internal.linearThreshold = linearThreshold;
        return this;
    }
    public ScaleDistributionConfig build() {
        return this.internal;
    }
}
