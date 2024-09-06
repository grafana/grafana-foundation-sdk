// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardlist;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonInclude;

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
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("tags")
    public List<String> tags;
    // folderId is deprecated, and migrated to folderUid on panel init
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("folderId")
    public Long folderId;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("folderUID")
    public String folderUID;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
