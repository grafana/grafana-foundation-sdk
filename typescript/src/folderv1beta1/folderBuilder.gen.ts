// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as folderv1beta1 from '../folderv1beta1';
import * as resource from '../resource';

export class FolderBuilder implements cog.Builder<folderv1beta1.Folder> {
    protected readonly internal: folderv1beta1.Folder;

    constructor(title: string) {
        this.internal = folderv1beta1.defaultFolder();
        this.internal.title = title;
    }

    /**
     * Builds the object.
     */
    build(): folderv1beta1.Folder {
        return this.internal;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    description(description: string): this {
        this.internal.description = description;
        return this;
    }
}

/**
 * Creates a resource manifest from a Folder.
 */
export function manifest(name: string,folder: cog.Builder<folderv1beta1.Folder>): cog.Builder<resource.Manifest> {
	const builder = new resource.ManifestBuilder();
	builder.apiVersion("folder.grafana.app/v1beta1");
	builder.kind("Folder");
	builder.metadata(resource.named(name));
	builder.spec(folder.build());

	return builder;
}

