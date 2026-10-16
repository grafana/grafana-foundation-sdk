<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class V2PanelKindConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\V2PanelKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\V2PanelKindBuilder())',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\V2PanelSpecConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

