<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class PanelRefConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\PanelRef $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\PanelRefBuilder())',
        ];
            if ($input->dashboardUID !== "") {
    
        
    $buffer = 'dashboardUID(';
        $arg0 =\var_export($input->dashboardUID, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'panelID(';
        $arg0 =\var_export($input->panelID, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

