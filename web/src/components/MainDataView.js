import React, { useRef, useState, useEffect } from 'react';
import styled from 'styled-components';
import { Loader } from './standard/Loader';

import { InfoBox } from './InfoBox';
import { JSONQueryForm } from './JSONQueryForm';
import { runQuery, getQueryFromURL } from '../data/QueryController';
import { QueryResultView } from './QueryResultView';
import * as Plot from '@observablehq/plot';
console.log(Plot);
const Body = styled.div`
  padding: 0px 20px;
`;

const ResultView = styled.div`
  min-height: 750px;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
`;

export function MainDataView() {
  const [queryResult, setQueryResult] = useState(undefined);
  const [loading, setLoading] = useState(false);

  const queryFromURLRef = useRef(getQueryFromURL());

  const run = async (query) => {
    setLoading(true);
    setQueryResult(undefined);
    setQueryResult(await runQuery(query));
    setLoading(false);
  };

  useEffect(() => {
    // on load, run query from URL if there is one
    if (queryFromURLRef.current) {
      run(queryFromURLRef.current);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // useEffect(() => {
  //   try {
  //     const el = document.getElementById('plot');
  //     const data = queryResult.response.result;
  //     const body = data.map((d) => d.body);
  //     console.log('QR', data, body);
  //     const plot = Plot.dot(data, {
  //       x: 'body.x',
  //       y: 'body.y',
  //     }).plot();
  //     console.log(el, plot);
  //     el.replaceChildren(plot);
  //   } catch (e) {}
  // }, [queryResult]);

  const onSubmitJSONQueryForm = (j) => run(j);

  return (
    <div>
      <JSONQueryForm
        initialValue={queryFromURLRef.current}
        onSubmit={onSubmitJSONQueryForm}
      />
      <ResultView>
        {loading && <Loader size={'25px'} borderSize={'5px'} />}

        {queryResult && <QueryResultView queryResult={queryResult} />}
      </ResultView>
      {/* <div
        style={{
          width: '640px',
          height: '396px',
          padding: '100px',
          border: 'solid black 3px',
        }}
        id="plot"
      ></div> */}
    </div>
  );
}
