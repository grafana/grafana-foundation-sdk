// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as starsv1alpha1 from '../starsv1alpha1';

export class ResourceBuilder implements cog.Builder<starsv1alpha1.Resource> {
    protected readonly internal: starsv1alpha1.Resource;

    constructor() {
        this.internal = starsv1alpha1.defaultResource();
    }

    /**
     * Builds the object.
     */
    build(): starsv1alpha1.Resource {
        return this.internal;
    }

    group(group: string): this {
        this.internal.group = group;
        return this;
    }

    kind(kind: string): this {
        this.internal.kind = kind;
        return this;
    }

    // The set of resources
    // +listType=set
    names(names: string[]): this {
        this.internal.names = names;
        return this;
    }
}

