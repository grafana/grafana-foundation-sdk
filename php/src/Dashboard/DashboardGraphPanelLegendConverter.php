<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardGraphPanelLegendConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DashboardGraphPanelLegend $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardGraphPanelLegendBuilder())',
        ];
            if ($input->show !== true) {
    
        
    $buffer = 'show(';
        $arg0 =\var_export($input->show, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sort !== null && $input->sort !== "") {
    
        
    $buffer = 'sort(';
        $arg0 =\var_export($input->sort, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sortDesc !== null) {
    
        
    $buffer = 'sortDesc(';
        $arg0 =\var_export($input->sortDesc, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

