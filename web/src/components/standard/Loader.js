import React from 'react';
import { Colors } from '../../utils/Colors';

export function Loader(props) {
  return (
    <div
      style={{
        border: `${props.borderSize} solid #f5f5f5`,
        borderRadius: '50%',
        borderTop: `${props.borderSize} solid ${Colors.Blue}`,
        borderBottom: `${props.borderSize} solid ${Colors.Green}`,
        width: props.size,
        height: props.size,
        animation: 'spin 2s linear infinite',
      }}
    ></div>
  );
}
