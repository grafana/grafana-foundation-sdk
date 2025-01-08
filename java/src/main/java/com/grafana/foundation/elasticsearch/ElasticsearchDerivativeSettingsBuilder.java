// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchDerivativeSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchDerivativeSettings> {
    protected final ElasticsearchDerivativeSettings internal;
    
    public ElasticsearchDerivativeSettingsBuilder() {
        this.internal = new ElasticsearchDerivativeSettings();
    }
    public ElasticsearchDerivativeSettingsBuilder unit(String unit) {
    this.internal.unit = unit;
        return this;
    }
    public ElasticsearchDerivativeSettings build() {
        return this.internal;
    }
}
