// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class TopMetricsBuilder implements com.grafana.foundation.cog.Builder<TopMetrics> {
    protected final TopMetrics internal;
    
    public TopMetricsBuilder() {
        this.internal = new TopMetrics();
        this.internal.type = "top_metrics";
    }
    public TopMetricsBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public TopMetricsBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchTopMetricsSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public TopMetricsBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public TopMetrics build() {
        return this.internal;
    }
}
