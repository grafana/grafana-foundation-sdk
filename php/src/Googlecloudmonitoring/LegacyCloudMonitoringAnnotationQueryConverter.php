<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class LegacyCloudMonitoringAnnotationQueryConverter
{
    public static function convert(\Grafana\Foundation\Googlecloudmonitoring\LegacyCloudMonitoringAnnotationQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Googlecloudmonitoring\LegacyCloudMonitoringAnnotationQueryBuilder())',
        ];
            if ($input->projectName !== "") {
    
        
    $buffer = 'projectName(';
        $arg0 =\var_export($input->projectName, true);
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
            if ($input->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->filters) >= 1) {
    
        
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
            
    
        {
    $buffer = 'metricKind(';
        $arg0 ='\Grafana\Foundation\Googlecloudmonitoring\MetricKind::fromValue("'.$input->metricKind.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->valueType !== "") {
    
        
    $buffer = 'valueType(';
        $arg0 =\var_export($input->valueType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->text !== "") {
    
        
    $buffer = 'text(';
        $arg0 =\var_export($input->text, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

