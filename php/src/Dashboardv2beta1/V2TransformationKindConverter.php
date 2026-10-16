<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class V2TransformationKindConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\V2TransformationKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\V2TransformationKindBuilder())',
        ];
            if ($input->group !== "") {
    
        
    $buffer = 'group(';
        $arg0 =\var_export($input->group, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'spec(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\V2TransformationSpecConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

