// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

export class NamedRoutingTreeBuilder implements cog.Builder<alertingv0alpha1.NamedRoutingTree> {
    protected readonly internal: alertingv0alpha1.NamedRoutingTree;

    constructor() {
        this.internal = alertingv0alpha1.defaultNamedRoutingTree();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.NamedRoutingTree {
        return this.internal;
    }

    routingTree(routingTree: string): this {
        this.internal.routingTree = routingTree;
        return this;
    }
}

