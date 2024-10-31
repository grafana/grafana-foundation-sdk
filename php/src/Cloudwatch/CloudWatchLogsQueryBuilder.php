<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * Shape of a CloudWatch Logs query
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\CloudWatchLogsQuery>
 */
class CloudWatchLogsQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\CloudWatchLogsQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\CloudWatchLogsQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\CloudWatchLogsQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Whether a query is a Metrics, Logs, or Annotations query
     */
    public function queryMode(\Grafana\Foundation\Cloudwatch\CloudWatchQueryMode $queryMode): static
    {
        $this->internal->queryMode = $queryMode;
    
        return $this;
    }
    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    /**
     * AWS region to query for the logs
     */
    public function region(string $region): static
    {
        $this->internal->region = $region;
    
        return $this;
    }
    /**
     * The CloudWatch Logs Insights query to execute
     */
    public function expression(string $expression): static
    {
        $this->internal->expression = $expression;
    
        return $this;
    }
    /**
     * Fields to group the results by, this field is automatically populated whenever the query is updated
     * @param array<string> $statsGroups
     */
    public function statsGroups(array $statsGroups): static
    {
        $this->internal->statsGroups = $statsGroups;
    
        return $this;
    }
    /**
     * Log groups to query
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\LogGroup>> $logGroups
     */
    public function logGroups(array $logGroups): static
    {
            $logGroupsResources = [];
            foreach ($logGroups as $r1) {
                    $logGroupsResources[] = $r1->build();
            }
        $this->internal->logGroups = $logGroupsResources;
    
        return $this;
    }
    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public function refId(string $refId): static
    {
        $this->internal->refId = $refId;
    
        return $this;
    }
    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }
    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public function queryType(string $queryType): static
    {
        $this->internal->queryType = $queryType;
    
        return $this;
    }
    /**
     * @deprecated use logGroups
     * @param array<string> $logGroupNames
     */
    public function logGroupNames(array $logGroupNames): static
    {
        $this->internal->logGroupNames = $logGroupNames;
    
        return $this;
    }
    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public function datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }

}
