<?php

namespace Grafana\Foundation\Dashboardv2;

final class ControlSourceRefConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\ControlSourceRef $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\ControlSourceRefBuilder())',
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

