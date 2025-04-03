// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class LogsBuilder implements com.grafana.foundation.cog.Builder<Logs> {
    protected final Logs internal;
    
    public LogsBuilder() {
        this.internal = new Logs();
    }
    public LogsBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public LogsBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchLogsSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public LogsBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public Logs build() {
        return this.internal;
    }
}
