<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class TabRepeatOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\TabRepeatOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\TabRepeatOptionsBuilder())',
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

