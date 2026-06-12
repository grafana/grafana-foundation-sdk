// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.playlistv1;


public class ItemBuilder implements com.grafana.foundation.cog.Builder<Item> {
    protected final Item internal;
    
    public ItemBuilder() {
        this.internal = new Item();
    }
    public ItemBuilder type(PlaylistItemType type) {
        this.internal.type = type;
        return this;
    }
    
    public ItemBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    public Item build() {
        return this.internal;
    }
}
