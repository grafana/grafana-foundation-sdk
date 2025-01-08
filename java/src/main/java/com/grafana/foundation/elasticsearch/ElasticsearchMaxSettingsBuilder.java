// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchMaxSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchMaxSettings> {
    protected final ElasticsearchMaxSettings internal;
    
    public ElasticsearchMaxSettingsBuilder() {
        this.internal = new ElasticsearchMaxSettings();
    }
    public ElasticsearchMaxSettingsBuilder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    
    public ElasticsearchMaxSettingsBuilder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    public ElasticsearchMaxSettings build() {
        return this.internal;
    }
}
