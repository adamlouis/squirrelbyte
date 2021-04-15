import React, { useRef, useState, useEffect } from 'react';
import styled from 'styled-components';
import { Loader } from './Loader';

import { Header } from './components/Header';
import { InfoBox } from './components/InfoBox';
import { JSONQueryForm } from './components/JSONQueryForm';
import { runQuery, getQueryFromURL } from './data/QueryController';
import { QueryResultView } from './components/QueryResultView';

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

function App() {
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

  const onSubmitJSONQueryForm = (j) => run(j);

  return (
    <div>
      <Header />
      <Body>
        <InfoBox />
        <JSONQueryForm
          initialValue={queryFromURLRef.current}
          onSubmit={onSubmitJSONQueryForm}
        />
        <ResultView>
          {loading && <Loader size={'25px'} borderSize={'5px'} />}
          {queryResult && <QueryResultView queryResult={queryResult} />}
        </ResultView>
      </Body>
    </div>
  );
}

export default App;
