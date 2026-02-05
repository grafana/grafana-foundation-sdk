"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultDataquery = exports.defaultUSAQuery = exports.defaultTimeRange = exports.defaultStreamingQuery = exports.defaultSimulationQuery = exports.defaultKey = exports.defaultResultAssertions = exports.defaultPulseWaveQuery = exports.defaultNodesQuery = exports.defaultCSVWave = void 0;
const defaultCSVWave = () => ({});
exports.defaultCSVWave = defaultCSVWave;
const defaultNodesQuery = () => ({});
exports.defaultNodesQuery = defaultNodesQuery;
const defaultPulseWaveQuery = () => ({});
exports.defaultPulseWaveQuery = defaultPulseWaveQuery;
const defaultResultAssertions = () => ({
    typeVersion: [],
});
exports.defaultResultAssertions = defaultResultAssertions;
const defaultKey = () => ({
    tick: 0,
    type: "",
});
exports.defaultKey = defaultKey;
const defaultSimulationQuery = () => ({
    key: (0, exports.defaultKey)(),
});
exports.defaultSimulationQuery = defaultSimulationQuery;
const defaultStreamingQuery = () => ({
    noise: 0,
    speed: 0,
    spread: 0,
    type: "fetch",
});
exports.defaultStreamingQuery = defaultStreamingQuery;
const defaultTimeRange = () => ({
    from: "now-6h",
    to: "now",
});
exports.defaultTimeRange = defaultTimeRange;
const defaultUSAQuery = () => ({});
exports.defaultUSAQuery = defaultUSAQuery;
const defaultDataquery = () => ({
    _implementsDataqueryVariant: () => { },
});
exports.defaultDataquery = defaultDataquery;
//# sourceMappingURL=types.gen.js.map