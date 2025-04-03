// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchMovingAverageEWMAModelSettingsSettings> {
    protected final ElasticsearchMovingAverageEWMAModelSettingsSettings internal;
    
    public ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder() {
        this.internal = new ElasticsearchMovingAverageEWMAModelSettingsSettings();
    }
    public ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder alpha(String alpha) {
        this.internal.alpha = alpha;
        return this;
    }
    public ElasticsearchMovingAverageEWMAModelSettingsSettings build() {
        return this.internal;
    }
}
