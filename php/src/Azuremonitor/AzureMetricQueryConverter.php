<?php

namespace Grafana\Foundation\Azuremonitor;

final class AzureMetricQueryConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\AzureMetricQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\AzureMetricQueryBuilder())',
        ];
            if ($input->resources !== null && count($input->resources) >= 1) {
    
        
    $buffer = 'resources(';
        $tmparg0 = [];
        foreach ($input->resources as $arg1) {
        $tmpresourcesarg1 = \Grafana\Foundation\Azuremonitor\AzureMonitorResourceConverter::convert($arg1);
        $tmparg0[] = $tmpresourcesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metricNamespace !== null && $input->metricNamespace !== "") {
    
        
    $buffer = 'metricNamespace(';
        $arg0 =\var_export($input->metricNamespace, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->customNamespace !== null && $input->customNamespace !== "") {
    
        
    $buffer = 'customNamespace(';
        $arg0 =\var_export($input->customNamespace, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metricName !== null && $input->metricName !== "") {
    
        
    $buffer = 'metricName(';
        $arg0 =\var_export($input->metricName, true);
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
            if ($input->timeGrain !== null && $input->timeGrain !== "") {
    
        
    $buffer = 'timeGrain(';
        $arg0 =\var_export($input->timeGrain, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->aggregation !== null && $input->aggregation !== "") {
    
        
    $buffer = 'aggregation(';
        $arg0 =\var_export($input->aggregation, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->dimensionFilters !== null && count($input->dimensionFilters) >= 1) {
    
        
    $buffer = 'dimensionFilters(';
        $tmparg0 = [];
        foreach ($input->dimensionFilters as $arg1) {
        $tmpdimensionFiltersarg1 = \Grafana\Foundation\Azuremonitor\AzureMetricDimensionConverter::convert($arg1);
        $tmparg0[] = $tmpdimensionFiltersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->top !== null && $input->top !== "") {
    
        
    $buffer = 'top(';
        $arg0 =\var_export($input->top, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->allowedTimeGrainsMs !== null && count($input->allowedTimeGrainsMs) >= 1) {
    
        
    $buffer = 'allowedTimeGrainsMs(';
        $tmparg0 = [];
        foreach ($input->allowedTimeGrainsMs as $arg1) {
        $tmpallowedTimeGrainsMsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpallowedTimeGrainsMsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->alias !== null && $input->alias !== "") {
    
        
    $buffer = 'alias(';
        $arg0 =\var_export($input->alias, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeGrainUnit !== null && $input->timeGrainUnit !== "") {
    
        
    $buffer = 'timeGrainUnit(';
        $arg0 =\var_export($input->timeGrainUnit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->dimension !== null && $input->dimension !== "") {
    
        
    $buffer = 'dimension(';
        $arg0 =\var_export($input->dimension, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->dimensionFilter !== null && $input->dimensionFilter !== "") {
    
        
    $buffer = 'dimensionFilter(';
        $arg0 =\var_export($input->dimensionFilter, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metricDefinition !== null && $input->metricDefinition !== "") {
    
        
    $buffer = 'metricDefinition(';
        $arg0 =\var_export($input->metricDefinition, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resourceUri !== null && $input->resourceUri !== "") {
    
        
    $buffer = 'resourceUri(';
        $arg0 =\var_export($input->resourceUri, true);
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
            if ($input->resourceName !== null && $input->resourceName !== "") {
    
        
    $buffer = 'resourceName(';
        $arg0 =\var_export($input->resourceName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

