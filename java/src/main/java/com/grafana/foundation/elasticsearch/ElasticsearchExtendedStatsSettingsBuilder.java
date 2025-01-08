// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchExtendedStatsSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchExtendedStatsSettings> {
    protected final ElasticsearchExtendedStatsSettings internal;
    
    public ElasticsearchExtendedStatsSettingsBuilder() {
        this.internal = new ElasticsearchExtendedStatsSettings();
    }
    public ElasticsearchExtendedStatsSettingsBuilder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    
    public ElasticsearchExtendedStatsSettingsBuilder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    
    public ElasticsearchExtendedStatsSettingsBuilder sigma(String sigma) {
    this.internal.sigma = sigma;
        return this;
    }
    public ElasticsearchExtendedStatsSettings build() {
        return this.internal;
    }
}
