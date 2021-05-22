import { createSetExchangeParamsAction, dispatch } from './store';
import * as OAuthController from './OAuthController';

export const OnLoad = async () => {
  OAuthController.getAllProviders();

  const path = window.location.pathname;
  if (path.startsWith('/oauth/providers/') && path.endsWith('/token')) {
    const provider = path
      .replaceAll('/oauth/providers/', '')
      .replaceAll('/token', '');
    const params = getQueryParameters();
    dispatch(createSetExchangeParamsAction({ params, provider }));
  }
};

const getQueryParameters = () => {
  var params = {};
  var prmstr = window.location.search.substr(1) || '';
  var prmarr = prmstr.split('&');
  for (let i = 0; i < prmarr.length; i++) {
    let tmp = prmarr[i].split('=');
    if (tmp.length === 2) {
      params[tmp[0]] = tmp[1];
    }
  }
  return params;
};
