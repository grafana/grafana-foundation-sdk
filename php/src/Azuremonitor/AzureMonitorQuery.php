<?php

namespace Grafana\Foundation\Azuremonitor;

class AzureMonitorQuery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * Note this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    /**
     * Azure subscription containing the resource(s) to be queried.
     */
    public ?string $subscription;

    /**
     * Subscriptions to be queried via Azure Resource Graph.
     * @var array<string>|null
     */
    public ?array $subscriptions;

    /**
     * Azure Monitor Metrics sub-query properties.
     */
    public ?\Grafana\Foundation\Azuremonitor\AzureMetricQuery $azureMonitor;

    /**
     * Azure Monitor Logs sub-query properties.
     */
    public ?\Grafana\Foundation\Azuremonitor\AzureLogsQuery $azureLogAnalytics;

    /**
     * Azure Resource Graph sub-query properties.
     */
    public ?\Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery $azureResourceGraph;

    /**
     * Application Insights Traces sub-query properties.
     */
    public ?\Grafana\Foundation\Azuremonitor\AzureTracesQuery $azureTraces;

    /**
     * @deprecated Legacy template variable support.
     * @var \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery|\Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery|\Grafana\Foundation\Azuremonitor\SubscriptionsQuery|\Grafana\Foundation\Azuremonitor\ResourceGroupsQuery|\Grafana\Foundation\Azuremonitor\ResourceNamesQuery|\Grafana\Foundation\Azuremonitor\MetricNamespaceQuery|\Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery|\Grafana\Foundation\Azuremonitor\MetricNamesQuery|\Grafana\Foundation\Azuremonitor\WorkspacesQuery|\Grafana\Foundation\Azuremonitor\UnknownQuery
     */
    public $grafanaTemplateVariableFn;

    /**
     * Template variables params. These exist for backwards compatiblity with legacy template variables.
     */
    public ?string $resourceGroup;

    public ?string $namespace;

    public ?string $resource;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * Azure Monitor query type.
     * queryType: #AzureQueryType
     */
    public ?string $region;

    /**
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param string|null $subscription
     * @param array<string>|null $subscriptions
     * @param \Grafana\Foundation\Azuremonitor\AzureMetricQuery|null $azureMonitor
     * @param \Grafana\Foundation\Azuremonitor\AzureLogsQuery|null $azureLogAnalytics
     * @param \Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery|null $azureResourceGraph
     * @param \Grafana\Foundation\Azuremonitor\AzureTracesQuery|null $azureTraces
     * @param \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery|\Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery|\Grafana\Foundation\Azuremonitor\SubscriptionsQuery|\Grafana\Foundation\Azuremonitor\ResourceGroupsQuery|\Grafana\Foundation\Azuremonitor\ResourceNamesQuery|\Grafana\Foundation\Azuremonitor\MetricNamespaceQuery|\Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery|\Grafana\Foundation\Azuremonitor\MetricNamesQuery|\Grafana\Foundation\Azuremonitor\WorkspacesQuery|\Grafana\Foundation\Azuremonitor\UnknownQuery|null $grafanaTemplateVariableFn
     * @param string|null $resourceGroup
     * @param string|null $namespace
     * @param string|null $resource
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param string|null $region
     */
    public function __construct(?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?string $subscription = null, ?array $subscriptions = null, ?\Grafana\Foundation\Azuremonitor\AzureMetricQuery $azureMonitor = null, ?\Grafana\Foundation\Azuremonitor\AzureLogsQuery $azureLogAnalytics = null, ?\Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery $azureResourceGraph = null, ?\Grafana\Foundation\Azuremonitor\AzureTracesQuery $azureTraces = null,  $grafanaTemplateVariableFn = null, ?string $resourceGroup = null, ?string $namespace = null, ?string $resource = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?string $region = null)
    {
        $this->refId = $refId ?: "";
        $this->hide = $hide;
        $this->queryType = $queryType;
        $this->subscription = $subscription;
        $this->subscriptions = $subscriptions;
        $this->azureMonitor = $azureMonitor;
        $this->azureLogAnalytics = $azureLogAnalytics;
        $this->azureResourceGraph = $azureResourceGraph;
        $this->azureTraces = $azureTraces;
        $this->grafanaTemplateVariableFn = $grafanaTemplateVariableFn ?: new \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery();
        $this->resourceGroup = $resourceGroup;
        $this->namespace = $namespace;
        $this->resource = $resource;
        $this->datasource = $datasource;
        $this->region = $region;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{refId?: string, hide?: bool, queryType?: string, subscription?: string, subscriptions?: array<string>, azureMonitor?: mixed, azureLogAnalytics?: mixed, azureResourceGraph?: mixed, azureTraces?: mixed, grafanaTemplateVariableFn?: mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed, resourceGroup?: string, namespace?: string, resource?: string, datasource?: mixed, region?: string} $inputData */
        $data = $inputData;
        return new self(
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
            subscription: $data["subscription"] ?? null,
            subscriptions: $data["subscriptions"] ?? null,
            azureMonitor: isset($data["azureMonitor"]) ? (function($input) {
    	/** @var array{resources?: array<mixed>, metricNamespace?: string, customNamespace?: string, metricName?: string, region?: string, timeGrain?: string, aggregation?: string, dimensionFilters?: array<mixed>, top?: string, allowedTimeGrainsMs?: array<int>, alias?: string, timeGrainUnit?: string, dimension?: string, dimensionFilter?: string, metricDefinition?: string, resourceUri?: string, resourceGroup?: string, resourceName?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\AzureMetricQuery::fromArray($val);
    })($data["azureMonitor"]) : null,
            azureLogAnalytics: isset($data["azureLogAnalytics"]) ? (function($input) {
    	/** @var array{query?: string, resultFormat?: string, resources?: array<string>, dashboardTime?: bool, timeColumn?: string, workspace?: string, resource?: string, intersectTime?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\AzureLogsQuery::fromArray($val);
    })($data["azureLogAnalytics"]) : null,
            azureResourceGraph: isset($data["azureResourceGraph"]) ? (function($input) {
    	/** @var array{query?: string, resultFormat?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery::fromArray($val);
    })($data["azureResourceGraph"]) : null,
            azureTraces: isset($data["azureTraces"]) ? (function($input) {
    	/** @var array{resultFormat?: string, resources?: array<string>, operationId?: string, traceTypes?: array<string>, filters?: array<mixed>, query?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\AzureTracesQuery::fromArray($val);
    })($data["azureTraces"]) : null,
            grafanaTemplateVariableFn: isset($data["grafanaTemplateVariableFn"]) ? (function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
    
        switch ($input["kind"]) {
        case "AppInsightsGroupByQuery":
            return AppInsightsGroupByQuery::fromArray($input);
        case "AppInsightsMetricNameQuery":
            return AppInsightsMetricNameQuery::fromArray($input);
        case "MetricDefinitionsQuery":
            return MetricDefinitionsQuery::fromArray($input);
        case "MetricNamesQuery":
            return MetricNamesQuery::fromArray($input);
        case "MetricNamespaceQuery":
            return MetricNamespaceQuery::fromArray($input);
        case "ResourceGroupsQuery":
            return ResourceGroupsQuery::fromArray($input);
        case "ResourceNamesQuery":
            return ResourceNamesQuery::fromArray($input);
        case "SubscriptionsQuery":
            return SubscriptionsQuery::fromArray($input);
        case "UnknownQuery":
            return UnknownQuery::fromArray($input);
        case "WorkspacesQuery":
            return WorkspacesQuery::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    })($data["grafanaTemplateVariableFn"]) : null,
            resourceGroup: $data["resourceGroup"] ?? null,
            namespace: $data["namespace"] ?? null,
            resource: $data["resource"] ?? null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            region: $data["region"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "refId" => $this->refId,
            "grafanaTemplateVariableFn" => $this->grafanaTemplateVariableFn,
        ];
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        if (isset($this->queryType)) {
            $data["queryType"] = $this->queryType;
        }
        if (isset($this->subscription)) {
            $data["subscription"] = $this->subscription;
        }
        if (isset($this->subscriptions)) {
            $data["subscriptions"] = $this->subscriptions;
        }
        if (isset($this->azureMonitor)) {
            $data["azureMonitor"] = $this->azureMonitor;
        }
        if (isset($this->azureLogAnalytics)) {
            $data["azureLogAnalytics"] = $this->azureLogAnalytics;
        }
        if (isset($this->azureResourceGraph)) {
            $data["azureResourceGraph"] = $this->azureResourceGraph;
        }
        if (isset($this->azureTraces)) {
            $data["azureTraces"] = $this->azureTraces;
        }
        if (isset($this->resourceGroup)) {
            $data["resourceGroup"] = $this->resourceGroup;
        }
        if (isset($this->namespace)) {
            $data["namespace"] = $this->namespace;
        }
        if (isset($this->resource)) {
            $data["resource"] = $this->resource;
        }
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        if (isset($this->region)) {
            $data["region"] = $this->region;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "grafana-azure-monitor-datasource";
    }
}
