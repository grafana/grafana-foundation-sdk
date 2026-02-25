// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.folderv1beta1;

import com.grafana.foundation.resource.ManifestBuilder;
import com.grafana.foundation.resource.Resource;
public class Folderv1beta1 {    
    /**
     * Creates a resource manifest from a Folder.
     */
    public static ManifestBuilder manifest(String name,com.grafana.foundation.cog.Builder<Folder> folder) {
        ManifestBuilder builder = new ManifestBuilder();
        builder.apiVersion("folder.grafana.app/v1beta1");
        builder.kind("Folder");
        builder.metadata(Resource.named(name));
        builder.spec(folder.build());
        return builder;
    }
}
