import _ from 'lodash';
import React, { useState, useEffect } from 'react';
import styled from 'styled-components';
import { safePretty, safeStringify } from '../utils/JSON';

import { JSONGrid } from './JSONGrid';
import { JSONPre, JSONForm } from './standard/JSON';
import { CodeEditor } from './standard/CodeEditor';
import { Button } from './standard/Button';
import { Modal } from './standard/Modal';
import { SuccessBanner, ErrorBanner } from './standard/Banner';
import { connect } from 'react-redux';
import { createSetExchangeParamsAction } from '../data/store';
import * as OAuthController from '../data/OAuthController';

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

const GridContainer = styled.div`
  height: 300px;
`;

const SectionHeader = styled.div`
  padding: 5px 0px;
  font-weight: bold;
`;

const mapStateToProps = (state) => {
  return {
    exchange: _.get(state, 'oauth.exchange'),
    providers: _.get(state, 'oauth.providers'),
  };
};

export const MainOAuthView = connect(mapStateToProps)(function View(props) {
  const [showCreateNew, setShowCreateNew] = useState(false);

  const onExitTokenModal = () => {
    props.dispatch(createSetExchangeParamsAction());
  };

  const OnClickCreateNew = () => setShowCreateNew(true);

  // TODO: step through token exchange
  return (
    <div>
      <SectionHeader>oauth providers</SectionHeader>
      <div
        style={{
          display: 'flex',
          justifyContent: 'flex-end',
          margin: '5px 0px',
        }}
      >
        <Button onClick={OnClickCreateNew}>create new</Button>
      </div>
      <div>
        {_.map(props.providers, (p) => {
          return <OAuthProviderView provider={p} key={p.name} />;
        })}
      </div>
      {!_.isEmpty(props.exchange) && (
        <OAuthTokenExchangeForm
          exchange={props.exchange}
          onClose={onExitTokenModal}
        />
      )}
      {showCreateNew && (
        <OAuthProviderConfigForm
          provider={getNewProviderTemplate()}
          onClose={() => setShowCreateNew(false)}
        />
      )}
    </div>
  );
});

function OAuthProviderView(props) {
  // state
  const [authorizationURL, setAuthorizationURL] = useState(undefined);
  const [authorizationURLError, setAuthorizationURLError] = useState(undefined);

  const [showEditor, setShowEditor] = useState(false);

  // authorize button
  const onClickAuthorize = async () => {
    try {
      const resp = await fetch(
        `/api/oauth/providers/${props.provider.name}/authorize`
      );
      const respBody = await resp.json();
      setAuthorizationURL(_.get(respBody, 'url'));
    } catch (e) {
      setAuthorizationURLError(`${e}`);
    }
  };
  const onClickRedirect = () => {
    window.location = authorizationURL;
  };
  const onExitAuthorizationModal = () => {
    setAuthorizationURL(undefined);
    setAuthorizationURLError(undefined);
  };

  // edit button
  const onClickEdit = () => setShowEditor(true);
  const onClickDelete = async () => {
    OAuthController.deleteConfig(props.provider.name);
  };

  return (
    <div
      style={{
        border: 'solid black 1px',
        marginBottom: '5px',
        padding: '10px',
      }}
    >
      <div>{props.provider.name}</div>
      <div>
        <Button onClick={onClickAuthorize}>authorize</Button>
        <Button onClick={onClickEdit}>edit</Button>
        <Button onClick={onClickDelete}>delete</Button>
      </div>
      {showEditor && (
        <OAuthProviderConfigForm
          provider={props.provider}
          onClose={() => setShowEditor(false)}
        />
      )}
      {authorizationURL && (
        <Modal onExit={onExitAuthorizationModal}>
          <SectionHeader>authorization url</SectionHeader>
          {authorizationURLError && (
            <ErrorBanner>{authorizationURLError}</ErrorBanner>
          )}
          <JSONPre>{safePretty(authorizationURL)}</JSONPre>
          <div>
            <Button onClick={onClickRedirect}>go to url</Button>
            <Button onClick={onExitAuthorizationModal}>cancel</Button>
          </div>
        </Modal>
      )}
    </div>
  );
}

