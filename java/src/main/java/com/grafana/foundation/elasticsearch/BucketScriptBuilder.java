// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;
import java.util.LinkedList;

public class BucketScriptBuilder implements com.grafana.foundation.cog.Builder<BucketScript> {
    protected final BucketScript internal;
    
    public BucketScriptBuilder() {
        this.internal = new BucketScript();
    }
    public BucketScriptBuilder pipelineVariables(List<com.grafana.foundation.cog.Builder<PipelineVariable>> pipelineVariables) {
        List<PipelineVariable> pipelineVariablesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<PipelineVariable> r1 : pipelineVariables) {
                PipelineVariable pipelineVariablesDepth1 = r1.build();
                pipelineVariablesResources.add(pipelineVariablesDepth1); 
        }
        this.internal.pipelineVariables = pipelineVariablesResources;
        return this;
    }
    
    public BucketScriptBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public BucketScriptBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchBucketScriptSettings> settings) {
    ElasticsearchBucketScriptSettings settingsResource = settings.build();
        this.internal.settings = settingsResource;
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
