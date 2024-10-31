// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

// Links with references to other dashboards or external resources
public class DashboardLink {
    // Title to display with the link
    @JsonProperty("title")
    public String title;
    // Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public DashboardLinkType type;
    // Icon name to be displayed with the link
    @JsonProperty("icon")
    public String icon;
    // Tooltip to display when the user hovers their mouse over it
    @JsonProperty("tooltip")
    public String tooltip;
    // Link URL. Only required/valid if the type is link
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("url")
    public String url;
    // List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("tags")
    public List<String> tags;
    // If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
    @JsonProperty("asDropdown")
    public Boolean asDropdown;
    // If true, the link will be opened in a new tab
    @JsonProperty("targetBlank")
    public Boolean targetBlank;
    // If true, includes current template variables values in the link as query params
    @JsonProperty("includeVars")
    public Boolean includeVars;
    // If true, includes current time range in the link as query params
    @JsonProperty("keepTime")
    public Boolean keepTime;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DashboardLink> {
        protected final DashboardLink internal;
        
        public Builder(String title) {
            this.internal = new DashboardLink();
    this.internal.title = title;
        this.asDropdown(false);
        this.targetBlank(false);
        this.includeVars(false);
        this.keepTime(false);
        }
    public Builder title(String title) {
    this.internal.title = title;
        return this;
    }
    
    public Builder type(DashboardLinkType type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder icon(String icon) {
    this.internal.icon = icon;
        return this;
    }
    
    public Builder tooltip(String tooltip) {
    this.internal.tooltip = tooltip;
        return this;
    }
    
    public Builder url(String url) {
    this.internal.url = url;
        return this;
    }
    
    public Builder tags(List<String> tags) {
    this.internal.tags = tags;
        return this;
    }
    
    public Builder asDropdown(Boolean asDropdown) {
    this.internal.asDropdown = asDropdown;
        return this;
    }
    
    public Builder targetBlank(Boolean targetBlank) {
    this.internal.targetBlank = targetBlank;
        return this;
    }
    
    public Builder includeVars(Boolean includeVars) {
    this.internal.includeVars = includeVars;
        return this;
    }
    
    public Builder keepTime(Boolean keepTime) {
    this.internal.keepTime = keepTime;
        return this;
    }
    public DashboardLink build() {
            return this.internal;
        }
    }
}