function OAuthProviderConfigForm(props) {
  const [editorValue, setEditorValue] = useState(
    safeStringify(_.get(props, 'provider.config'), undefined, 2)
  );
  const [editConfigError, setEditConfigError] = useState(undefined);
  const [editConfigSuccess, setEditConfigSuccess] = useState(false);

  const onClickCancel = () => props.onClose();

  const onClickSave = async () => {
    setEditConfigSuccess(false);
    setEditConfigError(undefined);
    try {
      const parsed = JSON.parse(editorValue);
      await OAuthController.putConfig(parsed);
      setEditConfigSuccess(true);
    } catch (e) {
      setEditConfigError(`${e}`);
    }
  };

  return (
    <Modal onExit={onClickCancel}>
      <SectionHeader>{props.provider.name} oauth config</SectionHeader>
      {editConfigError && <ErrorBanner>{editConfigError}</ErrorBanner>}
      {editConfigSuccess && <SuccessBanner>success!</SuccessBanner>}
      <CodeEditor
        initialValue={editorValue}
        height={'100%'}
        width={'100%'}
        onChange={setEditorValue}
      />
      <div>
        <Button onClick={onClickSave}>save</Button>
        <Button onClick={onClickCancel}>cancel</Button>
      </div>
    </Modal>
  );
}

async function sha256(message) {
  // encode as UTF-8
  const msgBuffer = new TextEncoder().encode(message);

  // hash the message
  const hashBuffer = await crypto.subtle.digest('SHA-256', msgBuffer);

  // convert ArrayBuffer to Array
  const hashArray = Array.from(new Uint8Array(hashBuffer));

  // convert bytes to hex string
  const hashHex = hashArray
    .map((b) => b.toString(16).padStart(2, '0'))
    .join('');
  return hashHex;
}

function OAuthTokenExchangeForm(props) {
  const [formState, setFormState] = useState('PENDING');
  const [key, setKey] = useState('');
  const [token, setToken] = useState(undefined);
  const [tokenError, setTokenError] = useState(undefined);
  const params = _.get(props, 'exchange.params') || {};
  const provider = _.get(props, 'exchange.provider') || '';

  const onClickComplete = async () => {
    try {
      const res = await fetch(`/api/oauth/providers/${provider}/token`, {
        method: 'POST',
        body: JSON.stringify({ code: _.get(params, 'code') }),
      });
      if (res.status !== 200) {
        const e = new Error(await res.text());
        e.name = '';
        throw e;
      }
      const t = await res.json();
      setToken(t);
      const hash = await sha256(safeStringify(t));
      setKey(`oauth-${provider}-${hash.slice(0, 8)}`);
      setFormState('SUCCESS');
    } catch (e) {
      console.warn(e);
      setFormState('ERROR');
      setTokenError(`${e}`);
    }
  };
  const onClickCopy = () => {};
  const onClickCancel = () => props.onClose();
  const onChangeKey = (e) => {
    setKey(e.target.value);
  };

  const onClickSaveKV = () => {
    fetch(`/api/kvs/${key}`, {
      method: 'PUT',
      body: safeStringify({ key, value: token }),
    });
  };

  return (
    <Modal onExit={onClickCancel}>
      {formState === 'PENDING' && (
        <div>
          <SectionHeader>pending {provider} token exchange</SectionHeader>
          <JSONPre>{safeStringify(params, undefined, 2)}</JSONPre>
          <div>
            <Button onClick={onClickComplete}>complete</Button>
            <Button onClick={onClickCopy}>copy</Button>
            <Button onClick={onClickCancel}>cancel</Button>
          </div>
        </div>
      )}
      {formState === 'SUCCESS' && (
        <div>
          <SectionHeader>{provider} token exchange success</SectionHeader>
          <JSONPre>{safeStringify(token, undefined, 2)}</JSONPre>
          <div>
            <input
              type="text"
              placeholder="key"
              value={key}
              onChange={onChangeKey}
            />
            <Button onClick={onClickSaveKV}>save as key-value</Button>
            <Button onClick={onClickCopy}>copy</Button>
            <Button onClick={onClickCancel}>cancel</Button>
          </div>
        </div>
      )}
      {formState === 'ERROR' && (
        <div>
          <SectionHeader>{provider} token exchange error</SectionHeader>
          <ErrorBanner>
            <JSONPre>{safePretty(tokenError, undefined, 2)}</JSONPre>
          </ErrorBanner>
          <div>
            <Button onClick={onClickCancel}>cancel</Button>
          </div>
        </div>
      )}
    </Modal>
  );
}

const getNewProviderTemplate = () => {
  return {
    name: 'new',
    config: {
      name: '',
      client_id: '',
      client_secret: '',
      auth_url: '',
      token_url: '',
      redirect_url: '',
      scopes: [],
      auth_url_params: {},
    },
  };
};
