<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class RowConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\RowBuilder())',
        ];
            if ($input->spec->title !== null && $input->spec->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->spec->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->collapse !== null) {
    
        
    $buffer = 'collapse(';
        $arg0 =\var_export($input->spec->collapse, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->hideHeader !== null) {
    
        
    $buffer = 'hideHeader(';
        $arg0 =\var_export($input->spec->hideHeader, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fillScreen !== null) {
    
        
    $buffer = 'fillScreen(';
        $arg0 =\var_export($input->spec->fillScreen, true);
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
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\RowRepeatOptionsConverter::convert($input->spec->repeat);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'layout(';
        switch (true) {
            case $input->spec->layout instanceof \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\GridConverter::convert($input->spec->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->spec->layout instanceof \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\AutoGridConverter::convert($input->spec->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->spec->layout instanceof \Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\TabsConverter::convert($input->spec->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->spec->layout instanceof \Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\RowsConverter::convert($input->spec->layout);
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

