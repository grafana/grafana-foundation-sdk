<?php

namespace Grafana\Foundation\Azuremonitor;

final class MetricNamespaceQueryConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\MetricNamespaceQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\MetricNamespaceQueryBuilder())',
        ];
            if ($input->rawQuery !== null && $input->rawQuery !== "") {
    
        
    $buffer = 'rawQuery(';
        $arg0 =\var_export($input->rawQuery, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->subscription !== "") {
    
        
    $buffer = 'subscription(';
        $arg0 =\var_export($input->subscription, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resourceGroup !== "") {
    
        
    $buffer = 'resourceGroup(';
        $arg0 =\var_export($input->resourceGroup, true);
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

