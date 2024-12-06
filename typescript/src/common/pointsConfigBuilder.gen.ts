// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class PointsConfigBuilder implements cog.Builder<common.PointsConfig> {
    protected readonly internal: common.PointsConfig;

    constructor() {
        this.internal = common.defaultPointsConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.PointsConfig {
        return this.internal;
    }

    showPoints(showPoints: common.VisibilityMode): this {
        this.internal.showPoints = showPoints;
        return this;
    }

    pointSize(pointSize: number): this {
        this.internal.pointSize = pointSize;
        return this;
    }

    pointColor(pointColor: string): this {
        this.internal.pointColor = pointColor;
        return this;
    }

    pointSymbol(pointSymbol: string): this {
        this.internal.pointSymbol = pointSymbol;
        return this;
    }
}
