<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class TimeSeriesListConverter
{
    public static function convert(\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesListBuilder())',
        ];
            if ($input->projectName !== "") {
    
        
    $buffer = 'projectName(';
        $arg0 =\var_export($input->projectName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->crossSeriesReducer !== "") {
    
        
    $buffer = 'crossSeriesReducer(';
        $arg0 =\var_export($input->crossSeriesReducer, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->alignmentPeriod !== null && $input->alignmentPeriod !== "") {
    
        
    $buffer = 'alignmentPeriod(';
        $arg0 =\var_export($input->alignmentPeriod, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->perSeriesAligner !== null && $input->perSeriesAligner !== "") {
    
        
    $buffer = 'perSeriesAligner(';
        $arg0 =\var_export($input->perSeriesAligner, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupBys !== null && count($input->groupBys) >= 1) {
    
        
    $buffer = 'groupBys(';
        $tmparg0 = [];
        foreach ($input->groupBys as $arg1) {
        $tmpgroupBysarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpgroupBysarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->filters !== null && count($input->filters) >= 1) {
    
        
    $buffer = 'filters(';
        $tmparg0 = [];
        foreach ($input->filters as $arg1) {
        $tmpfiltersarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpfiltersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->view !== null && $input->view !== "") {
    
        
    $buffer = 'view(';
        $arg0 =\var_export($input->view, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->title !== null && $input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->text !== null && $input->text !== "") {
    
        
    $buffer = 'text(';
        $arg0 =\var_export($input->text, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->secondaryCrossSeriesReducer !== null && $input->secondaryCrossSeriesReducer !== "") {
    
        
    $buffer = 'secondaryCrossSeriesReducer(';
        $arg0 =\var_export($input->secondaryCrossSeriesReducer, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->secondaryAlignmentPeriod !== null && $input->secondaryAlignmentPeriod !== "") {
    
        
    $buffer = 'secondaryAlignmentPeriod(';
        $arg0 =\var_export($input->secondaryAlignmentPeriod, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->secondaryPerSeriesAligner !== null && $input->secondaryPerSeriesAligner !== "") {
    
        
    $buffer = 'secondaryPerSeriesAligner(';
        $arg0 =\var_export($input->secondaryPerSeriesAligner, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->secondaryGroupBys !== null && count($input->secondaryGroupBys) >= 1) {
    
        
    $buffer = 'secondaryGroupBys(';
        $tmparg0 = [];
        foreach ($input->secondaryGroupBys as $arg1) {
        $tmpsecondaryGroupBysarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpsecondaryGroupBysarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->preprocessor !== null) {
    
        
    $buffer = 'preprocessor(';
        $arg0 ='\Grafana\Foundation\Googlecloudmonitoring\PreprocessorType::fromValue("'.$input->preprocessor.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

