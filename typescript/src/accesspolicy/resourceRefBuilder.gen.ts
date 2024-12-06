// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as accesspolicy from '../accesspolicy';

export class ResourceRefBuilder implements cog.Builder<accesspolicy.ResourceRef> {
    protected readonly internal: accesspolicy.ResourceRef;

    constructor() {
        this.internal = accesspolicy.defaultResourceRef();
    }

    /**
     * Builds the object.
     */
    build(): accesspolicy.ResourceRef {
        return this.internal;
    }

    kind(kind: string): this {
        this.internal.kind = kind;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}
