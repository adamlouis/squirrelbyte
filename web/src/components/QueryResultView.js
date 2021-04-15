import React, { useState } from 'react';
import _ from 'lodash';
import styled from 'styled-components';
import { SuccessBanner, ErrorBanner } from './standard/Banner';
import { safePretty } from '../utils/JSON';
import { Colors } from '../utils/Colors';
import { JSONGrid } from './JSONGrid';

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

const TabContent = styled.pre`
  display: flex;
`;
const QueryView = styled.pre`
  background-color: ${Colors.GrayD};
  padding: 20px;
  border-radius: 3px;
  margin: 0px;
`;

export function QueryResultView(props) {
  const [selectedTab, setSelectedTab] = useState('result');
  const onClickTab = (t) => setSelectedTab(t);

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
      <TabContent>
        {selectedTab === 'query' && (
          <QueryView>
            {safePretty(_.get(props, 'queryResult.metadata.query'))}
          </QueryView>
        )}
      </TabContent>
      {selectedTab === 'result' && (
        <JSONGrid documents={_.get(props, 'queryResult.response.result')} />
      )}
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
