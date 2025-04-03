// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as accesspolicy from '../accesspolicy';

export class AccessPolicyBuilder implements cog.Builder<accesspolicy.AccessPolicy> {
    protected readonly internal: accesspolicy.AccessPolicy;

    constructor() {
        this.internal = accesspolicy.defaultAccessPolicy();
    }

    /**
     * Builds the object.
     */
    build(): accesspolicy.AccessPolicy {
        return this.internal;
    }

    // The scope where these policies should apply
    scope(scope: cog.Builder<accesspolicy.ResourceRef>): this {
        const scopeResource = scope.build();
        this.internal.scope = scopeResource;
        return this;
    }

    // The role that must apply this policy
    role(role: cog.Builder<accesspolicy.RoleRef>): this {
        const roleResource = role.build();
        this.internal.role = roleResource;
        return this;
    }

    // The set of rules to apply.  Note that * is required to modify
    // access policy rules, and that "none" will reject all actions
    rules(rule: cog.Builder<accesspolicy.AccessRule>): this {
        if (!this.internal.rules) {
            this.internal.rules = [];
        }
        const ruleResource = rule.build();
        this.internal.rules.push(ruleResource);
        return this;
    }
}

