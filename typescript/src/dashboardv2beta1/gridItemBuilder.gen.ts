// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class GridItemBuilder implements cog.Builder<dashboardv2beta1.GridLayoutItemKind> {
    protected readonly internal: dashboardv2beta1.GridLayoutItemKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultGridLayoutItemKind();
        this.internal.kind = "GridLayoutItem";
        this.internal.spec.element.kind = "ElementReference";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.GridLayoutItemKind {
        return this.internal;
    }

    x(x: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        this.internal.spec.x = x;
        return this;
    }

    y(y: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        this.internal.spec.y = y;
        return this;
    }

    width(width: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        this.internal.spec.width = width;
        return this;
    }

    height(height: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        this.internal.spec.height = height;
        return this;
    }

    repeat(repeat: cog.Builder<dashboardv2beta1.RepeatOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        if (!this.internal.spec.element) {
            this.internal.spec.element = dashboardv2beta1.defaultElementReference();
        }
        this.internal.spec.element.name = name;
        return this;
    }
}

export function gridItem(name: string): GridItemBuilder {
	const builder = new GridItemBuilder();
	builder.name(name);

	return builder;
}

