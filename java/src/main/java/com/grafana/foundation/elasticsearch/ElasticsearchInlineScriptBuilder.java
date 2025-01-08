// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchInlineScriptBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchInlineScript> {
    protected final ElasticsearchInlineScript internal;
    
    public ElasticsearchInlineScriptBuilder() {
        this.internal = new ElasticsearchInlineScript();
    }
    public ElasticsearchInlineScriptBuilder inline(String inline) {
    this.internal.inline = inline;
        return this;
    }
    public ElasticsearchInlineScript build() {
        return this.internal;
    }
}
