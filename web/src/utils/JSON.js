import _ from 'lodash';
import { isValidHttpUrl } from './Browser';

export const safePretty = (o) => {
  try {
    return JSON.stringify(JSON.parse(o), undefined, 2);
  } catch (e) {}
  return `${o}`;
};

export const safeStringify = (...args) => {
  try {
    return JSON.stringify(...args);
  } catch (e) {
    return `${e}`;
  }
};

export const getAllPaths = (docs) => {
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
      result.add(p);
      if (_.isObject(obj[k]) && !_.isArray(obj[k])) {
        getPathsRec(obj[k], p, result);
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
export const JSONToReact = (j, depth, indentation = 2) => {
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
