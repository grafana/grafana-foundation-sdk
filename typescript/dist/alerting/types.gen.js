"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultLocation = exports.defaultMonthRange = exports.defaultYearRange = exports.defaultDayOfMonthRange = exports.defaultWeekdayRange = exports.defaultTimeRange = exports.defaultTimeInterval = exports.defaultNotificationPolicy = exports.defaultRelativeTimeRange = exports.defaultRecordRule = exports.defaultRule = exports.defaultProvenance = exports.defaultObjectMatchers = exports.defaultObjectMatcher = exports.defaultNotificationTemplate = exports.defaultMuteTiming = exports.defaultMatchers = exports.defaultMatcher = exports.defaultMatchType = exports.MatchType = exports.defaultMatchRegexps = exports.defaultJson = exports.defaultContactPoint = exports.defaultDuration = exports.defaultNotificationSettings = exports.defaultRuleGroup = exports.defaultQuery = void 0;
const defaultQuery = () => ({});
exports.defaultQuery = defaultQuery;
const defaultRuleGroup = () => ({});
exports.defaultRuleGroup = defaultRuleGroup;
const defaultNotificationSettings = () => ({
    group_by: [
        "alertname",
        "grafana_folder",
    ],
    receiver: "",
});
exports.defaultNotificationSettings = defaultNotificationSettings;
const defaultDuration = () => (0);
exports.defaultDuration = defaultDuration;
const defaultContactPoint = () => ({
    settings: (0, exports.defaultJson)(),
    type: "alertmanager",
});
exports.defaultContactPoint = defaultContactPoint;
const defaultJson = () => ({});
exports.defaultJson = defaultJson;
const defaultMatchRegexps = () => ({});
exports.defaultMatchRegexps = defaultMatchRegexps;
var MatchType;
(function (MatchType) {
    MatchType["Equal"] = "=";
    MatchType["NotEqual"] = "!=";
    MatchType["EqualRegex"] = "=~";
    MatchType["NotEqualRegex"] = "!~";
})(MatchType || (exports.MatchType = MatchType = {}));
const defaultMatchType = () => (MatchType.Equal);
exports.defaultMatchType = defaultMatchType;
const defaultMatcher = () => ({});
exports.defaultMatcher = defaultMatcher;
const defaultMatchers = () => ([]);
exports.defaultMatchers = defaultMatchers;
const defaultMuteTiming = () => ({});
exports.defaultMuteTiming = defaultMuteTiming;
const defaultNotificationTemplate = () => ({});
exports.defaultNotificationTemplate = defaultNotificationTemplate;
const defaultObjectMatcher = () => ([]);
exports.defaultObjectMatcher = defaultObjectMatcher;
const defaultObjectMatchers = () => ([]);
exports.defaultObjectMatchers = defaultObjectMatchers;
const defaultProvenance = () => ("");
exports.defaultProvenance = defaultProvenance;
const defaultRule = () => ({
    condition: "",
    data: [],
    execErrState: "OK",
    folderUID: "",
    for: "",
    noDataState: "Alerting",
    orgID: 0,
    ruleGroup: "",
    title: "",
});
exports.defaultRule = defaultRule;
const defaultRecordRule = () => ({
    from: "",
    metric: "",
});
exports.defaultRecordRule = defaultRecordRule;
const defaultRelativeTimeRange = () => ({});
exports.defaultRelativeTimeRange = defaultRelativeTimeRange;
const defaultNotificationPolicy = () => ({});
exports.defaultNotificationPolicy = defaultNotificationPolicy;
const defaultTimeInterval = () => ({});
exports.defaultTimeInterval = defaultTimeInterval;
const defaultTimeRange = () => ({});
exports.defaultTimeRange = defaultTimeRange;
const defaultWeekdayRange = () => ({});
exports.defaultWeekdayRange = defaultWeekdayRange;
const defaultDayOfMonthRange = () => ({});
exports.defaultDayOfMonthRange = defaultDayOfMonthRange;
const defaultYearRange = () => ({});
exports.defaultYearRange = defaultYearRange;
const defaultMonthRange = () => ({});
exports.defaultMonthRange = defaultMonthRange;
const defaultLocation = () => ("");
exports.defaultLocation = defaultLocation;
//# sourceMappingURL=types.gen.js.map