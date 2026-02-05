"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.RuleGroupBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
class RuleGroupBuilder {
    constructor(title) {
        this.internal = alerting.defaultRuleGroup();
        this.internal.title = title;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    folderUid(folderUid) {
        this.internal.folderUid = folderUid;
        return this;
    }
    // The interval, in seconds, at which all rules in the group are evaluated.
    // If a group contains many rules, the rules are evaluated sequentially.
    interval(interval) {
        this.internal.interval = interval;
        return this;
    }
    rules(rules) {
        const rulesResources = rules.map(builder1 => builder1.build());
        this.internal.rules = rulesResources;
        return this;
    }
    withRule(rule) {
        if (!this.internal.rules) {
            this.internal.rules = [];
        }
        const ruleResource = rule.build();
        this.internal.rules.push(ruleResource);
        return this;
    }
    title(title) {
        this.internal.title = title;
        return this;
    }
}
exports.RuleGroupBuilder = RuleGroupBuilder;
//# sourceMappingURL=ruleGroupBuilder.gen.js.map