// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class GraphThresholdsStyleConfigBuilder implements com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> {
    protected final GraphThresholdsStyleConfig internal;
    
    public GraphThresholdsStyleConfigBuilder() {
        this.internal = new GraphThresholdsStyleConfig();
    }
    public GraphThresholdsStyleConfigBuilder mode(GraphThresholdsStyleMode mode) {
    this.internal.mode = mode;
        return this;
    }
    public GraphThresholdsStyleConfig build() {
        return this.internal;
    }
}
