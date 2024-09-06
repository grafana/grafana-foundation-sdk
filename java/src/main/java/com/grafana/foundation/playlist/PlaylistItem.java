// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.playlist;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class PlaylistItem {
    // Type of the item.
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public PlaylistItemType type;
    // Value depends on type and describes the playlist item.
    // 
    //  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
    //  is not portable as the numerical identifier is non-deterministic between different instances.
    //  Will be replaced by dashboard_by_uid in the future. (deprecated)
    //  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
    //  dashboards behind the tag will be added to the playlist.
    //  - dashboard_by_uid: The value is the dashboard UID
    @JsonProperty("value")
    public String value;
    // Title is an unused property -- it will be removed in the future
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<PlaylistItem> {
        private final PlaylistItem internal;
        
        public Builder() {
            this.internal = new PlaylistItem();
        }
    public Builder type(PlaylistItemType type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder value(String value) {
    this.internal.value = value;
        return this;
    }
    
    public Builder title(String title) {
    this.internal.title = title;
        return this;
    }
    public PlaylistItem build() {
            return this.internal;
        }
    }
}
