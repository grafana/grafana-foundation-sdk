// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class RawDocumentBuilder implements com.grafana.foundation.cog.Builder<RawDocument> {
    protected final RawDocument internal;
    
    public RawDocumentBuilder() {
        this.internal = new RawDocument();
    }
    public RawDocumentBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public RawDocumentBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchRawDocumentSettings> settings) {
    ElasticsearchRawDocumentSettings settingsResource = settings.build();
        this.internal.settings = settingsResource;
        return this;
    }
    
    public RawDocumentBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public RawDocument build() {
        return this.internal;
    }
}
