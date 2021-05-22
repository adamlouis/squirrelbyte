import _ from 'lodash';
import { safeStringify } from '../utils/JSON';
import {
  createSetJobAction,
  createSetJobsAction,
  createDeleteJobAction,
  dispatch,
} from './store';

export const listJobs = async (params) => {
  try {
    let url = `/api/jobs`;
    // TODO: params
    const resp = await fetch(url);
    const responseBody = await resp.json();
    const jobs = _.get(responseBody, 'jobs');
    dispatch(createSetJobsAction({ jobs: _.keyBy(jobs, 'id') }));
  } catch (e) {
    console.warn(e);
  }
};

export const putJob = async (job) => {
  try {
    const resp = await fetch(`/api/jobs/${job.id}`, {
      method: 'PUT',
      body: safeStringify(job),
    });
    if (resp.status !== 200) {
      throw new Error(await resp.text());
    }
    const responseBody = await resp.json();
    dispatch(createSetJobAction(responseBody));
  } catch (e) {
    console.warn(e);
  }
};

export const deleteJob = async (id) => {
  try {
    const resp = await fetch(`/api/jobs/${id}`, {
      method: 'DELETE',
    });
    if (resp.status !== 200) {
      throw new Error(await resp.text());
    }
    dispatch(createDeleteJobAction({ id }));
  } catch (e) {
    console.warn(e);
  }
};
