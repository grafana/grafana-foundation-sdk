// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as accesspolicy from '../accesspolicy';

export class AccessRuleBuilder implements cog.Builder<accesspolicy.AccessRule> {
    protected readonly internal: accesspolicy.AccessRule;

    constructor() {
        this.internal = accesspolicy.defaultAccessRule();
    }

    /**
     * Builds the object.
     */
    build(): accesspolicy.AccessRule {
        return this.internal;
    }

    // The kind this rule applies to (dashboards, alert, etc)
    kind(kind: string): this {
        this.internal.kind = kind;
        return this;
    }

    // READ, WRITE, CREATE, DELETE, ...
    // should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
    verb(verb: "*" | "none" | string): this {
        this.internal.verb = verb;
        return this;
    }

    // Specific sub-elements like "alert.rules" or "dashboard.permissions"????
    target(target: string): this {
        this.internal.target = target;
        return this;
    }
}
