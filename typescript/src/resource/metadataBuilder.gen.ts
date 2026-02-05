// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as resource from '../resource';

export class MetadataBuilder implements cog.Builder<resource.Metadata> {
    protected readonly internal: resource.Metadata;

    constructor() {
        this.internal = resource.defaultMetadata();
    }

    /**
     * Builds the object.
     */
    build(): resource.Metadata {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    namespace(namespace: string): this {
        this.internal.namespace = namespace;
        return this;
    }

    labels(labels: Record<string, string>): this {
        this.internal.labels = labels;
        return this;
    }

    annotations(annotations: Record<string, string>): this {
        this.internal.annotations = annotations;
        return this;
    }

    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }

    resourceVersion(resourceVersion: string): this {
        this.internal.resourceVersion = resourceVersion;
        return this;
    }

    generation(generation: number): this {
        this.internal.generation = generation;
        return this;
    }

    creationTimestamp(creationTimestamp: string): this {
        this.internal.creationTimestamp = creationTimestamp;
        return this;
    }

    updateTimestamp(updateTimestamp: string): this {
        this.internal.updateTimestamp = updateTimestamp;
        return this;
    }

    deletionTimestamp(deletionTimestamp: string): this {
        this.internal.deletionTimestamp = deletionTimestamp;
        return this;
    }
}

