// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class SwitchVariableKindBuilder implements cog.Builder<dashboardv2beta1.SwitchVariableKind> {
    protected readonly internal: dashboardv2beta1.SwitchVariableKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultSwitchVariableKind();
        this.internal.kind = "SwitchVariable";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.SwitchVariableKind {
        return this.internal;
    }

    spec(spec: cog.Builder<dashboardv2beta1.SwitchVariableSpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
}

