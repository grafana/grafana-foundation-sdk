// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;


public class MatcherBuilder implements com.grafana.foundation.cog.Builder<Matcher> {
    protected final Matcher internal;
    
    public MatcherBuilder() {
        this.internal = new Matcher();
    }
    public MatcherBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public MatcherBuilder type(MatchType type) {
    this.internal.type = type;
        return this;
    }
    
    public MatcherBuilder value(String value) {
    this.internal.value = value;
        return this;
    }
    public Matcher build() {
        return this.internal;
    }
}
