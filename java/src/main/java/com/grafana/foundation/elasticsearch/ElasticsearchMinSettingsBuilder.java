// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchMinSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchMinSettings> {
    protected final ElasticsearchMinSettings internal;
    
    public ElasticsearchMinSettingsBuilder() {
        this.internal = new ElasticsearchMinSettings();
    }
    public ElasticsearchMinSettingsBuilder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    
    public ElasticsearchMinSettingsBuilder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    public ElasticsearchMinSettings build() {
        return this.internal;
    }
}
