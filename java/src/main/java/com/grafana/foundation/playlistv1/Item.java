// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.playlistv1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Item {
    // type of the item.
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public PlaylistItemType type;
    // Value depends on type and describes the playlist item.
    //  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
    //  is not portable as the numerical identifier is non-deterministic between different instances.
    //  Will be replaced by dashboard_by_uid in the future. (deprecated)
    //  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
    //  dashboards behind the tag will be added to the playlist.
    //  - dashboard_by_uid: The value is the dashboard UID
    @JsonProperty("value")
    public String value;
    public Item() {
        this.type = PlaylistItemType.DASHBOARD_BY_TAG;
        this.value = "";
    }
    public Item(PlaylistItemType type,String value) {
        this.type = type;
        this.value = value;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
