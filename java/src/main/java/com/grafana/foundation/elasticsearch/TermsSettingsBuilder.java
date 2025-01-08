// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class TermsSettingsBuilder implements com.grafana.foundation.cog.Builder<TermsSettings> {
    protected final TermsSettings internal;
    
    public TermsSettingsBuilder() {
        this.internal = new TermsSettings();
    }
    public TermsSettingsBuilder order(TermsOrder order) {
    this.internal.order = order;
        return this;
    }
    
    public TermsSettingsBuilder size(String size) {
    this.internal.size = size;
        return this;
    }
    
    public TermsSettingsBuilder minDocCount(String minDocCount) {
    this.internal.minDocCount = minDocCount;
        return this;
    }
    
    public TermsSettingsBuilder orderBy(String orderBy) {
    this.internal.orderBy = orderBy;
        return this;
    }
    
    public TermsSettingsBuilder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    public TermsSettings build() {
        return this.internal;
    }
}
