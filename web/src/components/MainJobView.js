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
import * as JobController from '../data/JobController';

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
    jobs: _.get(state, 'jobs.jobs'),
  };
};

export const MainJobView = connect(mapStateToProps)(function View(props) {
  const [selectedJob, setSelectedJob] = useState(undefined);
  const jobs = _.values(props.jobs);

  useEffect(async () => {
    JobController.listJobs();
  }, []);

  const onClickBracket = (job) => setSelectedJob(job);
  const onExitModal = () => setSelectedJob(undefined);

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
            documents={jobs}
            paths={['key', 'value']}
            headers={[
              '',
              'id',
              'name',
              'status',
              'input',
              'created_at',
              'scheduled_for',
              'succeeded_at',
            ]}
            renderers={[
              (p) => {
                return (
                  <Center>
                    <Bracket
                      onClick={() => onClickBracket(jobs[p.index])}
                    >{`{}`}</Bracket>
                  </Center>
                );
              },
              (p) => {
                return <JSONCell>{safeStringify(jobs[p.index].id)}</JSONCell>;
              },
              (p) => {
                return <JSONCell>{safeStringify(jobs[p.index].name)}</JSONCell>;
              },
              (p) => {
                return (
                  <JSONCell>{safeStringify(jobs[p.index].status)}</JSONCell>
                );
              },
              (p) => {
                return (
                  <JSONCell>{safeStringify(jobs[p.index].input)}</JSONCell>
                );
              },
              (p) => {
                return (
                  <JSONCell>{safeStringify(jobs[p.index].created_at)}</JSONCell>
                );
              },
              (p) => {
                return (
                  <JSONCell>
                    {safeStringify(jobs[p.index].scheduled_for)}
                  </JSONCell>
                );
              },
            ]}
          />
        </AbsoluteContainer>
      </RelativeContainer>
      {selectedJob && (
        <Modal onExit={onExitModal}>
          <JobForm job={selectedJob} onClose={onExitModal} />
        </Modal>
      )}
      {showQueueNew && (
        <JobForm job={{}} onClose={() => setShowQueueNew(false)} />
      )}
    </Container>
  );
});

function JobForm(props) {
  const [mode, setMode] = useState(props.initMode || 'VIEW');
  const [editorValue, setEditorValue] = useState(
    safeStringify(props.job, undefined, 2)
  );
  const [editConfigError, setEditConfigError] = useState(undefined);
  const [editConfigSuccess, setEditConfigSuccess] = useState(false);

  const onClickCancel = () => props.onClose();
  const onClickEdit = () => setMode('EDIT');

  const onClickDelete = () => {
    JobController.deleteJob(props.job.id);
  };

  const onClickSave = async () => {
    setEditConfigSuccess(false);
    setEditConfigError(undefined);
    try {
      const parsed = JSON.parse(editorValue);
      await JobController.putJob(parsed);
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
        <JSONPre>{safeStringify(props.job, undefined, 2)}</JSONPre>
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
