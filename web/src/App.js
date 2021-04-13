import React, { useCallback, useState, useEffect } from 'react';
import _ from 'lodash';
import styled from 'styled-components';
import { DocumentList } from './DocumentList';
import Util from './Util';
import { Loader } from './Loader';

import { Header } from './components/Header';
import { InfoHeader } from './components/InfoHeader';
import { JSONQueryForm } from './components/JSONQueryForm';

const TABS = {
  result: { name: 'result' },
  query: { name: 'query' },
};

const Body = styled.div`
  padding: 0px 20px;
`;

const Row = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
`;

const RowLeft = styled(Row)`
  justify-content: flex-start;
`;

const TabRow = styled(RowLeft)`
  border-bottom: solid black 1px;
`;

const Tab = styled.div`
  cursor: pointer;
  padding: 10px;
  border: solid black 1px;
  margin-right: -1px;
  margin-bottom: -1px;
  background-color: ${(p) => (p.selected ? '#ccc' : '#fff')};
  :hover {
    background-color: #ccc;
  }
`;

const Banner = styled.div`
  background-color: ${(p) => p.backgroundColor};
  padding: 3px 8px;
  margin-bottom: 10px;
`;

const QueryView = styled.pre`
  background-color: #ddd;
  padding: 20px;
  border-radius: 3px;
`;

// hand pick for demo
const defaultQuery = {
  select: [],
  where: {},
  group_by: [],
  order_by: [],
  limit: 1000,
};

const getEmptyResult = () => ({
  query: '',
  documents: undefined,
  paths: undefined,
  insights: undefined,
  error: '',
});

const getQueryFromURL = () => {
  try {
    const q = Util.getUrlParameter('q');
    if (!q) {
      Util.clearURLParameters();
      return;
    }

    const j = JSON.parse(q);

    if (!_.isObject(j)) {
      Util.clearURLParameters();
      return;
    }

    return JSON.stringify(j, undefined, 2);
  } catch (e) {
    console.warn(e);
    Util.clearURLParameters();
  }
};

const queryFromUrl = getQueryFromURL();

function App() {
  const [query, setQuery] = useState(
    queryFromUrl || JSON.stringify(defaultQuery, undefined, 2)
  );
  const [result, setResult] = useState(getEmptyResult());
  const [selectedTab, setSelectedTab] = useState('result');
  const [loading, setLoading] = useState(false);

  const onClickTab = (t) => setSelectedTab(t);
  const onSubmitForm = async (e) => {
    e.preventDefault();
    submitForm();
  };

  const submitForm = useCallback(async () => {
    setLoading(true);
    setResult(getEmptyResult());

    const submittedQuery = query;
    try {
      const start = performance.now();
      const res = await fetch('/api/documents:search', {
        method: 'POST',
        body: submittedQuery,
        json: true,
      });
      const elapsed = performance.now() - start;

      const j = await res.json();

      setResult({
        query: submittedQuery,
        documents: j.result,
        paths: Util.getAllPaths(j.result),
        insights: j.insights,
        error: j.message,
        elapsed,
      });

      try {
        Util.setUrlParameter('q', JSON.stringify(JSON.parse(submittedQuery)));
      } catch (e) {
        console.warn(e);
      }
    } catch (e) {
      setResult({
        query: submittedQuery,
        paths: [],
        documents: [],
        insights: {},
        error: `${e}`,
      });
    }
    setLoading(false);
  }, [setLoading, setResult, query]);

  const runDocumentQuery = useCallback(
    async (q) => {
      setLoading(true);
      setResult(getEmptyResult());

      const submittedQuery = q;
      try {
        const start = performance.now();
        const res = await fetch('/api/documents:search', {
          method: 'POST',
          body: submittedQuery,
          json: true,
        });
        const elapsed = performance.now() - start;

        const j = await res.json();

        setResult({
          query: submittedQuery,
          documents: j.result,
          paths: Util.getAllPaths(j.result),
          insights: j.insights,
          error: j.message,
          elapsed,
        });

        try {
          Util.setUrlParameter('q', JSON.stringify(JSON.parse(submittedQuery)));
        } catch (e) {
          console.warn(e);
        }
      } catch (e) {
        setResult({
          query: submittedQuery,
          paths: [],
          documents: [],
          insights: {},
          error: `${e}`,
        });
      }
      setLoading(false);
    },
    [setLoading, setResult]
  );

  useEffect(() => {
    if (queryFromUrl) {
      submitForm();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const onSubmitJSONQueryForm = (j) => {
    runDocumentQuery(j);
  };

  let banner;
  if (result.error) {
    banner = (
      <Banner backgroundColor={Util.Colors.Yellow}>{result.error}</Banner>
    );
  } else {
    if (result.query && result.insights) {
      banner = (
        <Banner backgroundColor={Util.Colors.Green}>{`${_.size(
          _.get(result, 'documents')
        )} records - ${Math.round(
          _.get(result, 'insights.duration_ms')
        )}ms service time - ${Math.round(
          result.elapsed
        )}ms round trip time`}</Banner>
      );
    }
  }

  return (
    <div>
      <Header />
      <Body>
        <InfoHeader />
        <JSONQueryForm onSubmit={onSubmitJSONQueryForm} />
        <div>
          <div style={{ minHeight: '750px' }}>
            {loading && (
              <div
                style={{
                  width: '100%',
                  height: '100%',
                  display: 'flex',
                  justifyContent: 'center',
                }}
              >
                <Loader size={'25px'} borderSize={'5px'} />
              </div>
            )}
            {banner}
            {result.documents && (
              <div>
                <TabRow>
                  {_.map(TABS, (v, k) => (
                    <Tab
                      onClick={() => onClickTab(k)}
                      key={k}
                      selected={k === selectedTab}
                    >
                      {v.name}
                    </Tab>
                  ))}
                </TabRow>
                {selectedTab === 'query' && (
                  <div style={{ display: 'flex' }}>
                    <QueryView>{Util.safePretty(result.query)}</QueryView>
                  </div>
                )}
                {selectedTab === 'result' && (
                  <div
                    style={{
                      margin: '10px 0px',
                    }}
                  >
                    <DocumentList
                      paths={result.paths}
                      documents={result.documents}
                    />
                  </div>
                )}
              </div>
            )}
          </div>
        </div>
      </Body>
    </div>
  );
}

export default App;
