<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class TabsLayoutTabSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabSpecBuilder())',
        ];
            if ($input->title !== null && $input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'layout(';
        switch (true) {
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\GridLayoutConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\RowsLayoutConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\TabsLayoutConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->conditionalRendering !== null) {
    
        
    $buffer = 'conditionalRendering(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupConverter::convert($input->conditionalRendering);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->repeat !== null) {
    
        
    $buffer = 'repeat(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\TabRepeatOptionsConverter::convert($input->repeat);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

