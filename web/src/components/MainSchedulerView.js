import _ from 'lodash';
import React, { useRef, useState, useEffect } from 'react';
import { connect } from 'react-redux';
import { BigGrid } from './BigGrid';
import styled from 'styled-components';
import { Loader } from './standard/Loader';

import { Button } from './standard/Button';
import { InfoBox } from './InfoBox';
import { JSONPre } from './standard/JSON';
import { JSONGrid } from './JSONGrid';
import { runQuery, getQueryFromURL } from '../data/QueryController';
import { QueryResultView } from './QueryResultView';
import { ListView, DynamicHeightList } from './ListView';
import { Modal } from './standard/Modal';
import { Colors } from '../utils/Colors';
import { SuccessBanner, ErrorBanner } from './standard/Banner';
import { CodeEditor } from './standard/CodeEditor';
import { safeStringify } from '../utils/JSON';
import * as SchedulerController from '../data/SchedulerController';

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
    schedulers: _.get(state, 'schedulers.schedulers'),
  };
};

export const MainSchedulerView = connect(mapStateToProps)(function View(props) {
  const [selectedScheduler, setSelectedScheduler] = useState(undefined);
  const schedulers = _.values(props.schedulers);

  useEffect(async () => {
    SchedulerController.listSchedulers();
  }, []);

  const onClickBracket = (scheduler) => setSelectedScheduler(scheduler);
  const onExitModal = () => setSelectedScheduler(undefined);

  const [showQueueNew, setShowQueueNew] = useState(false);
  const onClickQueueNew = () => setShowQueueNew(true);

  return (
    <Container>
      <SectionHeader>key value pairs</SectionHeader>
      <RowRight>
        <Button onClick={onClickQueueNew}>queue new</Button>
      </RowRight>
      {/* <ListContainer> */}
      <RelativeContainer>
        <AbsoluteContainer>
          <BigGrid
            documents={schedulers}
            paths={['key', 'value']}
            headers={['', 'id', 'schedule', 'job_name', 'input']}
            renderers={[
              (p) => {
                return (
                  <Center>
                    <Bracket
                      onClick={() => onClickBracket(schedulers[p.index])}
                    >{`{}`}</Bracket>
                  </Center>
                );
              },
              (p) => {
                return (
                  <JSONCell>{safeStringify(schedulers[p.index].id)}</JSONCell>
                );
              },
              (p) => {
                return (
                  <JSONCell>
                    {safeStringify(schedulers[p.index].schedule)}
                  </JSONCell>
                );
              },
              (p) => {
                return (
                  <JSONCell>
                    {safeStringify(schedulers[p.index].job_name)}
                  </JSONCell>
                );
              },
              (p) => {
                return (
                  <JSONCell>
                    {safeStringify(schedulers[p.index].input)}
                  </JSONCell>
                );
              },
            ]}
          />
        </AbsoluteContainer>
      </RelativeContainer>
      {selectedScheduler && (
        <Modal onExit={onExitModal}>
          <SchedulerForm scheduler={selectedScheduler} onClose={onExitModal} />
        </Modal>
      )}
      {showQueueNew && (
        <SchedulerForm scheduler={{}} onClose={() => setShowQueueNew(false)} />
      )}
    </Container>
  );
});

function SchedulerForm(props) {
  const [mode, setMode] = useState(props.initMode || 'VIEW');
  const [editorValue, setEditorValue] = useState(
    safeStringify(props.scheduler, undefined, 2)
  );
  const [editConfigError, setEditConfigError] = useState(undefined);
  const [editConfigSuccess, setEditConfigSuccess] = useState(false);

  const onClickCancel = () => props.onClose();
  const onClickEdit = () => setMode('EDIT');

  const onClickDelete = () => {
    SchedulerController.deleteScheduler(props.scheduler.id);
  };

  const onClickSave = async () => {
    setEditConfigSuccess(false);
    setEditConfigError(undefined);
    try {
      const parsed = JSON.parse(editorValue);
      await SchedulerController.putScheduler(parsed);
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
        <JSONPre>{safeStringify(props.scheduler, undefined, 2)}</JSONPre>
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
