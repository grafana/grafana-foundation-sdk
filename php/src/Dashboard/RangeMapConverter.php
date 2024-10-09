<?php

namespace Grafana\Foundation\Dashboard;

final class RangeMapConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\RangeMap $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\RangeMapBuilder())',
        ];
            
    
        {
    $buffer = 'options(';
        $arg0 = \Grafana\Foundation\Dashboard\DashboardRangeMapOptionsConverter::convert($input->options);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

