// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchTermsSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchTermsSettings> {
    protected final ElasticsearchTermsSettings internal;
    
    public ElasticsearchTermsSettingsBuilder() {
        this.internal = new ElasticsearchTermsSettings();
    }
    public ElasticsearchTermsSettingsBuilder order(TermsOrder order) {
    this.internal.order = order;
        return this;
    }
    
    public ElasticsearchTermsSettingsBuilder size(String size) {
    this.internal.size = size;
        return this;
    }
    
    public ElasticsearchTermsSettingsBuilder minDocCount(String minDocCount) {
    this.internal.minDocCount = minDocCount;
        return this;
    }
    
    public ElasticsearchTermsSettingsBuilder orderBy(String orderBy) {
    this.internal.orderBy = orderBy;
        return this;
    }
    
    public ElasticsearchTermsSettingsBuilder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    public ElasticsearchTermsSettings build() {
        return this.internal;
    }
}
