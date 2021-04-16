import _ from 'lodash';

const Colors = {
  Green: '#55efc4',
  Yellow: '#fcf9d7',
  Blue: '#74b9ff',
};

const safePretty = (o) => {
  try {
    return JSON.stringify(JSON.parse(o), undefined, 2);
  } catch (e) {}
  return `${o}`;
};

const getAllPaths = (docs) => {
  const result = new Set();
  _.forEach(docs, (d) => getPathsRec(d, '', result));
  return _.sortBy(Array.from(result));
};

const getPathsRec = (obj, currentPath, result) => {
  if (!_.isObject(obj)) {
    return;
  }

  if (_.isArray(obj)) {
    return;
  }

  try {
    const keys = _.keys(obj);
    for (let k of keys) {
      let p = currentPath === '' ? k : currentPath + '.' + k;
      if (_.isObject(obj[k]) && !_.isArray(obj[k])) {
        getPathsRec(obj[k], p, result);
      } else {
        result.add(p);
      }
    }
  } catch (e) {}

  return;
};

const getSpaces = (n, identation = 2) => {
  let s = [];
  for (let i = 0; i < n * identation; i++) {
    s.push(<span key={`${i}`}>&nbsp;</span>);
  }
  return <span>{s}</span>;
};

// todo: use tail rec
// todo: test (fuzz) agains standard JSON.stringify
const JSONToReact = (j, depth, indentation = 2) => {
  if (_.isArray(j)) {
    let children = [];
    for (let i = 0; i < j.length; i++) {
      const v = j[i];
      children.push(getSpaces(depth + 1, indentation));
      children.push(<span>{JSONToReact(v, depth + 1, indentation)}</span>);
      if (i !== j.length - 1) {
        children.push(<span>{','}</span>);
      }
      if (indentation > 0) {
        children.push(<br />);
      }
    }

    let intro = (
      <span>
        {'['}
        {indentation > 0 && j.length > 0 && <br />}
      </span>
    );
    let outro = (
      <span>
        {j.length > 0 && getSpaces(depth, indentation)}
        {']'}
      </span>
    );
    return (
      <span>
        {intro}
        {children}
        {outro}
      </span>
    );
  }

  if (_.isObject(j)) {
    let children = [];
    const keys = _.sortBy(_.keys(j));

    let intro = (
      <span>
        {'{'}
        {indentation > 0 && keys.length > 0 && <br />}
      </span>
    );
    let outro = (
      <span>
        {keys.length > 0 && getSpaces(depth, indentation)}
        {'}'}
      </span>
    );

    for (let i = 0; i < keys.length; i++) {
      const k = keys[i];
      const v = j[k];
      children.push(getSpaces(depth + 1, indentation));
      children.push(<span>{JSON.stringify(k)}</span>);
      children.push(<span>:{indentation > 0 ? ' ' : ''}</span>);
      children.push(<span>{JSONToReact(v, depth + 1, indentation)}</span>);
      if (i !== keys.length - 1) {
        children.push(<span>{','}</span>);
      }
      if (indentation > 0) {
        children.push(<br />);
      }
    }
    return (
      <span>
        {intro}
        {children}
        {outro}
      </span>
    );
  }

  // todo: pass 1) condition fn & 2) render fn as args to allow override rendering of any matching field
  if (_.isString(j) && isValidHttpUrl(j)) {
    return (
      <a target="_blank" rel="noreferrer" href={j}>
        {JSON.stringify(j)}
      </a>
    );
  }

  return <span>{JSON.stringify(j)}</span>;
};

function isValidHttpUrl(s) {
  let url;

  try {
    url = new URL(s);
  } catch (_) {
    return false;
  }

  return url.protocol === 'http:' || url.protocol === 'https:';
}

// https://stackoverflow.com/questions/400212/how-do-i-copy-to-the-clipboard-in-javascript
function fallbackCopyTextToClipboard(text) {
  var textArea = document.createElement('textarea');
  textArea.value = text;

  // Avoid scrolling to bottom
  textArea.style.top = '0';
  textArea.style.left = '0';
  textArea.style.position = 'fixed';

  document.body.appendChild(textArea);
  textArea.focus();
  textArea.select();

  try {
    var successful = document.execCommand('copy');
    var msg = successful ? 'successful' : 'unsuccessful';
    console.log('Fallback: Copying text command was ' + msg);
  } catch (err) {
    console.error('Fallback: Oops, unable to copy', err);
  }

  document.body.removeChild(textArea);
}

function copyTextToClipboard(text) {
  if (!navigator.clipboard) {
    fallbackCopyTextToClipboard(text);
    return;
  }
  navigator.clipboard.writeText(text).then(
    function () {
      console.log('Async: Copying to clipboard was successful!');
    },
    function (err) {
      console.error('Async: Could not copy text: ', err);
    }
  );
}

// https://stackoverflow.com/questions/1090948/change-url-parameters
/**
 * http://stackoverflow.com/a/10997390/11236
 */
function updateURLParameter(url, param, paramVal) {
  var newAdditionalURL = '';
  var tempArray = url.split('?');
  var baseURL = tempArray[0];
  var additionalURL = tempArray[1];
  var temp = '';
  if (additionalURL) {
    tempArray = additionalURL.split('&');
    for (var i = 0; i < tempArray.length; i++) {
      if (tempArray[i].split('=')[0] !== param) {
        newAdditionalURL += temp + tempArray[i];
        temp = '&';
      }
    }
  }

  var encodedParamVal = encodeURIComponent(paramVal);
  var rows_txt = temp + '' + param + '=' + encodedParamVal;
  return baseURL + '?' + newAdditionalURL + rows_txt;
}

function setUrlParameter(k, v) {
  window.history.replaceState(
    '',
    '',
    updateURLParameter(window.location.href, k, v)
  );
}

function clearURLParameters() {
  window.history.replaceState(null, null, window.location.pathname);
}

// https://stackoverflow.com/questions/901115/how-can-i-get-query-string-values-in-javascript
function getUrlParameter(name, url = window.location.href) {
  name = name.replace(/[[\]]/g, '\\$&');
  var regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)'),
    results = regex.exec(url);
  if (!results) return null;
  if (!results[2]) return '';
  return decodeURIComponent(results[2].replace(/\+/g, ' '));
}

const Util = {
  safePretty,
  getAllPaths,
  getSpaces,
  JSONToReact,
  isValidHttpUrl,
  copyTextToClipboard,
  Colors,
  setUrlParameter,
  getUrlParameter,
  clearURLParameters,
};

export default Util;
