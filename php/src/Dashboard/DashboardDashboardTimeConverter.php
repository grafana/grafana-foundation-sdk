<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardDashboardTimeConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DashboardDashboardTime $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardDashboardTimeBuilder())',
        ];
            if ($input->from !== "" && $input->from !== "now-6h") {
    
        
    $buffer = 'from(';
        $arg0 =\var_export($input->from, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->to !== "" && $input->to !== "now") {
    
        
    $buffer = 'to(';
        $arg0 =\var_export($input->to, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

