<?php

namespace Grafana\Foundation\Azuremonitor;

final class AzureLogsQueryConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\AzureLogsQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\AzureLogsQueryBuilder())',
        ];
            if ($input->query !== null && $input->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resultFormat !== null) {
    
        
    $buffer = 'resultFormat(';
        $arg0 ='\Grafana\Foundation\Azuremonitor\ResultFormat::fromValue("'.$input->resultFormat.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resources !== null && count($input->resources) >= 1) {
    
        
    $buffer = 'resources(';
        $tmparg0 = [];
        foreach ($input->resources as $arg1) {
        $tmpresourcesarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpresourcesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->dashboardTime !== null) {
    
        
    $buffer = 'dashboardTime(';
        $arg0 =\var_export($input->dashboardTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeColumn !== null && $input->timeColumn !== "") {
    
        
    $buffer = 'timeColumn(';
        $arg0 =\var_export($input->timeColumn, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->workspace !== null && $input->workspace !== "") {
    
        
    $buffer = 'workspace(';
        $arg0 =\var_export($input->workspace, true);
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
            if ($input->intersectTime !== null) {
    
        
    $buffer = 'intersectTime(';
        $arg0 =\var_export($input->intersectTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

