<?php

namespace Grafana\Foundation\Azuremonitor;

final class AzureMonitorResourceConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\AzureMonitorResource $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\AzureMonitorResourceBuilder())',
        ];
            if ($input->subscription !== null && $input->subscription !== "") {
    
        
    $buffer = 'subscription(';
        $arg0 =\var_export($input->subscription, true);
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
            if ($input->metricNamespace !== null && $input->metricNamespace !== "") {
    
        
    $buffer = 'metricNamespace(';
        $arg0 =\var_export($input->metricNamespace, true);
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

        return \implode("\n\t->", $calls);
    }
}

