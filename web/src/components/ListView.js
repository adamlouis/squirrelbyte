import { PureComponent, useEffect } from 'react';
import {
  List,
  AutoSizer,
  CellMeasurerCache,
  CellMeasurer,
} from 'react-virtualized';
import * as All from 'react-virtualized';
console.log(All);

// import { List as ImmutableList } from 'immutable';
// import PropTypes from 'prop-types';
// import * as React from 'react';
// import {
//   ContentBox,
//   ContentBoxHeader,
//   ContentBoxParagraph,
// } from '../demo/ContentBox';
// import AutoSizer from './AutoSizer';
// import List, { type RowRendererParams } from '../List';
// import styles from './AutoSizer.example.css';

// type State = {
//   hideDescription: boolean,
// };

export function ListView(props) {
  return (
    <AutoSizer>
      {({ width, height }) => (
        <List
          // className={styles.List}
          height={height}
          width={width}
          rowCount={props.items.length}
          rowHeight={props.rowHeight || 30}
          rowRenderer={props.rowRenderer}
        />
      )}
    </AutoSizer>
  );
}

export class DynamicHeightList extends PureComponent {
  constructor(props) {
    super(props);

    this._cache = new CellMeasurerCache({
      fixedWidth: true,
    });

    this._rowRenderer = this._rowRenderer.bind(this);
  }

  render() {
    // const { width } = this.props;
    // console.log(this.props);
    return (
      <AutoSizer>
        {({ width, height }) => (
          <List
            // className={styles.BodyGrid}
            deferredMeasurementCache={this._cache}
            height={height}
            width={width}
            overscanRowCount={0}
            rowCount={this.props.items.length}
            rowHeight={this._cache.rowHeight}
            rowRenderer={this._rowRenderer}
          />
        )}
      </AutoSizer>
    );
  }

  _rowRenderer({ index, key, parent, style }) {
    console.log('_rowRenderer');
    const parentRowRenderer = this.props.rowRenderer;
    return (
      <CellMeasurer
        cache={this._cache}
        columnIndex={0}
        key={key}
        rowIndex={index}
        parent={parent}
      >
        {({ measure, registerChild }) => {
          // setTimeout(() => {
          //   measure();
          //   console.log('measure', measure);
          // }, 1000);
          return (
            <div ref={registerChild} style={style}>
              {parentRowRenderer({ index })}
            </div>
            // <RowWrapper
            //   measure={measure}
            //   registerChild={registerChild}
            //   parentRowRenderer={parentRowRenderer}
            //   style={style}
            //   key={key}
            //   k={key}
            //   index={index}
            // />
          );
        }}
      </CellMeasurer>
    );
  }
}

function RowWrapper(props) {
  const { measure, parentRowRenderer, registerChild, style, k, index } = props;
  // <div ref={registerChild} style={style} key={key}>
  //   {parentRowRenderer({ index })}
  // </div>
  useEffect(() => {
    setTimeout(() => {
      measure();
      console.log('measure', measure);
    }, 1000);
  }, []);
  return (
    <div ref={registerChild} style={style}>
      {parentRowRenderer({ index })}
    </div>
  );
}
