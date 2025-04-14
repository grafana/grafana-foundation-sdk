// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferences;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Spec defines user, team or org Grafana preferences
// swagger:model Preferences
public class Preferences {
    // UID for the home dashboard
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("homeDashboardUID")
    public String homeDashboardUID;
    // The timezone selection
    // TODO: this should use the timezone defined in common
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timezone")
    public String timezone;
    // day of the week (sunday, monday, etc)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("weekStart")
    public String weekStart;
    // light, dark, empty is default
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("theme")
    public String theme;
    // Selected language (beta)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("language")
    public String language;
    // Explore query history preferences
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryHistory")
    public QueryHistoryPreference queryHistory;
    // Cookie preferences
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("cookiePreferences")
    public CookiePreferences cookiePreferences;
    // Navigation preferences
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("navbar")
    public NavbarPreference navbar;
    public Preferences() {
    }
    public Preferences(String homeDashboardUID,String timezone,String weekStart,String theme,String language,QueryHistoryPreference queryHistory,CookiePreferences cookiePreferences,NavbarPreference navbar) {
        this.homeDashboardUID = homeDashboardUID;
        this.timezone = timezone;
        this.weekStart = weekStart;
        this.theme = theme;
        this.language = language;
        this.queryHistory = queryHistory;
        this.cookiePreferences = cookiePreferences;
        this.navbar = navbar;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
