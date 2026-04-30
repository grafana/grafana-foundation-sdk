<?php

namespace Grafana\Foundation\Dashboardv2;

final class RowRepeatOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\RowRepeatOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\RowRepeatOptionsBuilder())',
        ];
            if ($input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

