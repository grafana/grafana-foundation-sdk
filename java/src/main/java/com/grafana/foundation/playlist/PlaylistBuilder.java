// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.playlist;

import java.util.List;

public class PlaylistBuilder implements com.grafana.foundation.cog.Builder<Playlist> {
    protected final Playlist internal;
    
    public PlaylistBuilder() {
        this.internal = new Playlist();
    }
    public PlaylistBuilder uid(String uid) {
    this.internal.uid = uid;
        return this;
    }
    
    public PlaylistBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public PlaylistBuilder interval(String interval) {
    this.internal.interval = interval;
        return this;
    }
    
    public PlaylistBuilder items(com.grafana.foundation.cog.Builder<List<PlaylistItem>> items) {
    this.internal.items = items.build();
        return this;
    }
    public Playlist build() {
        return this.internal;
    }
}
