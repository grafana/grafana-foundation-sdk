// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchLogsSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchLogsSettings> {
    protected final ElasticsearchLogsSettings internal;
    
    public ElasticsearchLogsSettingsBuilder() {
        this.internal = new ElasticsearchLogsSettings();
    }
    public ElasticsearchLogsSettingsBuilder limit(String limit) {
        this.internal.limit = limit;
        return this;
    }
    public ElasticsearchLogsSettings build() {
        return this.internal;
    }
}
