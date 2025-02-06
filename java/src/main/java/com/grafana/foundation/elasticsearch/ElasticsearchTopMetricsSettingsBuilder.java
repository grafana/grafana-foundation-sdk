// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;

public class ElasticsearchTopMetricsSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchTopMetricsSettings> {
    protected final ElasticsearchTopMetricsSettings internal;
    
    public ElasticsearchTopMetricsSettingsBuilder() {
        this.internal = new ElasticsearchTopMetricsSettings();
    }
    public ElasticsearchTopMetricsSettingsBuilder order(String order) {
        this.internal.order = order;
        return this;
    }
    
    public ElasticsearchTopMetricsSettingsBuilder orderBy(String orderBy) {
        this.internal.orderBy = orderBy;
        return this;
    }
    
    public ElasticsearchTopMetricsSettingsBuilder metrics(List<String> metrics) {
        this.internal.metrics = metrics;
        return this;
    }
    public ElasticsearchTopMetricsSettings build() {
        return this.internal;
    }
}
