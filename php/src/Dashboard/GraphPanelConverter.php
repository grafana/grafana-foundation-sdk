<?php

namespace Grafana\Foundation\Dashboard;

final class GraphPanelConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\GraphPanel $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\GraphPanelBuilder())',
        ];
            if ($input->legend !== null) {
    
        
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Dashboard\DashboardGraphPanelLegendConverter::convert($input->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

