// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;


public class LibraryPanelBuilder implements com.grafana.foundation.cog.Builder<LibraryPanel> {
    protected final LibraryPanel internal;
    
    public LibraryPanelBuilder() {
        this.internal = new LibraryPanel();
    }
    public LibraryPanelBuilder folderUid(String folderUid) {
        this.internal.folderUid = folderUid;
        return this;
    }
    
    public LibraryPanelBuilder uid(String uid) {
        this.internal.uid = uid;
        return this;
    }
    
    public LibraryPanelBuilder name(String name) {
        if (!(name.length() >= 1)) {
            throw new IllegalArgumentException("name.length() must be >= 1");
        }
        this.internal.name = name;
        return this;
    }
    
    public LibraryPanelBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public LibraryPanelBuilder type(String type) {
        if (!(type.length() >= 1)) {
            throw new IllegalArgumentException("type.length() must be >= 1");
        }
        this.internal.type = type;
        return this;
    }
    
    public LibraryPanelBuilder schemaVersion(Short schemaVersion) {
        this.internal.schemaVersion = schemaVersion;
        return this;
    }
    
    public LibraryPanelBuilder version(Long version) {
        this.internal.version = version;
        return this;
    }
    
    public LibraryPanelBuilder model(com.grafana.foundation.cog.Builder<PanelModel> model) {
    PanelModel modelResource = model.build();
        this.internal.model = modelResource;
        return this;
    }
    
    public LibraryPanelBuilder meta(com.grafana.foundation.cog.Builder<LibraryElementDTOMeta> meta) {
    LibraryElementDTOMeta metaResource = meta.build();
        this.internal.meta = metaResource;
        return this;
    }
    public LibraryPanel build() {
        return this.internal;
    }
}
