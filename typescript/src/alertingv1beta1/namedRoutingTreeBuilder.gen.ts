// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv1beta1 from '../alertingv1beta1';

export class NamedRoutingTreeBuilder implements cog.Builder<alertingv1beta1.NamedRoutingTree> {
    protected readonly internal: alertingv1beta1.NamedRoutingTree;

    constructor() {
        this.internal = alertingv1beta1.defaultNamedRoutingTree();
    }

    /**
     * Builds the object.
     */
    build(): alertingv1beta1.NamedRoutingTree {
        return this.internal;
    }

    routingTree(routingTree: string): this {
        this.internal.routingTree = routingTree;
        return this;
    }
}

