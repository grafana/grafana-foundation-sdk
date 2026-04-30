// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class AutoGridItemBuilder implements cog.Builder<dashboardv2.AutoGridLayoutItemKind> {
    protected readonly internal: dashboardv2.AutoGridLayoutItemKind;

    constructor() {
        this.internal = dashboardv2.defaultAutoGridLayoutItemKind();
        this.internal.kind = "AutoGridLayoutItem";
        this.internal.spec.element.kind = "ElementReference";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.AutoGridLayoutItemKind {
        return this.internal;
    }

    repeat(repeat: cog.Builder<dashboardv2.AutoGridRepeatOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultAutoGridLayoutItemSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }

    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2.ConditionalRenderingGroupKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultAutoGridLayoutItemSpec();
        }
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultAutoGridLayoutItemSpec();
        }
        if (!this.internal.spec.element) {
            this.internal.spec.element = dashboardv2.defaultElementReference();
        }
        this.internal.spec.element.name = name;
        return this;
    }
}

export function autoGridItem(name: string): AutoGridItemBuilder {
	const builder = new AutoGridItemBuilder();
	builder.name(name);

	return builder;
}

