<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class RowsLayoutRowSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowSpecBuilder())',
        ];
            if ($input->title !== null && $input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->collapse !== null) {
    
        
    $buffer = 'collapse(';
        $arg0 =\var_export($input->collapse, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hideHeader !== null) {
    
        
    $buffer = 'hideHeader(';
        $arg0 =\var_export($input->hideHeader, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fillScreen !== null) {
    
        
    $buffer = 'fillScreen(';
        $arg0 =\var_export($input->fillScreen, true);
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
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\RowRepeatOptionsConverter::convert($input->repeat);
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
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\TabsLayoutConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\RowsLayoutConverter::convert($input->layout);
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

