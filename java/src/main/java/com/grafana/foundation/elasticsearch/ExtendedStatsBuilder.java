// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ExtendedStatsBuilder implements com.grafana.foundation.cog.Builder<ExtendedStats> {
    protected final ExtendedStats internal;
    
    public ExtendedStatsBuilder() {
        this.internal = new ExtendedStats();
    }
    public ExtendedStatsBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchExtendedStatsSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public ExtendedStatsBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public ExtendedStatsBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public ExtendedStatsBuilder meta(Object meta) {
        this.internal.meta = meta;
        return this;
    }
    
    public ExtendedStatsBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public ExtendedStats build() {
        return this.internal;
    }
}
