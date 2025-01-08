// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchAverageSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchAverageSettings> {
    protected final ElasticsearchAverageSettings internal;
    
    public ElasticsearchAverageSettingsBuilder() {
        this.internal = new ElasticsearchAverageSettings();
    }
    public ElasticsearchAverageSettingsBuilder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    
    public ElasticsearchAverageSettingsBuilder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    public ElasticsearchAverageSettings build() {
        return this.internal;
    }
}
