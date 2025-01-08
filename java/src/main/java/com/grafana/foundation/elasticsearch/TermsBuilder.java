// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class TermsBuilder implements com.grafana.foundation.cog.Builder<Terms> {
    protected final Terms internal;
    
    public TermsBuilder() {
        this.internal = new Terms();
    this.internal.type = "terms";
    }
    public TermsBuilder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public TermsBuilder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public TermsBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchTermsSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    public Terms build() {
        return this.internal;
    }
}
