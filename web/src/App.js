import _ from 'lodash';
import React, { useRef, useState, useEffect } from 'react';
import styled from 'styled-components';
import { Loader } from './components/standard/Loader';
import { OnLoad } from './data/OnLoad';

import { Header } from './components/Header';
import { MainDataView } from './components/MainDataView';
import { MainJobView } from './components/MainJobView';
import { MainSchedulerView } from './components/MainSchedulerView';
import { MainKVView } from './components/MainKVView';
import { MainOAuthView } from './components/MainOAuthView';
import { InfoBox } from './components/InfoBox';
import { JSONQueryForm } from './components/JSONQueryForm';
import { runQuery, getQueryFromURL } from './data/QueryController';
import { QueryResultView } from './components/QueryResultView';
import { Colors } from './utils/Colors';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  Redirect,
  useLocation,
} from 'react-router-dom';

const Body = styled.div`
  flex-grow: 1;
  padding: 0px 20px 20px 20px;
  display: flex;
  flex-direction: column;
`;

const TabRow = styled.div`
  display: flex;
`;
const Tab = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 10px;
  cursor: pointer;
  :hover {
    background-color: ${Colors.GrayD};
  }
  background-color: ${(p) => (p.selected ? Colors.GrayD : '')};
`;

OnLoad();

function Nav() {
  const location = useLocation();
  return (
    <TabRow>
      <Link to="/documents">
        <Tab selected={_.startsWith(location.pathname, '/documents')}>
          documents
        </Tab>
      </Link>
      <Link to="/jobs">
        <Tab selected={_.startsWith(location.pathname, '/jobs')}>jobs</Tab>
      </Link>
      <Link to="/schedulers">
        <Tab selected={_.startsWith(location.pathname, '/schedulers')}>
          schedulers
        </Tab>
      </Link>
      <Link to="/kvs">
        <Tab selected={_.startsWith(location.pathname, '/kvs')}>key-values</Tab>
      </Link>
      <Link to="/oauth">
        <Tab selected={_.startsWith(location.pathname, '/oauth')}>oauth</Tab>
      </Link>
    </TabRow>
  );
}

function App() {
  return (
    <div
      style={{
        height: '100%',
        width: '100%',
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      <Header />
      <Router>
        <Body>
          <Nav />
          <Switch>
            <Route path="/documents">
              <MainDataView />
            </Route>
            <Route path="/jobs">
              <MainJobView />
            </Route>
            <Route path="/schedulers">
              <MainSchedulerView />
            </Route>
            <Route path="/kvs">
              <MainKVView />
            </Route>
            <Route path="/oauth">
              <MainOAuthView />
            </Route>
            <Route>
              <Redirect push to="/documents" />
            </Route>
          </Switch>
        </Body>
      </Router>
    </div>
  );
}

export default App;
