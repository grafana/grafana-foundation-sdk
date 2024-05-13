// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class RuleGroupBuilder implements cog.Builder<alerting.RuleGroup> {
    protected readonly internal: alerting.RuleGroup;

    constructor(title: string) {
        this.internal = alerting.defaultRuleGroup();
        this.internal.title = title;
    }

    build(): alerting.RuleGroup {
        return this.internal;
    }

    folderUid(folderUid: string): this {
        this.internal.folderUid = folderUid;
        return this;
    }

    // The interval, in seconds, at which all rules in the group are evaluated.
    // If a group contains many rules, the rules are evaluated sequentially.
    interval(interval: alerting.Duration): this {
        this.internal.interval = interval;
        return this;
    }

    rules(rules: cog.Builder<alerting.Rule>[]): this {
        const rulesResources = rules.map(builder1 => builder1.build());
        this.internal.rules = rulesResources;
        return this;
    }

    withRule(rules: cog.Builder<alerting.Rule>): this {
        if (!this.internal.rules) {
            this.internal.rules = [];
        }
        const rulesResource = rules.build();
        this.internal.rules.push(rulesResource);
        return this;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }
}
