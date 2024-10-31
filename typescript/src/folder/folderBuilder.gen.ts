// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as folder from '../folder';

// TODO:
// common metadata will soon support setting the parent folder in the metadata
export class FolderBuilder implements cog.Builder<folder.Folder> {
    protected readonly internal: folder.Folder;

    constructor() {
        this.internal = folder.defaultFolder();
    }

    /**
     * Builds the object.
     */
    build(): folder.Folder {
        return this.internal;
    }

    // Unique folder id. (will be k8s name)
    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }

    // Folder title
    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    // Description of the folder.
    description(description: string): this {
        this.internal.description = description;
        return this;
    }
}
