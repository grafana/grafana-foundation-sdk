// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class InlineScriptBuilder implements com.grafana.foundation.cog.Builder<InlineScript> {
    protected final InlineScript internal;
    
    public InlineScriptBuilder() {
        this.internal = new InlineScript();
    }
    public InlineScriptBuilder string(String string) {
    this.internal.string = string;
        return this;
    }
    
    public InlineScriptBuilder elasticsearchInlineScript(com.grafana.foundation.cog.Builder<ElasticsearchInlineScript> elasticsearchInlineScript) {
    this.internal.elasticsearchInlineScript = elasticsearchInlineScript.build();
        return this;
    }
    public InlineScript build() {
        return this.internal;
    }
}
