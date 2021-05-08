export const OnLoad = async () => {
  const path = window.location.pathname;

  // if (path.startsWith("/authenticate/callback")) {
  //   const params = getSearchParameters();

  //   window.history.replaceState({}, document.title, "/");

  //   const urlState = params.state;
  //   const urlCode = params.code;

  //   const cookieState = CookieStore.getAuthRedirectState("squirrelbyte");
  //   CookieStore.removeAuthRedirectState("squirrelbyte");

  //   if (!urlState || urlState !== cookieState) {
  //     throw new Error("unexpected state token after oauth redirect");
  //   }

  //   const json = await API.completeAuth(urlCode);
  //   CookieStore.setTokens(json);
  // }

  if (path.startsWith('/oauth/providers/') && path.endsWith('/token')) {
    const provider = path
      .replaceAll('/oauth/providers/', '')
      .replaceAll('/token', '');

    const params = getSearchParameters();

    // TODO: clear params
    // window.history.replaceState({}, document.title, '/');

    const urlCode = params.code;
    // TODO: check state
    // const urlState = params.state;
    //   const cookieState = CookieStore.getAuthRedirectState(appID);
    //   CookieStore.removeAuthRedirectState(appID);

    //   if (!urlState || urlState !== cookieState) {
    //     throw new Error("unexpected state token after oauth redirect");
    //   }

    const res = await fetch(`/api/oauth/providers/${provider}/token`, {
      method: 'POST',
      body: JSON.stringify({ code: urlCode }),
    });

    console.log(await res.json());
    // TODO: handle UI
    // return res.json();
  }
};

const getSearchParameters = () => {
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
