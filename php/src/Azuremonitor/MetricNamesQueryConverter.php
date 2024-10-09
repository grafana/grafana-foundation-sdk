<?php

namespace Grafana\Foundation\Azuremonitor;

final class MetricNamesQueryConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\MetricNamesQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\MetricNamesQueryBuilder())',
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
            if ($input->resourceName !== "") {
    
        
    $buffer = 'resourceName(';
        $arg0 =\var_export($input->resourceName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metricNamespace !== "") {
    
        
    $buffer = 'metricNamespace(';
        $arg0 =\var_export($input->metricNamespace, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

