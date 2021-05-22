import _ from 'lodash';
import { safeStringify } from '../utils/JSON';
import {
  createSetProviderAction,
  createSetProvidersAction,
  createDeleteProviderAction,
  store,
} from './store';

export const getAllProviders = async () => {
  let providers = [];
  let pageToken = '';
  try {
    while (true) {
      let url = '/api/oauth/providers';
      if (pageToken) {
        url += `page_token=${pageToken}`;
      }

      const resp = await fetch(url);
      const responseBody = await resp.json();
      providers = _.concat(providers, responseBody.providers);
      if (!responseBody.next_page_token) {
        break;
      }
      pageToken = responseBody.next_page_token;
    }
  } catch (e) {
    console.warn(e);
  }

  store.dispatch(
    createSetProvidersAction({ providers: _.keyBy(providers, 'name') })
  );
};

export const putConfig = async (config) => {
  const resp = await fetch(`/api/oauth/providers/${config.name}/config`, {
    method: 'PUT',
    body: safeStringify(config),
  });
  if (resp.status !== 200) {
    throw new Error(await resp.text());
  }
  const responseBody = await resp.json();
  store.dispatch(
    createSetProviderAction({ name: responseBody.name, config: responseBody })
  );
};

export const deleteConfig = async (name) => {
  try {
    await fetch(`/api/oauth/providers/${name}/config`, {
      method: 'DELETE',
    });
    store.dispatch(createDeleteProviderAction({ name }));
  } catch (e) {
    console.warn(e);
  }
};
