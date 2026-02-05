// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class AnnotationActionsBuilder implements com.grafana.foundation.cog.Builder<AnnotationActions> {
    protected final AnnotationActions internal;
    
    public AnnotationActionsBuilder() {
        this.internal = new AnnotationActions();
    }
    public AnnotationActionsBuilder canAdd(Boolean canAdd) {
        this.internal.canAdd = canAdd;
        return this;
    }
    
    public AnnotationActionsBuilder canDelete(Boolean canDelete) {
        this.internal.canDelete = canDelete;
        return this;
    }
    
    public AnnotationActionsBuilder canEdit(Boolean canEdit) {
        this.internal.canEdit = canEdit;
        return this;
    }
    public AnnotationActions build() {
        return this.internal;
    }
}
