<?php

namespace Grafana\Foundation\Azuremonitor;

final class ResourceGroupsQueryConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\ResourceGroupsQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\ResourceGroupsQueryBuilder())',
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

        return \implode("\n\t->", $calls);
    }
}

