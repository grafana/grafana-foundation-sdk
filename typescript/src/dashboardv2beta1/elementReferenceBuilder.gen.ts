// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class ElementReferenceBuilder implements cog.Builder<dashboardv2beta1.ElementReference> {
    protected readonly internal: dashboardv2beta1.ElementReference;

    constructor() {
        this.internal = dashboardv2beta1.defaultElementReference();
        this.internal.kind = "ElementReference";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ElementReference {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}

