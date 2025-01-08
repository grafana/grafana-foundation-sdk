// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchUniqueCountSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchUniqueCountSettings> {
    protected final ElasticsearchUniqueCountSettings internal;
    
    public ElasticsearchUniqueCountSettingsBuilder() {
        this.internal = new ElasticsearchUniqueCountSettings();
    }
    public ElasticsearchUniqueCountSettingsBuilder precisionThreshold(String precisionThreshold) {
    this.internal.precisionThreshold = precisionThreshold;
        return this;
    }
    
    public ElasticsearchUniqueCountSettingsBuilder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    public ElasticsearchUniqueCountSettings build() {
        return this.internal;
    }
}
