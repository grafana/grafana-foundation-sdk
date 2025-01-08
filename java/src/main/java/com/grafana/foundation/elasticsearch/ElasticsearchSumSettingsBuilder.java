// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchSumSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchSumSettings> {
    protected final ElasticsearchSumSettings internal;
    
    public ElasticsearchSumSettingsBuilder() {
        this.internal = new ElasticsearchSumSettings();
    }
    public ElasticsearchSumSettingsBuilder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    
    public ElasticsearchSumSettingsBuilder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    public ElasticsearchSumSettings build() {
        return this.internal;
    }
}
