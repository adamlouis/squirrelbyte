import _ from 'lodash';
import { safeStringify } from '../utils/JSON';
import {
  createSetKVAction,
  createSetKVsAction,
  createDeleteKVAction,
  dispatch,
} from './store';

export const listKVs = async (params) => {
  try {
    let url = `/api/kvs`;
    // TODO: params
    const resp = await fetch(url);
    const responseBody = await resp.json();
    const kvs = _.get(responseBody, 'kvs');
    dispatch(createSetKVsAction({ kvs: _.keyBy(kvs, 'key') }));
  } catch (e) {
    console.warn(e);
  }
};

export const putKV = async (kv) => {
  try {
    const resp = await fetch(`/api/kvs/${kv.key}`, {
      method: 'PUT',
      body: safeStringify(kv),
    });
    if (resp.status !== 200) {
      throw new Error(await resp.text());
    }
    const responseBody = await resp.json();
    dispatch(createSetKVAction(responseBody));
  } catch (e) {
    console.warn(e);
  }
};

export const deleteKV = async (key) => {
  try {
    const resp = await fetch(`/api/kvs/${key}`, {
      method: 'DELETE',
    });
    if (resp.status !== 200) {
      throw new Error(await resp.text());
    }
    dispatch(createDeleteKVAction({ key }));
  } catch (e) {
    console.warn(e);
  }
};
