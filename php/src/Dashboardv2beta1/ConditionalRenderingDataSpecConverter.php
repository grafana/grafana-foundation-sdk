<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ConditionalRenderingDataSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataSpecBuilder())',
        ];
            
    
        {
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

