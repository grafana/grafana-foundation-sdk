<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class MetricQueryConverter
{
    public static function convert(\Grafana\Foundation\Googlecloudmonitoring\MetricQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Googlecloudmonitoring\MetricQueryBuilder())',
        ];
            if ($input->projectName !== "") {
    
        
    $buffer = 'projectName(';
        $arg0 =\var_export($input->projectName, true);
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
            if ($input->alignmentPeriod !== null && $input->alignmentPeriod !== "") {
    
        
    $buffer = 'alignmentPeriod(';
        $arg0 =\var_export($input->alignmentPeriod, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->aliasBy !== null && $input->aliasBy !== "") {
    
        
    $buffer = 'aliasBy(';
        $arg0 =\var_export($input->aliasBy, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->editorMode !== "") {
    
        
    $buffer = 'editorMode(';
        $arg0 =\var_export($input->editorMode, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metricType !== "") {
    
        
    $buffer = 'metricType(';
        $arg0 =\var_export($input->metricType, true);
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
            if ($input->metricKind !== null) {
    
        
    $buffer = 'metricKind(';
        $arg0 ='\Grafana\Foundation\Googlecloudmonitoring\MetricKind::fromValue("'.$input->metricKind.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->valueType !== null && $input->valueType !== "") {
    
        
    $buffer = 'valueType(';
        $arg0 =\var_export($input->valueType, true);
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
            if ($input->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->query, true);
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
            if ($input->graphPeriod !== null && $input->graphPeriod !== "" && $input->graphPeriod !== "disabled") {
    
        
    $buffer = 'graphPeriod(';
        $arg0 =\var_export($input->graphPeriod, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

