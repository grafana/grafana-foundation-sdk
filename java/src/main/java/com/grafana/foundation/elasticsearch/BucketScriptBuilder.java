// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;

public class BucketScriptBuilder implements com.grafana.foundation.cog.Builder<BucketScript> {
    protected final BucketScript internal;
    
    public BucketScriptBuilder() {
        this.internal = new BucketScript();
    }
    public BucketScriptBuilder pipelineVariables(com.grafana.foundation.cog.Builder<List<PipelineVariable>> pipelineVariables) {
        this.internal.pipelineVariables = pipelineVariables.build();
        return this;
    }
    
    public BucketScriptBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public BucketScriptBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchBucketScriptSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public BucketScriptBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public BucketScript build() {
        return this.internal;
    }
}
