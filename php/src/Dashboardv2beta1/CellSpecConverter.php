<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class CellSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\CellSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\CellSpecBuilder())',
        ];
            
    
        {
    $buffer = 'content(';
        switch (true) {
            case $input->content instanceof \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentKind:
                $disjunctioncontent = \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentKindConverter::convert($input->content);
                $arg0 = $disjunctioncontent;
                break;
            case $input->content instanceof \Grafana\Foundation\Dashboardv2beta1\CodeCellContentKind:
                $disjunctioncontent = \Grafana\Foundation\Dashboardv2beta1\CodeCellContentKindConverter::convert($input->content);
                $arg0 = $disjunctioncontent;
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

