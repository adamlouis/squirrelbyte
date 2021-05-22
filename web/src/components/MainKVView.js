import React, { useState, useEffect } from 'react';
import _ from 'lodash';
import { connect } from 'react-redux';
import styled from 'styled-components';

import { BigGrid } from './BigGrid';
import { safeStringify } from '../utils/JSON';
import { JSONPre, JSONPreLine } from './standard/JSON';
import { Button } from './standard/Button';
import { Modal } from './standard/Modal';
import { Colors } from '../utils/Colors';
import { SuccessBanner, ErrorBanner } from './standard/Banner';
import { CodeEditor } from './standard/CodeEditor';
import * as KVController from '../data/KVController';

const Container = styled.div`
  flex-grow: 1;
  border: solid green 2px;
  flex-direction: column;
  display: flex;
`;
const RelativeContainer = styled.div`
  flexgrow: 1;
  position: relative;
  height: 100%;
  width: 100%;
`;
const AbsoluteContainer = styled.div`
  position: absolute;
  height: 100%;
  width: 100%;
`;
const SectionHeader = styled.div`
  padding: 5px 0px;
  font-weight: bold;
`;
const RowRight = styled.div`
  width: 100%;
  display: flex;
  justify-content: flex-end;
  align-items: center;
`;
const Center = styled.div`
  justify-content: center;
  align-items: center;
  display: flex;
`;
const Elip = styled.div`
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  width: 100%;
`;
const JSONCell = styled(Elip)`
  padding: 3px 6px;
`;
const Bracket = styled.div`
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  background-color: ${Colors.Blue};
  cursor: pointer;

  :hover {
    background-color: ${Colors.DarkerBlue};
  }
`;

const mapStateToProps = (state) => {
  return {
    kvs: _.get(state, 'kvs.kvs'),
  };
};

export const MainKVView = connect(mapStateToProps)(function View(props) {
  const [selectedKV, setSelectedKV] = useState(undefined);
  const kvs = _.values(props.kvs);

  useEffect(async () => {
    KVController.listKVs();
  }, []);

  const onClickBracket = (kv) => setSelectedKV(kv);
  const onExitModal = () => setSelectedKV(undefined);

  const [showCreateNew, setShowCreateNew] = useState(false);
  const onClickCreateNew = () => setShowCreateNew(true);

  return (
    <Container>
      <SectionHeader>key value pairs</SectionHeader>
      <RowRight>
        <Button onClick={onClickCreateNew}>create new</Button>
      </RowRight>
      {/* <ListContainer> */}
      <RelativeContainer>
        <AbsoluteContainer>
          <BigGrid
            documents={kvs}
            paths={['key', 'value']}
            headers={['', 'key', 'value']}
            renderers={[
              (p) => {
                return (
                  <Center>
                    <Bracket
                      onClick={() => onClickBracket(kvs[p.index])}
                    >{`{}`}</Bracket>
                  </Center>
                );
              },
              (p) => {
                return <JSONCell>{kvs[p.index].key}</JSONCell>;
              },
              (p) => {
                return (
                  <JSONCell>
                    {kvs[p.index].value instanceof Object
                      ? safeStringify(kvs[p.index].value)
                      : `${kvs[p.index].value}`}
                  </JSONCell>
                );
              },
            ]}
          />
        </AbsoluteContainer>
      </RelativeContainer>
      {selectedKV && (
        <Modal onExit={onExitModal}>
          <KVForm kv={selectedKV} onClose={onExitModal} />
        </Modal>
      )}
      {showCreateNew && (
        <KVForm
          initMode="EDIT"
          kv={{ key: '', value: '' }}
          onClose={() => setShowCreateNew(false)}
        />
      )}
    </Container>
  );
});

function KVForm(props) {
  const [mode, setMode] = useState(props.initMode || 'VIEW');
  const [editorValue, setEditorValue] = useState(
    safeStringify(props.kv, undefined, 2)
  );
  const [editConfigError, setEditConfigError] = useState(undefined);
  const [editConfigSuccess, setEditConfigSuccess] = useState(false);

  const onClickCancel = () => props.onClose();
  const onClickEdit = () => setMode('EDIT');

  const onClickDelete = () => {
    KVController.deleteKV(props.kv.key);
  };

  const onClickSave = async () => {
    setEditConfigSuccess(false);
    setEditConfigError(undefined);
    try {
      const parsed = JSON.parse(editorValue);
      await KVController.putKV(parsed);
      setEditConfigSuccess(true);
    } catch (e) {
      setEditConfigError(`${e}`);
    }
  };

  return (
    <Modal onExit={onClickCancel}>
      {editConfigError && <ErrorBanner>{editConfigError}</ErrorBanner>}
      {editConfigSuccess && <SuccessBanner>success!</SuccessBanner>}
      {mode === 'VIEW' && (
        <JSONPre>{safeStringify(props.kv, undefined, 2)}</JSONPre>
      )}
      {mode === 'EDIT' && (
        <CodeEditor
          initialValue={editorValue}
          height={'100%'}
          width={'100%'}
          onChange={setEditorValue}
        />
      )}
      <div>
        <Button onClick={onClickEdit}>edit</Button>
        <Button onClick={onClickDelete}>delete</Button>
        <Button onClick={onClickSave}>save</Button>
        <Button onClick={onClickCancel}>cancel</Button>
      </div>
    </Modal>
  );
}
