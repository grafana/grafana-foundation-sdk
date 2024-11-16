// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as xychart from '../xychart';

export class XYSeriesConfigBuilder implements cog.Builder<xychart.XYSeriesConfig> {
    protected readonly internal: xychart.XYSeriesConfig;

    constructor() {
        this.internal = xychart.defaultXYSeriesConfig();
    }

    /**
     * Builds the object.
     */
    build(): xychart.XYSeriesConfig {
        return this.internal;
    }

    name(name: {
	fixed?: string;
}): this {
        this.internal.name = name;
        return this;
    }

    frame(frame: {
	matcher: xychart.MatcherConfig;
}): this {
        this.internal.frame = frame;
        return this;
    }

    x(x: {
	matcher: xychart.MatcherConfig;
}): this {
        this.internal.x = x;
        return this;
    }

    y(y: {
	matcher: xychart.MatcherConfig;
}): this {
        this.internal.y = y;
        return this;
    }

    color(color: {
	matcher: xychart.MatcherConfig;
}): this {
        this.internal.color = color;
        return this;
    }

    size(size: {
	matcher: xychart.MatcherConfig;
}): this {
        this.internal.size = size;
        return this;
    }
}
