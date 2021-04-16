import React from 'react';

export const A = (props) => (
  <a target="_blank" rel="noreferrer" href={props.href}>
    {props.children}
  </a>
);
