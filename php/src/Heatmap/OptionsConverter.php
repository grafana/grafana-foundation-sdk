<?php

namespace Grafana\Foundation\Heatmap;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Heatmap\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Heatmap\OptionsBuilder())',
        ];
            if ($input->calculate !== null && $input->calculate !== false) {
    
        
    $buffer = 'calculate(';
        $arg0 =\var_export($input->calculate, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->calculation !== null) {
    
        
    $buffer = 'calculation(';
        $arg0 = \Grafana\Foundation\Common\HeatmapCalculationOptionsConverter::convert($input->calculation);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'color(';
        $arg0 = \Grafana\Foundation\Heatmap\HeatmapColorOptionsConverter::convert($input->color);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->filterValues !== null) {
    
        
    $buffer = 'filterValues(';
        $arg0 = \Grafana\Foundation\Heatmap\FilterValueRangeConverter::convert($input->filterValues);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->rowsFrame !== null) {
    
        
    $buffer = 'rowsFrame(';
        $arg0 = \Grafana\Foundation\Heatmap\RowsHeatmapOptionsConverter::convert($input->rowsFrame);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'showValue(';
        $arg0 ='\Grafana\Foundation\Common\VisibilityMode::fromValue("'.$input->showValue.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->cellGap !== null && $input->cellGap !== 1) {
    
        
    $buffer = 'cellGap(';
        $arg0 =\var_export($input->cellGap, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->cellRadius !== null) {
    
        
    $buffer = 'cellRadius(';
        $arg0 =\var_export($input->cellRadius, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->cellValues !== null) {
    
        
    $buffer = 'cellValues(';
        $arg0 = \Grafana\Foundation\Heatmap\CellValuesConverter::convert($input->cellValues);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'yAxis(';
        $arg0 = \Grafana\Foundation\Heatmap\YAxisConfigConverter::convert($input->yAxis);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Heatmap\HeatmapLegendConverter::convert($input->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'tooltip(';
        $arg0 = \Grafana\Foundation\Heatmap\HeatmapTooltipConverter::convert($input->tooltip);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'exemplars(';
        $arg0 = \Grafana\Foundation\Heatmap\ExemplarConfigConverter::convert($input->exemplars);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

