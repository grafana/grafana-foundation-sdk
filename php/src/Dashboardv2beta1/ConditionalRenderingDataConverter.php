<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ConditionalRenderingDataConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataBuilder())',
        ];
            
    
        {
    $buffer = 'value(';
        $arg0 =\var_export($input->spec->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

