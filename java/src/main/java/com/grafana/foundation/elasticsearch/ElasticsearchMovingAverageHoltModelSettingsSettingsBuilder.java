// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchMovingAverageHoltModelSettingsSettings> {
    protected final ElasticsearchMovingAverageHoltModelSettingsSettings internal;
    
    public ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder() {
        this.internal = new ElasticsearchMovingAverageHoltModelSettingsSettings();
    }
    public ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder alpha(String alpha) {
        this.internal.alpha = alpha;
        return this;
    }
    
    public ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder beta(String beta) {
        this.internal.beta = beta;
        return this;
    }
    public ElasticsearchMovingAverageHoltModelSettingsSettings build() {
        return this.internal;
    }
}
