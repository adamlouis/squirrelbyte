import React, { useState } from 'react';
import _ from 'lodash';
import styled from 'styled-components';
import { SuccessBanner, ErrorBanner } from './standard/Banner';
import { safePretty, getAllPaths } from '../utils/JSON';
import { Colors } from '../utils/Colors';
import { JSONGrid } from './JSONGrid';
import { JSONPathSelector } from './JSONPathSelector';

const TABS = {
  result: { name: 'result' },
  query: { name: 'query' },
};

const Container = styled.div`
  width: 100%;

  > div {
    margin-bottom: 10px;
  }
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
  background-color: ${(p) => (p.selected ? Colors.GrayC : Colors.White)};
  :hover {
    background-color: ${Colors.GrayC};
  }
`;

const TabContent = styled.div`
  display: ${(p) => (p.show ? 'flex' : 'none')};
`;
const QueryView = styled.pre`
  background-color: ${Colors.GrayD};
  padding: 20px;
  border-radius: 3px;
  margin: 0px;
`;

const getAllPathsMemoized = _.memoize(getAllPaths);

const setPathSelectionPreferences = (paths) => {
  try {
    localStorage.setItem('preferences-selected-paths', JSON.stringify(paths));
  } catch (e) {
    console.warn(e);
  }
};

const getPathSelectionPreferences = () => {
  try {
    return JSON.parse(localStorage.getItem('preferences-selected-paths'));
  } catch (e) {
    console.warn(e);
  }
};

const intersection = (a, b) => {
  const sb = new Set(b);
  return _.filter(a, (e) => sb.has(e));
};

const getDefaultSelectedPaths = (paths) => {
  // select previously selected, if any of the columns present in the new query
  const prefs = intersection(getPathSelectionPreferences() || [], paths);
  if (_.size(prefs) !== 0) {
    return prefs;
  }

  console.log(
    _.uniq(
      _.compact(_.concat(intersection(['id', 'body', 'header'], paths), paths))
    )
  );
  // if none match, select the 1st 4, prefering the default top-level fields of the document resource
  return _.slice(
    _.uniq(
      _.compact(_.concat(intersection(['id', 'body', 'header'], paths), paths))
    ),
    0,
    4
  );
};

const later = async (fn) => fn();

export function QueryResultView(props) {
  const [selectedTab, setSelectedTab] = useState('result');

  const paths = getAllPathsMemoized(
    _.get(props, 'queryResult.response.result')
  );

  const [selectedPaths, setSelectedPaths] = useState(
    getDefaultSelectedPaths(paths)
  );

  const onClickTab = (t) => setSelectedTab(t);

  const onChangeSelectedPaths = (paths) => {
    setSelectedPaths(paths);
    later(() => setPathSelectionPreferences(paths));
  };

  return (
    <Container>
      <Banner queryResult={props.queryResult} />
      <TabRow>
        {_.map(TABS, (v, k) => (
          <Tab
            key={k}
            selected={k === selectedTab}
            onClick={() => onClickTab(k)}
          >
            {v.name}
          </Tab>
        ))}
      </TabRow>
      <TabContent show={selectedTab === 'query'}>
        <QueryView>
          {safePretty(_.get(props, 'queryResult.metadata.query'))}
        </QueryView>
      </TabContent>
      <TabContent show={selectedTab === 'result'}>
        <JSONPathSelector
          paths={paths}
          initialSelectedPaths={selectedPaths}
          onChange={onChangeSelectedPaths}
        />
        <JSONGrid
          documents={_.get(props, 'queryResult.response.result')}
          paths={selectedPaths}
        />
      </TabContent>
    </Container>
  );
}

function Banner(props) {
  const status = _.get(props, 'queryResult.metadata.status');
  const error = _.get(props, 'queryResult.metadata.error');
  const roundTripTime = _.get(props, 'queryResult.metadata.durationMs');

  const result = _.get(props, 'queryResult.response.result');
  const message = _.get(props, 'queryResult.response.message');
  const serviceTime = _.get(props, 'queryResult.response.insights.duration_ms');

  const statusStr = status ? `code=${status}` : '';
  const roundTripTimeStr = status
    ? `round trip time=${_.round(roundTripTime)}ms`
    : '';

  const messageStr = message ? `message="${_.replace(message, '"', '"')}"` : '';
  const errorStr = error ? `error="${_.replace(error, '"', '"')}"` : '';

  const resultCountStr = result ? `rows=${_.size(result)}` : '';
  const serviceTimeStr = result ? `service time=${_.round(serviceTime)}ms` : '';

  const content = _.join(
    _.compact([
      statusStr,
      resultCountStr,
      errorStr,
      messageStr,
      serviceTimeStr,
      roundTripTimeStr,
    ]),
    ' | '
  );

  if (status === 200) {
    return <SuccessBanner>{content}</SuccessBanner>;
  }

  return <ErrorBanner>{content}</ErrorBanner>;
}
