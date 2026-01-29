<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AutoGridRepeatOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptionsBuilder())',
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

