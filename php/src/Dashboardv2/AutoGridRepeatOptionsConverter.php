<?php

namespace Grafana\Foundation\Dashboardv2;

final class AutoGridRepeatOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\AutoGridRepeatOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\AutoGridRepeatOptionsBuilder())',
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

