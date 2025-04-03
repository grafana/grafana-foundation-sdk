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
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableAutoCellOptions(type: '.\var_export($input->cellOptions->type, true).','.(($input->cellOptions->wrapText !== null) ? 'wrapText: '.\var_export($input->cellOptions->wrapText, true).', ' : '').'))';
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableSparklineCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableSparklineCellOptions(type: '.\var_export($input->cellOptions->type, true).','.(($input->cellOptions->drawStyle !== null) ? 'drawStyle: '.'\Grafana\Foundation\Common\GraphDrawStyle::fromValue("'.$input->cellOptions->drawStyle.'")'.', ' : '').''.(($input->cellOptions->gradientMode !== null) ? 'gradientMode: '.'\Grafana\Foundation\Common\GraphGradientMode::fromValue("'.$input->cellOptions->gradientMode.'")'.', ' : '').''.(($input->cellOptions->thresholdsStyle !== null) ? 'thresholdsStyle: '.'(new \Grafana\Foundation\Common\GraphThresholdsStyleConfig(mode: '.'\Grafana\Foundation\Common\GraphThresholdsStyleMode::fromValue("'.$input->cellOptions->thresholdsStyle->mode.'")'.',))'.', ' : '').''.(($input->cellOptions->transform !== null) ? 'transform: '.'\Grafana\Foundation\Common\GraphTransform::fromValue("'.$input->cellOptions->transform.'")'.', ' : '').''.(($input->cellOptions->lineColor !== null) ? 'lineColor: '.\var_export($input->cellOptions->lineColor, true).', ' : '').''.(($input->cellOptions->lineWidth !== null) ? 'lineWidth: '.\var_export($input->cellOptions->lineWidth, true).', ' : '').''.(($input->cellOptions->lineInterpolation !== null) ? 'lineInterpolation: '.'\Grafana\Foundation\Common\LineInterpolation::fromValue("'.$input->cellOptions->lineInterpolation.'")'.', ' : '').''.(($input->cellOptions->lineStyle !== null) ? 'lineStyle: '.'(new \Grafana\Foundation\Common\LineStyle('.(($input->cellOptions->lineStyle->fill !== null) ? 'fill: '.'\Grafana\Foundation\Common\LineStyleFill::fromValue("'.$input->cellOptions->lineStyle->fill.'")'.', ' : '').''.(($input->cellOptions->lineStyle->dash !== null) ? 'dash: '.\var_export($input->cellOptions->lineStyle->dash, true).', ' : '').'))'.', ' : '').''.(($input->cellOptions->fillColor !== null) ? 'fillColor: '.\var_export($input->cellOptions->fillColor, true).', ' : '').''.(($input->cellOptions->fillOpacity !== null) ? 'fillOpacity: '.\var_export($input->cellOptions->fillOpacity, true).', ' : '').''.(($input->cellOptions->showPoints !== null) ? 'showPoints: '.'\Grafana\Foundation\Common\VisibilityMode::fromValue("'.$input->cellOptions->showPoints.'")'.', ' : '').''.(($input->cellOptions->pointSize !== null) ? 'pointSize: '.\var_export($input->cellOptions->pointSize, true).', ' : '').''.(($input->cellOptions->pointColor !== null) ? 'pointColor: '.\var_export($input->cellOptions->pointColor, true).', ' : '').''.(($input->cellOptions->axisPlacement !== null) ? 'axisPlacement: '.'\Grafana\Foundation\Common\AxisPlacement::fromValue("'.$input->cellOptions->axisPlacement.'")'.', ' : '').''.(($input->cellOptions->axisColorMode !== null) ? 'axisColorMode: '.'\Grafana\Foundation\Common\AxisColorMode::fromValue("'.$input->cellOptions->axisColorMode.'")'.', ' : '').''.(($input->cellOptions->axisLabel !== null) ? 'axisLabel: '.\var_export($input->cellOptions->axisLabel, true).', ' : '').''.(($input->cellOptions->axisWidth !== null) ? 'axisWidth: '.\var_export($input->cellOptions->axisWidth, true).', ' : '').''.(($input->cellOptions->axisSoftMin !== null) ? 'axisSoftMin: '.\var_export($input->cellOptions->axisSoftMin, true).', ' : '').''.(($input->cellOptions->axisSoftMax !== null) ? 'axisSoftMax: '.\var_export($input->cellOptions->axisSoftMax, true).', ' : '').''.(($input->cellOptions->axisGridShow !== null) ? 'axisGridShow: '.\var_export($input->cellOptions->axisGridShow, true).', ' : '').''.(($input->cellOptions->scaleDistribution !== null) ? 'scaleDistribution: '.'(new \Grafana\Foundation\Common\ScaleDistributionConfig(type: '.'\Grafana\Foundation\Common\ScaleDistribution::fromValue("'.$input->cellOptions->scaleDistribution->type.'")'.','.(($input->cellOptions->scaleDistribution->log !== null) ? 'log: '.\var_export($input->cellOptions->scaleDistribution->log, true).', ' : '').''.(($input->cellOptions->scaleDistribution->linearThreshold !== null) ? 'linearThreshold: '.\var_export($input->cellOptions->scaleDistribution->linearThreshold, true).', ' : '').'))'.', ' : '').''.(($input->cellOptions->axisCenteredZero !== null) ? 'axisCenteredZero: '.\var_export($input->cellOptions->axisCenteredZero, true).', ' : '').''.(($input->cellOptions->barAlignment !== null) ? 'barAlignment: '.'\Grafana\Foundation\Common\BarAlignment::fromValue("'.$input->cellOptions->barAlignment.'")'.', ' : '').''.(($input->cellOptions->barWidthFactor !== null) ? 'barWidthFactor: '.\var_export($input->cellOptions->barWidthFactor, true).', ' : '').''.(($input->cellOptions->stacking !== null) ? 'stacking: '.'(new \Grafana\Foundation\Common\StackingConfig('.(($input->cellOptions->stacking->mode !== null) ? 'mode: '.'\Grafana\Foundation\Common\StackingMode::fromValue("'.$input->cellOptions->stacking->mode.'")'.', ' : '').''.(($input->cellOptions->stacking->group !== null) ? 'group: '.\var_export($input->cellOptions->stacking->group, true).', ' : '').'))'.', ' : '').''.(($input->cellOptions->hideFrom !== null) ? 'hideFrom: '.'(new \Grafana\Foundation\Common\HideSeriesConfig(tooltip: '.\var_export($input->cellOptions->hideFrom->tooltip, true).',legend: '.\var_export($input->cellOptions->hideFrom->legend, true).',viz: '.\var_export($input->cellOptions->hideFrom->viz, true).',))'.', ' : '').''.(($input->cellOptions->hideValue !== null) ? 'hideValue: '.\var_export($input->cellOptions->hideValue, true).', ' : '').''.(($input->cellOptions->insertNulls !== null) ? 'insertNulls: '.\var_export($input->cellOptions->insertNulls, true).', ' : '').''.(($input->cellOptions->spanNulls !== null) ? 'spanNulls: '.\var_export($input->cellOptions->spanNulls, true).', ' : '').''.(($input->cellOptions->fillBelowTo !== null) ? 'fillBelowTo: '.\var_export($input->cellOptions->fillBelowTo, true).', ' : '').''.(($input->cellOptions->pointSymbol !== null) ? 'pointSymbol: '.\var_export($input->cellOptions->pointSymbol, true).', ' : '').''.(($input->cellOptions->axisBorderShow !== null) ? 'axisBorderShow: '.\var_export($input->cellOptions->axisBorderShow, true).', ' : '').''.(($input->cellOptions->barMaxWidth !== null) ? 'barMaxWidth: '.\var_export($input->cellOptions->barMaxWidth, true).', ' : '').'))';
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableBarGaugeCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableBarGaugeCellOptions(type: '.\var_export($input->cellOptions->type, true).','.(($input->cellOptions->mode !== null) ? 'mode: '.'\Grafana\Foundation\Common\BarGaugeDisplayMode::fromValue("'.$input->cellOptions->mode.'")'.', ' : '').''.(($input->cellOptions->valueDisplayMode !== null) ? 'valueDisplayMode: '.'\Grafana\Foundation\Common\BarGaugeValueMode::fromValue("'.$input->cellOptions->valueDisplayMode.'")'.', ' : '').'))';
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableColoredBackgroundCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableColoredBackgroundCellOptions(type: '.\var_export($input->cellOptions->type, true).','.(($input->cellOptions->mode !== null) ? 'mode: '.'\Grafana\Foundation\Common\TableCellBackgroundDisplayMode::fromValue("'.$input->cellOptions->mode.'")'.', ' : '').''.(($input->cellOptions->applyToRow !== null) ? 'applyToRow: '.\var_export($input->cellOptions->applyToRow, true).', ' : '').''.(($input->cellOptions->wrapText !== null) ? 'wrapText: '.\var_export($input->cellOptions->wrapText, true).', ' : '').'))';
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableColorTextCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableColorTextCellOptions(type: '.\var_export($input->cellOptions->type, true).','.(($input->cellOptions->wrapText !== null) ? 'wrapText: '.\var_export($input->cellOptions->wrapText, true).', ' : '').'))';
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableImageCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableImageCellOptions(type: '.\var_export($input->cellOptions->type, true).','.(($input->cellOptions->alt !== null) ? 'alt: '.\var_export($input->cellOptions->alt, true).', ' : '').''.(($input->cellOptions->title !== null) ? 'title: '.\var_export($input->cellOptions->title, true).', ' : '').'))';
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableDataLinksCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableDataLinksCellOptions(type: '.\var_export($input->cellOptions->type, true).',))';
                $arg0 = $disjunctioncellOptions;
                break;
            case $input->cellOptions instanceof \Grafana\Foundation\Common\TableActionsCellOptions:
                $disjunctioncellOptions ='(new \Grafana\Foundation\Common\TableActionsCellOptions(type: '.\var_export($input->cellOptions->type, true).',))';
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

