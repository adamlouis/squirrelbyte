import _ from 'lodash';
import { safeStringify } from '../utils/JSON';
import {
  createSetSchedulerAction,
  createSetSchedulersAction,
  createDeleteSchedulerAction,
  dispatch,
} from './store';

export const listSchedulers = async (params) => {
  try {
    let url = `/api/schedulers`;
    // TODO: params
    const resp = await fetch(url);
    const responseBody = await resp.json();
    const schedulers = _.get(responseBody, 'schedulers');
    dispatch(
      createSetSchedulersAction({ schedulers: _.keyBy(schedulers, 'id') })
    );
  } catch (e) {
    console.warn(e);
  }
};

export const putScheduler = async (scheduler) => {
  try {
    const resp = await fetch(`/api/schedulers/${scheduler.id}`, {
      method: 'PUT',
      body: safeStringify(scheduler),
    });
    if (resp.status !== 200) {
      throw new Error(await resp.text());
    }
    const responseBody = await resp.json();
    dispatch(createSetSchedulerAction(responseBody));
  } catch (e) {
    console.warn(e);
  }
};

export const deleteScheduler = async (id) => {
  try {
    const resp = await fetch(`/api/schedulers/${id}`, {
      method: 'DELETE',
    });
    if (resp.status !== 200) {
      throw new Error(await resp.text());
    }
    dispatch(createDeleteSchedulerAction({ id }));
  } catch (e) {
    console.warn(e);
  }
};
