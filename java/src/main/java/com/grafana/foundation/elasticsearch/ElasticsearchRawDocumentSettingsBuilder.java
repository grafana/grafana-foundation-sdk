// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchRawDocumentSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchRawDocumentSettings> {
    protected final ElasticsearchRawDocumentSettings internal;
    
    public ElasticsearchRawDocumentSettingsBuilder() {
        this.internal = new ElasticsearchRawDocumentSettings();
    }
    public ElasticsearchRawDocumentSettingsBuilder size(String size) {
    this.internal.size = size;
        return this;
    }
    public ElasticsearchRawDocumentSettings build() {
        return this.internal;
    }
}
