// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.playlist;

import java.util.List;
import java.util.LinkedList;

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
    
    public PlaylistBuilder items(List<com.grafana.foundation.cog.Builder<PlaylistItem>> items) {
        List<PlaylistItem> itemsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<PlaylistItem> r1 : items) {
                PlaylistItem itemsDepth1 = r1.build();
                itemsResources.add(itemsDepth1); 
        }
        this.internal.items = itemsResources;
        return this;
    }
    public Playlist build() {
        return this.internal;
    }
}
