// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.annotationslist;

import java.util.List;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder onlyFromThisDashboard(Boolean onlyFromThisDashboard) {
        this.internal.onlyFromThisDashboard = onlyFromThisDashboard;
        return this;
    }
    
    public OptionsBuilder onlyInTimeRange(Boolean onlyInTimeRange) {
        this.internal.onlyInTimeRange = onlyInTimeRange;
        return this;
    }
    
    public OptionsBuilder tags(List<String> tags) {
        this.internal.tags = tags;
        return this;
    }
    
    public OptionsBuilder limit(Integer limit) {
        this.internal.limit = limit;
        return this;
    }
    
    public OptionsBuilder showUser(Boolean showUser) {
        this.internal.showUser = showUser;
        return this;
    }
    
    public OptionsBuilder showTime(Boolean showTime) {
        this.internal.showTime = showTime;
        return this;
    }
    
    public OptionsBuilder showTags(Boolean showTags) {
        this.internal.showTags = showTags;
        return this;
    }
    
    public OptionsBuilder navigateToPanel(Boolean navigateToPanel) {
        this.internal.navigateToPanel = navigateToPanel;
        return this;
    }
    
    public OptionsBuilder navigateBefore(String navigateBefore) {
        this.internal.navigateBefore = navigateBefore;
        return this;
    }
    
    public OptionsBuilder navigateAfter(String navigateAfter) {
        this.internal.navigateAfter = navigateAfter;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
