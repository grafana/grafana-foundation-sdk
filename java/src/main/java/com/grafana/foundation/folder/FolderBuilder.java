// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.folder;


public class FolderBuilder implements com.grafana.foundation.cog.Builder<Folder> {
    protected final Folder internal;
    
    public FolderBuilder() {
        this.internal = new Folder();
    }
    public FolderBuilder uid(String uid) {
    this.internal.uid = uid;
        return this;
    }
    
    public FolderBuilder title(String title) {
    this.internal.title = title;
        return this;
    }
    
    public FolderBuilder description(String description) {
    this.internal.description = description;
        return this;
    }
    public Folder build() {
        return this.internal;
    }
}
