// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
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
    // FIXME: The type is generated as `type: DashboardLinkType | dashboardLinkType.Link;` but it should be `type: DashboardLinkType`
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
    // Placement can be used to display the link somewhere else on the dashboard other than above the visualisations.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("placement")
    public String placement;
    public DashboardLink() {
        this.title = "";
        this.type = DashboardLinkType.LINK;
        this.icon = "";
        this.tooltip = "";
        this.tags = new LinkedList<>();
        this.asDropdown = false;
        this.targetBlank = false;
        this.includeVars = false;
        this.keepTime = false;
        this.placement = Constants.DashboardLinkPlacement;
    }
    public DashboardLink(String title,DashboardLinkType type,String icon,String tooltip,String url,List<String> tags,Boolean asDropdown,Boolean targetBlank,Boolean includeVars,Boolean keepTime) {
        this.title = title;
        this.type = type;
        this.icon = icon;
        this.tooltip = tooltip;
        this.url = url;
        this.tags = tags;
        this.asDropdown = asDropdown;
        this.targetBlank = targetBlank;
        this.includeVars = includeVars;
        this.keepTime = keepTime;
        this.placement = Constants.DashboardLinkPlacement;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
