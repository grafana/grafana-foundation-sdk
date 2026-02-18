// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.playlistv0alpha1;

import java.util.List;
import java.util.LinkedList;

public class PlaylistBuilder implements com.grafana.foundation.cog.Builder<Playlist> {
    protected final Playlist internal;
    
    public PlaylistBuilder(String title) {
        this.internal = new Playlist();
        this.internal.title = title;
    }
    public PlaylistBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public PlaylistBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }
    
    public PlaylistBuilder items(List<com.grafana.foundation.cog.Builder<Item>> items) {
        List<Item> itemsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Item> r1 : items) {
                Item itemsDepth1 = r1.build();
                itemsResources.add(itemsDepth1); 
        }
        this.internal.items = itemsResources;
        return this;
    }
    
    public PlaylistBuilder item(com.grafana.foundation.cog.Builder<Item> item) {
		if (this.internal.items == null) {
			this.internal.items = new LinkedList<>();
		}
    Item itemResource = item.build();
        this.internal.items.add(itemResource);
        return this;
    }
    public Playlist build() {
        return this.internal;
    }
}
