import _ from 'lodash';

import {
  getURLQueryParameter,
  setURLQueryParameter,
  clearURLQueryParameters,
} from '../utils/Browser';

const QUERY_PARAM_NAME = 'q';

export const runQuery = async (query) => {
  const queryResult = {
    metadata: {
      query,
      durationMs: undefined,
      status: undefined,
      error: undefined,
    },
    response: undefined,
  };

  try {
    const start = performance.now();
    const res = await fetch('/api/documents:query', {
      method: 'POST',
      body: queryResult.metadata.query,
      json: true,
    });
    queryResult.metadata.durationMs = performance.now() - start;
    queryResult.metadata.status = res.status;

    const resTxt = res.clone();
    try {
      queryResult.response = await res.json();
    } catch (e) {
      // if parsing response as json fails, parse as text
      throw new Error(await resTxt.text());
    }

    // side effect: set URL query param
    try {
      setURLQueryParameter(
        QUERY_PARAM_NAME,
        JSON.stringify(JSON.parse(queryResult.metadata.query))
      );
    } catch (e) {
      console.warn(e);
    }
  } catch (e) {
    queryResult.metadata.error = `${e}`;
  }

  return queryResult;
};

export const getQueryFromURL = () => {
  try {
    const q = getURLQueryParameter(QUERY_PARAM_NAME);
    if (!q) {
      clearURLQueryParameters();
      return;
    }

    const j = JSON.parse(q);

    if (!_.isObject(j)) {
      clearURLQueryParameters();
      return;
    }

    return JSON.stringify(j, undefined, 2);
  } catch (e) {
    console.warn(e);
    clearURLQueryParameters();
  }
};
