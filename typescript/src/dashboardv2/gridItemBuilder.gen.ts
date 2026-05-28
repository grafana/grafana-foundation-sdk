// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class GridItemBuilder implements cog.Builder<dashboardv2.GridLayoutItemKind> {
    protected readonly internal: dashboardv2.GridLayoutItemKind;

    constructor() {
        this.internal = dashboardv2.defaultGridLayoutItemKind();
        this.internal.kind = "GridLayoutItem";
        this.internal.spec.element.kind = "ElementReference";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.GridLayoutItemKind {
        return this.internal;
    }

    x(x: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGridLayoutItemSpec();
        }
        this.internal.spec.x = x;
        return this;
    }

    y(y: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGridLayoutItemSpec();
        }
        this.internal.spec.y = y;
        return this;
    }

    width(width: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGridLayoutItemSpec();
        }
        this.internal.spec.width = width;
        return this;
    }

    height(height: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGridLayoutItemSpec();
        }
        this.internal.spec.height = height;
        return this;
    }

    repeat(repeat: cog.Builder<dashboardv2.RepeatOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGridLayoutItemSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGridLayoutItemSpec();
        }
        if (!this.internal.spec.element) {
            this.internal.spec.element = dashboardv2.defaultElementReference();
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

