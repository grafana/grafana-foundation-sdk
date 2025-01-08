// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;


public class LibraryElementDTOMetaBuilder implements com.grafana.foundation.cog.Builder<LibraryElementDTOMeta> {
    protected final LibraryElementDTOMeta internal;
    
    public LibraryElementDTOMetaBuilder() {
        this.internal = new LibraryElementDTOMeta();
    }
    public LibraryElementDTOMetaBuilder folderName(String folderName) {
    this.internal.folderName = folderName;
        return this;
    }
    
    public LibraryElementDTOMetaBuilder folderUid(String folderUid) {
    this.internal.folderUid = folderUid;
        return this;
    }
    
    public LibraryElementDTOMetaBuilder connectedDashboards(Long connectedDashboards) {
    this.internal.connectedDashboards = connectedDashboards;
        return this;
    }
    
    public LibraryElementDTOMetaBuilder created(String created) {
    this.internal.created = created;
        return this;
    }
    
    public LibraryElementDTOMetaBuilder updated(String updated) {
    this.internal.updated = updated;
        return this;
    }
    
    public LibraryElementDTOMetaBuilder createdBy(com.grafana.foundation.cog.Builder<LibraryElementDTOMetaUser> createdBy) {
    this.internal.createdBy = createdBy.build();
        return this;
    }
    
    public LibraryElementDTOMetaBuilder updatedBy(com.grafana.foundation.cog.Builder<LibraryElementDTOMetaUser> updatedBy) {
    this.internal.updatedBy = updatedBy.build();
        return this;
    }
    public LibraryElementDTOMeta build() {
        return this.internal;
    }
}
