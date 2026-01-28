// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardlist;

import java.util.List;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder keepTime(Boolean keepTime) {
        this.internal.keepTime = keepTime;
        return this;
    }
    
    public OptionsBuilder includeVars(Boolean includeVars) {
        this.internal.includeVars = includeVars;
        return this;
    }
    
    public OptionsBuilder showStarred(Boolean showStarred) {
        this.internal.showStarred = showStarred;
        return this;
    }
    
    public OptionsBuilder showRecentlyViewed(Boolean showRecentlyViewed) {
        this.internal.showRecentlyViewed = showRecentlyViewed;
        return this;
    }
    
    public OptionsBuilder showSearch(Boolean showSearch) {
        this.internal.showSearch = showSearch;
        return this;
    }
    
    public OptionsBuilder showHeadings(Boolean showHeadings) {
        this.internal.showHeadings = showHeadings;
        return this;
    }
    
    public OptionsBuilder maxItems(Long maxItems) {
        this.internal.maxItems = maxItems;
        return this;
    }
    
    public OptionsBuilder query(String query) {
        this.internal.query = query;
        return this;
    }
    
    public OptionsBuilder tags(List<String> tags) {
        this.internal.tags = tags;
        return this;
    }
    
    public OptionsBuilder folderId(Long folderId) {
        this.internal.folderId = folderId;
        return this;
    }
    
    public OptionsBuilder folderUID(String folderUID) {
        this.internal.folderUID = folderUID;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
