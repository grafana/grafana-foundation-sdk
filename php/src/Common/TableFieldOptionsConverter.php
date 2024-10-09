<?php

namespace Grafana\Foundation\Common;

final class TableFieldOptionsConverter
{
    public static function convert(\Grafana\Foundation\Common\TableFieldOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\TableFieldOptionsBuilder())',
        ];
            if ($input->width !== null) {
    
        
    $buffer = 'width(';
        $arg0 =\var_export($input->width, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->minWidth !== null) {
    
        
    $buffer = 'minWidth(';
        $arg0 =\var_export($input->minWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'align(';
        $arg0 ='\Grafana\Foundation\Common\FieldTextAlignment::fromValue("'.$input->align.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->displayMode !== null) {
    
        
    $buffer = 'displayMode(';
        $arg0 ='\Grafana\Foundation\Common\TableCellDisplayMode::fromValue("'.$input->displayMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'cellOptions(';
        switch (true) {
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableAutoCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableAutoCellOptions(type: '.\var_export($input->cellOptions->type, true).',))';
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableSparklineCellOptions:
                $disjunctioncellOptions = \Grafana\Foundation\Common\TableSparklineCellOptionsConverter::convert($input->cellOptions);
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableBarGaugeCellOptions:
                $disjunctioncellOptions = \Grafana\Foundation\Common\TableBarGaugeCellOptionsConverter::convert($input->cellOptions);
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableColoredBackgroundCellOptions:
                $disjunctioncellOptions = \Grafana\Foundation\Common\TableColoredBackgroundCellOptionsConverter::convert($input->cellOptions);
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableColorTextCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableColorTextCellOptions(type: '.\var_export($input->cellOptions->type, true).',))';
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableImageCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableImageCellOptions(type: '.\var_export($input->cellOptions->type, true).',))';
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableDataLinksCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableDataLinksCellOptions(type: '.\var_export($input->cellOptions->type, true).',))';
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableJsonViewCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableJsonViewCellOptions(type: '.\var_export($input->cellOptions->type, true).',))';
                $arg0 = $disjunctioncellOptions;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->hidden !== null) {
    
        
    $buffer = 'hidden(';
        $arg0 =\var_export($input->hidden, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->inspect !== false) {
    
        
    $buffer = 'inspect(';
        $arg0 =\var_export($input->inspect, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->filterable !== null) {
    
        
    $buffer = 'filterable(';
        $arg0 =\var_export($input->filterable, true);
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

        return \implode("\n\t->", $calls);
    }
}

