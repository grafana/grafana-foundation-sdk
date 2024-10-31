<?php

namespace Grafana\Foundation\Azuremonitor;

final class AzureMonitorQueryConverter
{
    public static function convert(\Grafana\Foundation\Cog\Dataquery $input): string
    {
        assert($input instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\AzureMonitorQueryBuilder())',
        ];
            if ($input->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryType !== null && $input->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->subscription !== null && $input->subscription !== "") {
    
        
    $buffer = 'subscription(';
        $arg0 =\var_export($input->subscription, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->subscriptions !== null && count($input->subscriptions) >= 1) {
    
        
    $buffer = 'subscriptions(';
        $tmparg0 = [];
        foreach ($input->subscriptions as $arg1) {
        $tmpsubscriptionsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpsubscriptionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->azureMonitor !== null) {
    
        
    $buffer = 'azureMonitor(';
        $arg0 = \Grafana\Foundation\Azuremonitor\AzureMetricQueryConverter::convert($input->azureMonitor);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->azureLogAnalytics !== null) {
    
        
    $buffer = 'azureLogAnalytics(';
        $arg0 = \Grafana\Foundation\Azuremonitor\AzureLogsQueryConverter::convert($input->azureLogAnalytics);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->azureResourceGraph !== null) {
    
        
    $buffer = 'azureResourceGraph(';
        $arg0 = \Grafana\Foundation\Azuremonitor\AzureResourceGraphQueryConverter::convert($input->azureResourceGraph);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->azureTraces !== null) {
    
        
    $buffer = 'azureTraces(';
        $arg0 = \Grafana\Foundation\Azuremonitor\AzureTracesQueryConverter::convert($input->azureTraces);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'grafanaTemplateVariableFn(';
        switch (true) {
            case $input->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQueryConverter::convert($input->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\AppInsightsGroupByQueryConverter::convert($input->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\SubscriptionsQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\SubscriptionsQueryConverter::convert($input->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\ResourceGroupsQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\ResourceGroupsQueryConverter::convert($input->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\ResourceNamesQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\ResourceNamesQueryConverter::convert($input->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\MetricNamespaceQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\MetricNamespaceQueryConverter::convert($input->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\MetricDefinitionsQueryConverter::convert($input->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\MetricNamesQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\MetricNamesQueryConverter::convert($input->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\WorkspacesQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\WorkspacesQueryConverter::convert($input->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\UnknownQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\UnknownQueryConverter::convert($input->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->resourceGroup !== null && $input->resourceGroup !== "") {
    
        
    $buffer = 'resourceGroup(';
        $arg0 =\var_export($input->resourceGroup, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->namespace !== null && $input->namespace !== "") {
    
        
    $buffer = 'namespace(';
        $arg0 =\var_export($input->namespace, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resource !== null && $input->resource !== "") {
    
        
    $buffer = 'resource(';
        $arg0 =\var_export($input->resource, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->region !== null && $input->region !== "") {
    
        
    $buffer = 'region(';
        $arg0 =\var_export($input->region, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->query !== null && $input->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

