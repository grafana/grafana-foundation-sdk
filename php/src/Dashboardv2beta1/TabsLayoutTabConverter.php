<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class TabsLayoutTabConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabBuilder('.\var_export($input->spec->title, true).'))',
        ];
            if ($input->spec->title !== null && $input->spec->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->spec->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'gridLayout(';
        switch (true) {
            case $input->spec->layout instanceof \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\GridLayoutConverter::convert($input->spec->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->spec->layout instanceof \Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\RowsLayoutConverter::convert($input->spec->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->spec->layout instanceof \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutConverter::convert($input->spec->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->spec->layout instanceof \Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\TabsLayoutConverter::convert($input->spec->layout);
                $arg0 = $disjunctionlayout;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->conditionalRendering !== null) {
    
        
    $buffer = 'conditionalRendering(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupConverter::convert($input->spec->conditionalRendering);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->repeat !== null) {
    
        
    $buffer = 'repeat(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\TabRepeatOptionsConverter::convert($input->spec->repeat);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

