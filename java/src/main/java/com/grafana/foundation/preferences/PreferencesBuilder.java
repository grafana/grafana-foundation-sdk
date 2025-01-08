// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferences;


public class PreferencesBuilder implements com.grafana.foundation.cog.Builder<Preferences> {
    protected final Preferences internal;
    
    public PreferencesBuilder() {
        this.internal = new Preferences();
    }
    public PreferencesBuilder homeDashboardUID(String homeDashboardUID) {
    this.internal.homeDashboardUID = homeDashboardUID;
        return this;
    }
    
    public PreferencesBuilder timezone(String timezone) {
    this.internal.timezone = timezone;
        return this;
    }
    
    public PreferencesBuilder weekStart(String weekStart) {
    this.internal.weekStart = weekStart;
        return this;
    }
    
    public PreferencesBuilder theme(String theme) {
    this.internal.theme = theme;
        return this;
    }
    
    public PreferencesBuilder language(String language) {
    this.internal.language = language;
        return this;
    }
    
    public PreferencesBuilder queryHistory(com.grafana.foundation.cog.Builder<QueryHistoryPreference> queryHistory) {
    this.internal.queryHistory = queryHistory.build();
        return this;
    }
    
    public PreferencesBuilder cookiePreferences(com.grafana.foundation.cog.Builder<CookiePreferences> cookiePreferences) {
    this.internal.cookiePreferences = cookiePreferences.build();
        return this;
    }
    
    public PreferencesBuilder navbar(com.grafana.foundation.cog.Builder<NavbarPreference> navbar) {
    this.internal.navbar = navbar.build();
        return this;
    }
    public Preferences build() {
        return this.internal;
    }
}
