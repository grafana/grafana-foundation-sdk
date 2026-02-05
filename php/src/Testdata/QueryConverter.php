<?php

namespace Grafana\Foundation\Testdata;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Testdata\QueryBuilder())',
        ];
            if ($input->version !== "" && $input->version !== "v0") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasourceConverter::convert($input->datasource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->alias !== null && $input->spec->alias !== "") {
    
        
    $buffer = 'alias(';
        $arg0 =\var_export($input->spec->alias, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->channel !== null && $input->spec->channel !== "") {
    
        
    $buffer = 'channel(';
        $arg0 =\var_export($input->spec->channel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->csvContent !== null && $input->spec->csvContent !== "") {
    
        
    $buffer = 'csvContent(';
        $arg0 =\var_export($input->spec->csvContent, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->csvFileName !== null && $input->spec->csvFileName !== "") {
    
        
    $buffer = 'csvFileName(';
        $arg0 =\var_export($input->spec->csvFileName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->csvWave !== null && count($input->spec->csvWave) >= 1) {
    
        
    $buffer = 'csvWave(';
        $tmparg0 = [];
        foreach ($input->spec->csvWave as $arg1) {
        $tmpcsvWavearg1 = \Grafana\Foundation\Testdata\CSVWaveConverter::convert($arg1);
        $tmparg0[] = $tmpcsvWavearg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->dropPercent !== null) {
    
        
    $buffer = 'dropPercent(';
        $arg0 =\var_export($input->spec->dropPercent, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->errorSource !== null) {
    
        
    $buffer = 'errorSource(';
        $arg0 ='\Grafana\Foundation\Testdata\DataqueryErrorSource::fromValue("'.$input->spec->errorSource.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->errorType !== null) {
    
        
    $buffer = 'errorType(';
        $arg0 ='\Grafana\Foundation\Testdata\DataqueryErrorType::fromValue("'.$input->spec->errorType.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->flamegraphDiff !== null) {
    
        
    $buffer = 'flamegraphDiff(';
        $arg0 =\var_export($input->spec->flamegraphDiff, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->spec->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->intervalMs !== null) {
    
        
    $buffer = 'intervalMs(';
        $arg0 =\var_export($input->spec->intervalMs, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->labels !== null && $input->spec->labels !== "") {
    
        
    $buffer = 'labels(';
        $arg0 =\var_export($input->spec->labels, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->levelColumn !== null) {
    
        
    $buffer = 'levelColumn(';
        $arg0 =\var_export($input->spec->levelColumn, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->lines !== null) {
    
        
    $buffer = 'lines(';
        $arg0 =\var_export($input->spec->lines, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->max !== null) {
    
        
    $buffer = 'max(';
        $arg0 =\var_export($input->spec->max, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->maxDataPoints !== null) {
    
        
    $buffer = 'maxDataPoints(';
        $arg0 =\var_export($input->spec->maxDataPoints, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->min !== null) {
    
        
    $buffer = 'min(';
        $arg0 =\var_export($input->spec->min, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->nodes !== null) {
    
        
    $buffer = 'nodes(';
        $arg0 = \Grafana\Foundation\Testdata\NodesQueryConverter::convert($input->spec->nodes);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->noise !== null) {
    
        
    $buffer = 'noise(';
        $arg0 =\var_export($input->spec->noise, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->points !== null && count($input->spec->points) >= 1) {
    
        
    $buffer = 'points(';
        $tmparg0 = [];
        foreach ($input->spec->points as $arg1) {
        $tmptmppointsarg1 = [];
        foreach ($arg1 as $arg1Value) {
        $tmparg1arg1Value =\var_export($arg1Value, true);
        $tmptmppointsarg1[] = $tmparg1arg1Value;
        }
        $tmppointsarg1 = "[" . implode(", \n", $tmptmppointsarg1) . "]";
        $tmparg0[] = $tmppointsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->pulseWave !== null) {
    
        
    $buffer = 'pulseWave(';
        $arg0 = \Grafana\Foundation\Testdata\PulseWaveQueryConverter::convert($input->spec->pulseWave);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->queryType !== null && $input->spec->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->spec->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->rawFrameContent !== null && $input->spec->rawFrameContent !== "") {
    
        
    $buffer = 'rawFrameContent(';
        $arg0 =\var_export($input->spec->rawFrameContent, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->refId !== null && $input->spec->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->resultAssertions !== null) {
    
        
    $buffer = 'resultAssertions(';
        $arg0 = \Grafana\Foundation\Testdata\ResultAssertionsConverter::convert($input->spec->resultAssertions);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->scenarioId !== null) {
    
        
    $buffer = 'scenarioId(';
        $arg0 ='\Grafana\Foundation\Testdata\DataqueryScenarioId::fromValue("'.$input->spec->scenarioId.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->seriesCount !== null) {
    
        
    $buffer = 'seriesCount(';
        $arg0 =\var_export($input->spec->seriesCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->sim !== null) {
    
        
    $buffer = 'sim(';
        $arg0 = \Grafana\Foundation\Testdata\SimulationQueryConverter::convert($input->spec->sim);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->spanCount !== null) {
    
        
    $buffer = 'spanCount(';
        $arg0 =\var_export($input->spec->spanCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->spread !== null) {
    
        
    $buffer = 'spread(';
        $arg0 =\var_export($input->spec->spread, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->startValue !== null) {
    
        
    $buffer = 'startValue(';
        $arg0 =\var_export($input->spec->startValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->stream !== null) {
    
        
    $buffer = 'stream(';
        $arg0 = \Grafana\Foundation\Testdata\StreamingQueryConverter::convert($input->spec->stream);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->stringInput !== null && $input->spec->stringInput !== "") {
    
        
    $buffer = 'stringInput(';
        $arg0 =\var_export($input->spec->stringInput, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->timeRange !== null) {
    
        
    $buffer = 'timeRange(';
        $arg0 = \Grafana\Foundation\Testdata\TimeRangeConverter::convert($input->spec->timeRange);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->usa !== null) {
    
        
    $buffer = 'usa(';
        $arg0 = \Grafana\Foundation\Testdata\USAQueryConverter::convert($input->spec->usa);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Testdata\Dataquery && $input->spec->withNil !== null) {
    
        
    $buffer = 'withNil(';
        $arg0 =\var_export($input->spec->withNil, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

