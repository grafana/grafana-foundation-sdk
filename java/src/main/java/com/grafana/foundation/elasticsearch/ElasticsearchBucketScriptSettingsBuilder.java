// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchBucketScriptSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchBucketScriptSettings> {
    protected final ElasticsearchBucketScriptSettings internal;
    
    public ElasticsearchBucketScriptSettingsBuilder() {
        this.internal = new ElasticsearchBucketScriptSettings();
    }
    public ElasticsearchBucketScriptSettingsBuilder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    public ElasticsearchBucketScriptSettings build() {
        return this.internal;
    }
}
