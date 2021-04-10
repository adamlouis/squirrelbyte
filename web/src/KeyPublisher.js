import _ from "lodash";

let subscribers = {};

document.addEventListener("keydown", (e) =>
  _.forEach(subscribers[`${e.code}.${e.metaKey}`], (f) => f())
);

export const subscribeKeyDown = (code, metaKey, fn) => {
  const k = `${code}.${metaKey}`;
  if (!subscribers[k]) {
    subscribers[k] = [];
  }

  subscribers[k].push(fn);

  return () => {
    // todo: fn equality is not the right condition to unsub ... fine if callers pass uniq fns ... consider later
    subscribers[k] = _.filter(subscribers[k], (f) => f !== fn);
  };
};
