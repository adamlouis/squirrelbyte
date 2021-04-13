import _ from 'lodash';

export const getQueryFromURL = () => {
  try {
    const q = getUrlParameter('q');
    if (!q) {
      clearURLParameters();
      return;
    }

    const j = JSON.parse(q);

    if (!_.isObject(j)) {
      clearURLParameters();
      return;
    }

    return JSON.stringify(j, undefined, 2);
  } catch (e) {
    console.warn(e);
    clearURLParameters();
  }
};

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
