// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.playlist;


public class PlaylistItemBuilder implements com.grafana.foundation.cog.Builder<PlaylistItem> {
    protected final PlaylistItem internal;
    
    public PlaylistItemBuilder() {
        this.internal = new PlaylistItem();
    }
    public PlaylistItemBuilder type(PlaylistItemType type) {
        this.internal.type = type;
        return this;
    }
    
    public PlaylistItemBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    
    public PlaylistItemBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    public PlaylistItem build() {
        return this.internal;
    }
}
