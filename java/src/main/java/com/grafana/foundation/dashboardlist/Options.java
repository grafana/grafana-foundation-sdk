// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardlist;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.List;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

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
    @JsonProperty("showFolderNames")
    public Boolean showFolderNames;
    @JsonProperty("maxItems")
    public Long maxItems;
    @JsonProperty("query")
    public String query;
    @JsonProperty("tags")
    public List<String> tags;
    // folderId is deprecated, and migrated to folderUid on panel init
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("folderId")
    public Long folderId;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("folderUID")
    public String folderUID;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
