// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as canvas from '../canvas';

export class CanvasElementOptionsBuilder implements cog.Builder<canvas.CanvasElementOptions> {
    protected readonly internal: canvas.CanvasElementOptions;

    constructor() {
        this.internal = canvas.defaultCanvasElementOptions();
    }

    /**
     * Builds the object.
     */
    build(): canvas.CanvasElementOptions {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    type(type: string): this {
        this.internal.type = type;
        return this;
    }

    // TODO: figure out how to define this (element config(s))
    config(config: any): this {
        this.internal.config = config;
        return this;
    }

    constraint(constraint: cog.Builder<canvas.Constraint>): this {
        const constraintResource = constraint.build();
        this.internal.constraint = constraintResource;
        return this;
    }

    placement(placement: cog.Builder<canvas.Placement>): this {
        const placementResource = placement.build();
        this.internal.placement = placementResource;
        return this;
    }

    background(background: cog.Builder<canvas.BackgroundConfig>): this {
        const backgroundResource = background.build();
        this.internal.background = backgroundResource;
        return this;
    }

    border(border: cog.Builder<canvas.LineConfig>): this {
        const borderResource = border.build();
        this.internal.border = borderResource;
        return this;
    }

    connections(connections: cog.Builder<canvas.CanvasConnection>[]): this {
        const connectionsResources = connections.map(builder1 => builder1.build());
        this.internal.connections = connectionsResources;
        return this;
    }
}
