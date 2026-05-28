<?php

namespace Grafana\Foundation\Dashboardv2;

final class PreferencesConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\Preferences $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\PreferencesBuilder())',
        ];
            if ($input->layout !== null) {
    
        
    $buffer = 'layout(';
        switch (true) {
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2\AutoGridLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2\AutoGridConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2\GridLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2\GridConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

