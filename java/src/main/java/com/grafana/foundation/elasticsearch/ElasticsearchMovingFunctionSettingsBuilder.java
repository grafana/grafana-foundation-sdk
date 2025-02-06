// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchMovingFunctionSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchMovingFunctionSettings> {
    protected final ElasticsearchMovingFunctionSettings internal;
    
    public ElasticsearchMovingFunctionSettingsBuilder() {
        this.internal = new ElasticsearchMovingFunctionSettings();
    }
    public ElasticsearchMovingFunctionSettingsBuilder window(String window) {
        this.internal.window = window;
        return this;
    }
    
    public ElasticsearchMovingFunctionSettingsBuilder script(InlineScript script) {
        this.internal.script = script;
        return this;
    }
    
    public ElasticsearchMovingFunctionSettingsBuilder shift(String shift) {
        this.internal.shift = shift;
        return this;
    }
    public ElasticsearchMovingFunctionSettings build() {
        return this.internal;
    }
}
