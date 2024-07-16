// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferences;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Preferences {
    // UID for the home dashboard 
    @JsonProperty("homeDashboardUID")
    public String homeDashboardUID;
    // The timezone selection
    // TODO: this should use the timezone defined in common 
    @JsonProperty("timezone")
    public String timezone;
    // day of the week (sunday, monday, etc) 
    @JsonProperty("weekStart")
    public String weekStart;
    // light, dark, empty is default 
    @JsonProperty("theme")
    public String theme;
    // Selected language (beta) 
    @JsonProperty("language")
    public String language;
    // Explore query history preferences 
    @JsonProperty("queryHistory")
    public QueryHistoryPreference queryHistory;
    // Cookie preferences 
    @JsonProperty("cookiePreferences")
    public CookiePreferences cookiePreferences;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Preferences> {
        private final Preferences internal;
        
        public Builder() {
            this.internal = new Preferences();
        }
    public Builder homeDashboardUID(String homeDashboardUID) {
    this.internal.homeDashboardUID = homeDashboardUID;
        return this;
    }
    
    public Builder timezone(String timezone) {
    this.internal.timezone = timezone;
        return this;
    }
    
    public Builder weekStart(String weekStart) {
    this.internal.weekStart = weekStart;
        return this;
    }
    
    public Builder theme(String theme) {
    this.internal.theme = theme;
        return this;
    }
    
    public Builder language(String language) {
    this.internal.language = language;
        return this;
    }
    
    public Builder queryHistory(com.grafana.foundation.cog.Builder<QueryHistoryPreference> queryHistory) {
    this.internal.queryHistory = queryHistory.build();
        return this;
    }
    
    public Builder cookiePreferences(com.grafana.foundation.cog.Builder<CookiePreferences> cookiePreferences) {
    this.internal.cookiePreferences = cookiePreferences.build();
        return this;
    }
    public Preferences build() {
            return this.internal;
        }
    }
}
