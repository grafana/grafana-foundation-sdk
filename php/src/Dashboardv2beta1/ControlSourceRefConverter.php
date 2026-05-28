<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ControlSourceRefConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\ControlSourceRef $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\ControlSourceRefBuilder())',
        ];
            if ($input->group !== "") {
    
        
    $buffer = 'group(';
        $arg0 =\var_export($input->group, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

