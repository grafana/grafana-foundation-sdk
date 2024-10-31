---
title: <span class="badge object-type-enum"></span> GrafanaTemplateVariableQueryType
---
# <span class="badge object-type-enum"></span> GrafanaTemplateVariableQueryType

## Definition

```php
final class GrafanaTemplateVariableQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, GrafanaTemplateVariableQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function appInsightsMetricNameQuery(): self
    {
        if (!isset(self::$instances["AppInsightsMetricNameQuery"])) {
            self::$instances["AppInsightsMetricNameQuery"] = new self("AppInsightsMetricNameQuery");
        }

        return self::$instances["AppInsightsMetricNameQuery"];
    }

    public static function appInsightsGroupByQuery(): self
    {
        if (!isset(self::$instances["AppInsightsGroupByQuery"])) {
            self::$instances["AppInsightsGroupByQuery"] = new self("AppInsightsGroupByQuery");
        }

        return self::$instances["AppInsightsGroupByQuery"];
    }

    public static function subscriptionsQuery(): self
    {
        if (!isset(self::$instances["SubscriptionsQuery"])) {
            self::$instances["SubscriptionsQuery"] = new self("SubscriptionsQuery");
        }

        return self::$instances["SubscriptionsQuery"];
    }

    public static function resourceGroupsQuery(): self
    {
        if (!isset(self::$instances["ResourceGroupsQuery"])) {
            self::$instances["ResourceGroupsQuery"] = new self("ResourceGroupsQuery");
        }

        return self::$instances["ResourceGroupsQuery"];
    }

    public static function resourceNamesQuery(): self
    {
        if (!isset(self::$instances["ResourceNamesQuery"])) {
            self::$instances["ResourceNamesQuery"] = new self("ResourceNamesQuery");
        }

        return self::$instances["ResourceNamesQuery"];
    }

    public static function metricNamespaceQuery(): self
    {
        if (!isset(self::$instances["MetricNamespaceQuery"])) {
            self::$instances["MetricNamespaceQuery"] = new self("MetricNamespaceQuery");
        }

        return self::$instances["MetricNamespaceQuery"];
    }

    public static function metricNamesQuery(): self
    {
        if (!isset(self::$instances["MetricNamesQuery"])) {
            self::$instances["MetricNamesQuery"] = new self("MetricNamesQuery");
        }

        return self::$instances["MetricNamesQuery"];
    }

    public static function workspacesQuery(): self
    {
        if (!isset(self::$instances["WorkspacesQuery"])) {
            self::$instances["WorkspacesQuery"] = new self("WorkspacesQuery");
        }

        return self::$instances["WorkspacesQuery"];
    }

    public static function unknownQuery(): self
    {
        if (!isset(self::$instances["UnknownQuery"])) {
            self::$instances["UnknownQuery"] = new self("UnknownQuery");
        }

        return self::$instances["UnknownQuery"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "AppInsightsMetricNameQuery") {
            return self::appInsightsMetricNameQuery();
        }

        if ($value === "AppInsightsGroupByQuery") {
            return self::appInsightsGroupByQuery();
        }

        if ($value === "SubscriptionsQuery") {
            return self::subscriptionsQuery();
        }

        if ($value === "ResourceGroupsQuery") {
            return self::resourceGroupsQuery();
        }

        if ($value === "ResourceNamesQuery") {
            return self::resourceNamesQuery();
        }

        if ($value === "MetricNamespaceQuery") {
            return self::metricNamespaceQuery();
        }

        if ($value === "MetricNamesQuery") {
            return self::metricNamesQuery();
        }

        if ($value === "WorkspacesQuery") {
            return self::workspacesQuery();
        }

        if ($value === "UnknownQuery") {
            return self::unknownQuery();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum GrafanaTemplateVariableQueryType");
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

```
