<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated
 */
final class QueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, QueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function monitor(): self
    {
        if (!isset(self::$instances["Monitor"])) {
            self::$instances["Monitor"] = new self("Azure Monitor");
        }

        return self::$instances["Monitor"];
    }

    public static function logAnalytics(): self
    {
        if (!isset(self::$instances["LogAnalytics"])) {
            self::$instances["LogAnalytics"] = new self("Azure Log Analytics");
        }

        return self::$instances["LogAnalytics"];
    }

    public static function resourceGraph(): self
    {
        if (!isset(self::$instances["ResourceGraph"])) {
            self::$instances["ResourceGraph"] = new self("Azure Resource Graph");
        }

        return self::$instances["ResourceGraph"];
    }

    public static function traces(): self
    {
        if (!isset(self::$instances["Traces"])) {
            self::$instances["Traces"] = new self("Azure Traces");
        }

        return self::$instances["Traces"];
    }

    public static function subscriptionsQuery(): self
    {
        if (!isset(self::$instances["SubscriptionsQuery"])) {
            self::$instances["SubscriptionsQuery"] = new self("Azure Subscriptions");
        }

        return self::$instances["SubscriptionsQuery"];
    }

    public static function resourceGroupsQuery(): self
    {
        if (!isset(self::$instances["ResourceGroupsQuery"])) {
            self::$instances["ResourceGroupsQuery"] = new self("Azure Resource Groups");
        }

        return self::$instances["ResourceGroupsQuery"];
    }

    public static function namespacesQuery(): self
    {
        if (!isset(self::$instances["NamespacesQuery"])) {
            self::$instances["NamespacesQuery"] = new self("Azure Namespaces");
        }

        return self::$instances["NamespacesQuery"];
    }

    public static function resourceNamesQuery(): self
    {
        if (!isset(self::$instances["ResourceNamesQuery"])) {
            self::$instances["ResourceNamesQuery"] = new self("Azure Resource Names");
        }

        return self::$instances["ResourceNamesQuery"];
    }

    public static function metricNamesQuery(): self
    {
        if (!isset(self::$instances["MetricNamesQuery"])) {
            self::$instances["MetricNamesQuery"] = new self("Azure Metric Names");
        }

        return self::$instances["MetricNamesQuery"];
    }

    public static function workspacesQuery(): self
    {
        if (!isset(self::$instances["WorkspacesQuery"])) {
            self::$instances["WorkspacesQuery"] = new self("Azure Workspaces");
        }

        return self::$instances["WorkspacesQuery"];
    }

    public static function locationsQuery(): self
    {
        if (!isset(self::$instances["LocationsQuery"])) {
            self::$instances["LocationsQuery"] = new self("Azure Regions");
        }

        return self::$instances["LocationsQuery"];
    }

    public static function grafanaTemplateVariableFn(): self
    {
        if (!isset(self::$instances["GrafanaTemplateVariableFn"])) {
            self::$instances["GrafanaTemplateVariableFn"] = new self("Grafana Template Variable Function");
        }

        return self::$instances["GrafanaTemplateVariableFn"];
    }

    public static function traceExemplar(): self
    {
        if (!isset(self::$instances["TraceExemplar"])) {
            self::$instances["TraceExemplar"] = new self("traceql");
        }

        return self::$instances["TraceExemplar"];
    }

    public static function customNamespacesQuery(): self
    {
        if (!isset(self::$instances["CustomNamespacesQuery"])) {
            self::$instances["CustomNamespacesQuery"] = new self("Azure Custom Namespaces");
        }

        return self::$instances["CustomNamespacesQuery"];
    }

    public static function customMetricNamesQuery(): self
    {
        if (!isset(self::$instances["CustomMetricNamesQuery"])) {
            self::$instances["CustomMetricNamesQuery"] = new self("Azure Custom Metric Names");
        }

        return self::$instances["CustomMetricNamesQuery"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "Azure Monitor") {
            return self::monitor();
        }

        if ($value === "Azure Log Analytics") {
            return self::logAnalytics();
        }

        if ($value === "Azure Resource Graph") {
            return self::resourceGraph();
        }

        if ($value === "Azure Traces") {
            return self::traces();
        }

        if ($value === "Azure Subscriptions") {
            return self::subscriptionsQuery();
        }

        if ($value === "Azure Resource Groups") {
            return self::resourceGroupsQuery();
        }

        if ($value === "Azure Namespaces") {
            return self::namespacesQuery();
        }

        if ($value === "Azure Resource Names") {
            return self::resourceNamesQuery();
        }

        if ($value === "Azure Metric Names") {
            return self::metricNamesQuery();
        }

        if ($value === "Azure Workspaces") {
            return self::workspacesQuery();
        }

        if ($value === "Azure Regions") {
            return self::locationsQuery();
        }

        if ($value === "Grafana Template Variable Function") {
            return self::grafanaTemplateVariableFn();
        }

        if ($value === "traceql") {
            return self::traceExemplar();
        }

        if ($value === "Azure Custom Namespaces") {
            return self::customNamespacesQuery();
        }

        if ($value === "Azure Custom Metric Names") {
            return self::customMetricNamesQuery();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum QueryType");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

