// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferencesv1alpha1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Preferences {
    // Explicit home URL (NOTE: this can only be modified in the system settings)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("homeURL")
    public String homeURL;
    // UID for the home dashboard
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("homeDashboardUID")
    public String homeDashboardUID;
    // The timezone selection
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timezone")
    public String timezone;
    // day of the week (sunday, monday, etc)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("weekStart")
    public String weekStart;
    // user interface theme
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("theme")
    public String theme;
    // Selected language
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("language")
    public String language;
    // Explore query history preferences
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryHistory")
    public QueryHistoryPreference queryHistory;
    // Navigation preferences
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("navbar")
    public NavbarPreference navbar;
    public Preferences() {
    }
    public Preferences(String homeURL,String homeDashboardUID,String timezone,String weekStart,String theme,String language,QueryHistoryPreference queryHistory,NavbarPreference navbar) {
        this.homeURL = homeURL;
        this.homeDashboardUID = homeDashboardUID;
        this.timezone = timezone;
        this.weekStart = weekStart;
        this.theme = theme;
        this.language = language;
        this.queryHistory = queryHistory;
        this.navbar = navbar;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
