// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferencesv1alpha1;


public class CookiePreferencesBuilder implements com.grafana.foundation.cog.Builder<CookiePreferences> {
    protected final CookiePreferences internal;
    
    public CookiePreferencesBuilder() {
        this.internal = new CookiePreferences();
    }
    public CookiePreferencesBuilder analytics(Object analytics) {
        this.internal.analytics = analytics;
        return this;
    }
    
    public CookiePreferencesBuilder performance(Object performance) {
        this.internal.performance = performance;
        return this;
    }
    
    public CookiePreferencesBuilder functional(Object functional) {
        this.internal.functional = functional;
        return this;
    }
    public CookiePreferences build() {
        return this.internal;
    }
}
