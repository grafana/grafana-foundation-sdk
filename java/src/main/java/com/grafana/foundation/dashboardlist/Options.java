// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardlist;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class Options {
    @JsonProperty("keepTime")
    public Boolean keepTime;
    @JsonProperty("includeVars")
    public Boolean includeVars;
    @JsonProperty("showStarred")
    public Boolean showStarred;
    @JsonProperty("showRecentlyViewed")
    public Boolean showRecentlyViewed;
    @JsonProperty("showSearch")
    public Boolean showSearch;
    @JsonProperty("showHeadings")
    public Boolean showHeadings;
    @JsonProperty("maxItems")
    public Long maxItems;
    @JsonProperty("query")
    public String query;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("folderId")
    public Long folderId;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("tags")
    public List<String> tags;
    public Options() {
        this.keepTime = false;
        this.includeVars = false;
        this.showStarred = true;
        this.showRecentlyViewed = false;
        this.showSearch = false;
        this.showHeadings = true;
        this.maxItems = 10L;
        this.query = "";
    }
    
    public Options(Boolean keepTime,Boolean includeVars,Boolean showStarred,Boolean showRecentlyViewed,Boolean showSearch,Boolean showHeadings,Long maxItems,String query,Long folderId,List<String> tags) {
        this.keepTime = keepTime;
        this.includeVars = includeVars;
        this.showStarred = showStarred;
        this.showRecentlyViewed = showRecentlyViewed;
        this.showSearch = showSearch;
        this.showHeadings = showHeadings;
        this.maxItems = maxItems;
        this.query = query;
        this.folderId = folderId;
        this.tags = tags;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
