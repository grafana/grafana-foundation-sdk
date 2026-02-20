// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.folderv1beta1;


public class FolderBuilder implements com.grafana.foundation.cog.Builder<Folder> {
    protected final Folder internal;
    
    public FolderBuilder(String title) {
        this.internal = new Folder();
        this.internal.title = title;
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
