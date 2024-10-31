// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as canvas from '../canvas';
import * as common from '../common';

export class CanvasConnectionBuilder implements cog.Builder<canvas.CanvasConnection> {
    protected readonly internal: canvas.CanvasConnection;

    constructor() {
        this.internal = canvas.defaultCanvasConnection();
    }

    /**
     * Builds the object.
     */
    build(): canvas.CanvasConnection {
        return this.internal;
    }

    source(source: cog.Builder<canvas.ConnectionCoordinates>): this {
        const sourceResource = source.build();
        this.internal.source = sourceResource;
        return this;
    }

    target(target: cog.Builder<canvas.ConnectionCoordinates>): this {
        const targetResource = target.build();
        this.internal.target = targetResource;
        return this;
    }

    targetName(targetName: string): this {
        this.internal.targetName = targetName;
        return this;
    }

    path(path: canvas.ConnectionPath): this {
        this.internal.path = path;
        return this;
    }

    color(color: cog.Builder<common.ColorDimensionConfig>): this {
        const colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }

    size(size: cog.Builder<common.ScaleDimensionConfig>): this {
        const sizeResource = size.build();
        this.internal.size = sizeResource;
        return this;
    }

    vertices(vertices: cog.Builder<canvas.ConnectionCoordinates>[]): this {
        const verticesResources = vertices.map(builder1 => builder1.build());
        this.internal.vertices = verticesResources;
        return this;
    }

    sourceOriginal(sourceOriginal: cog.Builder<canvas.ConnectionCoordinates>): this {
        const sourceOriginalResource = sourceOriginal.build();
        this.internal.sourceOriginal = sourceOriginalResource;
        return this;
    }

    targetOriginal(targetOriginal: cog.Builder<canvas.ConnectionCoordinates>): this {
        const targetOriginalResource = targetOriginal.build();
        this.internal.targetOriginal = targetOriginalResource;
        return this;
    }
}
