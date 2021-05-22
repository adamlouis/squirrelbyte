import _ from 'lodash';
import { createReducer, createAction, configureStore } from '@reduxjs/toolkit';
import { createLogger } from 'redux-logger';

export const createSetExchangeParamsAction = createAction('oauth/SET_EXCHANGE');
export const createSetProviderAction = createAction('oauth/SET_PROVIDER');
export const createSetProvidersAction = createAction('oauth/SET_PROVIDERS');
export const createDeleteProviderAction = createAction('oauth/DELETE_PROVIDER');

export const createSetKVAction = createAction('kvs/SET_KV');
export const createSetKVsAction = createAction('kvs/SET_KVS');
export const createDeleteKVAction = createAction('kvs/DELETE_KV');

export const createSetJobAction = createAction('jobs/SET_JOB');
export const createSetJobsAction = createAction('jobs/SET_JOBS');
export const createDeleteJobAction = createAction('jobs/DELETE_JOB');

export const createSetSchedulerAction = createAction(
  'schedulers/SET_SCHEDULER'
);
export const createSetSchedulersAction = createAction(
  'schedulers/SET_SCHEDULERS'
);
export const createDeleteSchedulerAction = createAction(
  'schedulers/DELETE_SCHEDULER'
);

const initOauthState = {
  exchange: undefined,
  providers: [],
};

const initKVsState = {
  kvs: {},
};

const initJobsState = {
  jobs: {},
};

const initSchedulersState = {
  schedulers: {},
};

const oauthReducer = createReducer(initOauthState, (builder) => {
  builder
    .addCase(createSetExchangeParamsAction, (state, action) => {
      state.exchange = _.get(action, 'payload');
    })
    .addCase(createSetProviderAction, (state, action) => {
      state.providers[_.get(action, 'payload.name')] = _.get(action, 'payload');
    })
    .addCase(createSetProvidersAction, (state, action) => {
      state.providers = _.get(action, 'payload.providers');
    })
    .addCase(createDeleteProviderAction, (state, action) => {
      delete state.providers[_.get(action, 'payload.name')];
    });
});

const kvsReducer = createReducer(initKVsState, (builder) => {
  builder
    .addCase(createSetKVAction, (state, action) => {
      state.kvs[_.get(action, 'payload.key')] = _.get(action, 'payload');
    })
    .addCase(createSetKVsAction, (state, action) => {
      state.kvs = _.get(action, 'payload.kvs');
    })
    .addCase(createDeleteKVAction, (state, action) => {
      delete state.kvs[_.get(action, 'payload.key')];
    });
});
const jobsReducer = createReducer(initJobsState, (builder) => {
  builder
    .addCase(createSetJobAction, (state, action) => {
      state.jobs[_.get(action, 'payload.id')] = _.get(action, 'payload');
    })
    .addCase(createSetJobsAction, (state, action) => {
      state.jobs = _.get(action, 'payload.jobs');
    })
    .addCase(createDeleteJobAction, (state, action) => {
      delete state.jobs[_.get(action, 'payload.id')];
    });
});
const schedulersReducer = createReducer(initSchedulersState, (builder) => {
  builder
    .addCase(createSetSchedulerAction, (state, action) => {
      state.schedulers[_.get(action, 'payload.id')] = _.get(action, 'payload');
    })
    .addCase(createSetSchedulersAction, (state, action) => {
      state.schedulers = _.get(action, 'payload.schedulers');
    })
    .addCase(createDeleteSchedulerAction, (state, action) => {
      delete state.schedulers[_.get(action, 'payload.id')];
    });
});

export const store = configureStore({
  reducer: {
    oauth: oauthReducer,
    kvs: kvsReducer,
    jobs: jobsReducer,
    schedulers: schedulersReducer,
  },
  middleware: [createLogger({})],
});

export const dispatch = store.dispatch;
