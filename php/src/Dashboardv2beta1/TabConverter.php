<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class TabConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\TabBuilder())',
        ];
            if ($input->spec->title !== null && $input->spec->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->spec->title, true);
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
            case $input->spec->layout instanceof \Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\RowsConverter::convert($input->spec->layout);
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
            if ($input->spec->variables !== null && count($input->spec->variables) >= 1) {
    
        
    $buffer = 'variables(';
        $tmparg0 = [];
        foreach ($input->spec->variables as $arg1) {
        switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\QueryVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\QueryVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\TextVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\TextVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\ConstantVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\DatasourceVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\IntervalVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\CustomVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\CustomVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\GroupByVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\AdhocVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\SwitchVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $tmparg0[] = $tmpvariablesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

