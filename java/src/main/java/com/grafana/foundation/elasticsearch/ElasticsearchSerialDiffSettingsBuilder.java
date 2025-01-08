// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchSerialDiffSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchSerialDiffSettings> {
    protected final ElasticsearchSerialDiffSettings internal;
    
    public ElasticsearchSerialDiffSettingsBuilder() {
        this.internal = new ElasticsearchSerialDiffSettings();
    }
    public ElasticsearchSerialDiffSettingsBuilder lag(String lag) {
    this.internal.lag = lag;
        return this;
    }
    public ElasticsearchSerialDiffSettings build() {
        return this.internal;
    }
}
