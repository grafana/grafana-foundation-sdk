// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class ElementReferenceBuilder implements cog.Builder<dashboardv2.ElementReference> {
    protected readonly internal: dashboardv2.ElementReference;

    constructor() {
        this.internal = dashboardv2.defaultElementReference();
        this.internal.kind = "ElementReference";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.ElementReference {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}

