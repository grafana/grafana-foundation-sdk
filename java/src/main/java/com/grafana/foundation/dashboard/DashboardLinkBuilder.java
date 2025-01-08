// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;

public class DashboardLinkBuilder implements com.grafana.foundation.cog.Builder<DashboardLink> {
    protected final DashboardLink internal;
    
    public DashboardLinkBuilder(String title) {
        this.internal = new DashboardLink();
    this.internal.title = title;
    }
    public DashboardLinkBuilder title(String title) {
    this.internal.title = title;
        return this;
    }
    
    public DashboardLinkBuilder type(DashboardLinkType type) {
    this.internal.type = type;
        return this;
    }
    
    public DashboardLinkBuilder icon(String icon) {
    this.internal.icon = icon;
        return this;
    }
    
    public DashboardLinkBuilder tooltip(String tooltip) {
    this.internal.tooltip = tooltip;
        return this;
    }
    
    public DashboardLinkBuilder url(String url) {
    this.internal.url = url;
        return this;
    }
    
    public DashboardLinkBuilder tags(List<String> tags) {
    this.internal.tags = tags;
        return this;
    }
    
    public DashboardLinkBuilder asDropdown(Boolean asDropdown) {
    this.internal.asDropdown = asDropdown;
        return this;
    }
    
    public DashboardLinkBuilder targetBlank(Boolean targetBlank) {
    this.internal.targetBlank = targetBlank;
        return this;
    }
    
    public DashboardLinkBuilder includeVars(Boolean includeVars) {
    this.internal.includeVars = includeVars;
        return this;
    }
    
    public DashboardLinkBuilder keepTime(Boolean keepTime) {
    this.internal.keepTime = keepTime;
        return this;
    }
    public DashboardLink build() {
        return this.internal;
    }
}
