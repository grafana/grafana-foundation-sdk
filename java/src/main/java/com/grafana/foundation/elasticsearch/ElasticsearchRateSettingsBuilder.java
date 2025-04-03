// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchRateSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchRateSettings> {
    protected final ElasticsearchRateSettings internal;
    
    public ElasticsearchRateSettingsBuilder() {
        this.internal = new ElasticsearchRateSettings();
    }
    public ElasticsearchRateSettingsBuilder unit(String unit) {
        this.internal.unit = unit;
        return this;
    }
    
    public ElasticsearchRateSettingsBuilder mode(String mode) {
        this.internal.mode = mode;
        return this;
    }
    public ElasticsearchRateSettings build() {
        return this.internal;
    }
}
