// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchMovingAverageHoltWintersModelSettingsSettings> {
    protected final ElasticsearchMovingAverageHoltWintersModelSettingsSettings internal;
    
    public ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder() {
        this.internal = new ElasticsearchMovingAverageHoltWintersModelSettingsSettings();
    }
    public ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder alpha(String alpha) {
        this.internal.alpha = alpha;
        return this;
    }
    
    public ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder beta(String beta) {
        this.internal.beta = beta;
        return this;
    }
    
    public ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder gamma(String gamma) {
        this.internal.gamma = gamma;
        return this;
    }
    
    public ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder period(String period) {
        this.internal.period = period;
        return this;
    }
    
    public ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder pad(Boolean pad) {
        this.internal.pad = pad;
        return this;
    }
    public ElasticsearchMovingAverageHoltWintersModelSettingsSettings build() {
        return this.internal;
    }
}
