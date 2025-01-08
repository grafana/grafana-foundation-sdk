// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.playlist;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class Playlist {
    // Unique playlist identifier. Generated on creation, either by the
    // creator of the playlist of by the application.
    @JsonProperty("uid")
    public String uid;
    // Name of the playlist.
    @JsonProperty("name")
    public String name;
    // Interval sets the time between switching views in a playlist.
    // FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
    @JsonProperty("interval")
    public String interval;
    // The ordered list of items that the playlist will iterate over.
    // FIXME! This should not be optional, but changing it makes the godegen awkward
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("items")
    public List<PlaylistItem> items;
    public Playlist() {
        this.interval = "5m";
    }
    
    public Playlist(String uid,String name,String interval,List<PlaylistItem> items) {
        this.uid = uid;
        this.name = name;
        this.interval = interval;
        this.items = items;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
