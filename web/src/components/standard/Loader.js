import React from 'react';
import Util from './Util';

export function Loader(props) {
  return (
    <div
      style={{
        border: `${props.borderSize} solid #f5f5f5`,
        borderRadius: '50%',
        borderTop: `${props.borderSize} solid ${Util.Colors.Blue}`,
        borderBottom: `${props.borderSize} solid ${Util.Colors.Green}`,
        width: props.size,
        height: props.size,
        animation: 'spin 2s linear infinite',
      }}
    ></div>
  );
}
