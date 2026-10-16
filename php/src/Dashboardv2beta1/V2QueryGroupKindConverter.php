<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class V2QueryGroupKindConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\V2QueryGroupKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\V2QueryGroupKindBuilder())',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\V2QueryGroupSpecConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

