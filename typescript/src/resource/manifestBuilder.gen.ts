// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as resource from '../resource';

export class ManifestBuilder implements cog.Builder<resource.Manifest> {
    protected readonly internal: resource.Manifest;

    constructor() {
        this.internal = resource.defaultManifest();
    }

    /**
     * Builds the object.
     */
    build(): resource.Manifest {
        return this.internal;
    }

    apiVersion(apiVersion: string): this {
        this.internal.apiVersion = apiVersion;
        return this;
    }

    kind(kind: string): this {
        this.internal.kind = kind;
        return this;
    }

    metadata(metadata: cog.Builder<resource.Metadata>): this {
        const metadataResource = metadata.build();
        this.internal.metadata = metadataResource;
        return this;
    }

    spec(spec: any): this {
        this.internal.spec = spec;
        return this;
    }
}

