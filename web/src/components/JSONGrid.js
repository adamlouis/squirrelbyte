import React, { useState, useEffect, useRef } from 'react';
import _ from 'lodash';
import styled from 'styled-components';
import Draggable from 'react-draggable';
import { Grid } from 'react-virtualized';
import { Colors } from '../utils/Colors';
import { isValidHttpUrl, copyTextToClipboard } from '../utils/Browser';
import { JSONToReact, safeStringify } from '../utils/JSON';
import { A } from './standard/Link';
import { subscribeKeyDown } from '../utils/KeyPublisher';
import { Modal } from './standard/Modal';
import { Button } from './standard/Button';

const Sizes = {
  DefaultRowHeight: 30,
  DefaultColumnWidth: 256,
  DefaultMinColumnWidth: 30,
  ReRenderSizeThreshold: 5,
};

const BodyContainer = styled.div`
  font-family: 'Courier New', Courier, monospace;
  width: 100%;
  height: 100%;
`;

const DataContainer = styled.div`
  width: 100%;
  height: 100%;
  border: solid black 1px;
  background-color: #f8f8f8;
`;

const Elip = styled.div`
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  width: 100%;
`;

const HeaderLabel = styled(Elip)`
  font-weight: bold;
  padding: 4px 8px;
  font-size: 12px;
  text-align: center;
`;

const JSONCell = styled(Elip)`
  padding: 3px 6px;
`;

const Cell = styled.div`
  font-family: 'Courier New', Courier, monospace;
  background-color: ${(p) => p.backgroundColor || ''};
  border: solid #ddd 1px;
  font-size: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const Bracket = styled.div`
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  background-color: ${Colors.Blue};
  cursor: pointer;

  :hover {
    background-color: ${Colors.DarkerBlue};
  }
`;

const ModalButtonRow = styled.div`
  display: flex;
  justify-content: flex-end;

  > * {
    margin-left: 10px;
  }
`;

const JSONPre = styled.pre`
  background-color: #eee;
  padding: 20px;
  white-space: pre-wrap; /* Since CSS 2.1 */
  white-space: -moz-pre-wrap; /* Mozilla, since 1999 */
  white-space: -pre-wrap; /* Opera 4-6 */
  white-space: -o-pre-wrap; /* Opera 7 */
  word-wrap: break-word; /* Internet Explorer 5.5+ */
`;

const setColumnWidthPreference = (column, width) => {
  try {
    const preferences = getColumnWidthPreferences() || {};
    preferences[column] = width;
    localStorage.setItem(
      'preferences-column-width',
      JSON.stringify(preferences)
    );
  } catch (e) {
    console.warn(e);
  }
};

const getColumnWidthPreferences = () => {
  try {
    return JSON.parse(localStorage.getItem('preferences-column-width'));
  } catch (e) {
    console.warn(e);
  }
};

const later = async (fn) => fn();

export function JSONGrid(props) {
  const [gridWidth, setGridWidth] = useState(0);
  const [gridHeight, setGridHeight] = useState(
    Math.min(20, _.size(props.documents) + 1) * Sizes.DefaultRowHeight
  );

  const bodyRef = useRef(undefined);
  const bodyContainerRef = useRef(undefined);

  useEffect(() => {
    try {
      const el = bodyContainerRef.current;
      const w = el?.getBoundingClientRect().width || 0;
      const h = el?.getBoundingClientRect().height || 0;
      if (Math.abs(gridWidth - w) > Sizes.ReRenderSizeThreshold) {
        setGridWidth(w);
      }
      if (Math.abs(gridHeight - h) > Sizes.ReRenderSizeThreshold) {
        setGridHeight(h);
      }
    } catch (e) {
      console.warn(e);
    }
  }, [gridWidth, gridHeight]);

  const [selectedDoc, setSelectedDoc] = useState(undefined);
  const onClickBracket = (doc) => setSelectedDoc(doc);

  const columns = ['', ...props.paths];
  const onExitModal = () => setSelectedDoc(undefined);

  useEffect(() => {
    const unsubscribeEscape = subscribeKeyDown('Escape', false, () => {
      setSelectedDoc(undefined);
    });
    return unsubscribeEscape;
  }, [setSelectedDoc]);

  const [widthsByColumn, setWidthsByColumn] = useState({ '': 30 });

  const getColumnWidth = ({ index }) => {
    const column = columns[index];

    if (widthsByColumn[column]) {
      return widthsByColumn[column];
    }

    if (index === 0) {
      return Sizes.DefaultRowHeight;
    }

    return Sizes.DefaultColumnWidth;
  };

  const [scrollOffset, setScrollOffset] = useState(0);
  const onScrollBody = (data) => setScrollOffset(data.scrollLeft);

  const onChangeHeaderWidths = (widths) => {
    setWidthsByColumn(widths);
    if (bodyRef.current) {
      bodyRef.current.recomputeGridSize();
    }
  };

  return (
    <DataContainer>
      {selectedDoc && (
        <Modal onExit={onExitModal}>
          <ModalButtonRow>
            <Button
              onClick={() => copyTextToClipboard(safeStringify(selectedDoc))}
            >
              copy
            </Button>
            <Button onClick={onExitModal}>close</Button>
          </ModalButtonRow>
          <JSONPre>{JSONToReact(selectedDoc, 0, 2)}</JSONPre>
        </Modal>
      )}
      <DraggableHeaders
        key={_.join(columns, ',')} // the key is cat of all columns, re-mount if any col changes
        initialWidths={widthsByColumn}
        onChangeWidths={onChangeHeaderWidths}
        paths={columns}
        dragDisabledPaths={['']}
        width={gridWidth}
        scrollOffset={scrollOffset}
      />
      <BodyContainer ref={bodyContainerRef}>
        <Grid
          ref={bodyRef}
          height={gridHeight}
          width={gridWidth}
          rowCount={props.documents.length}
          rowHeight={Sizes.DefaultRowHeight}
          onScroll={onScrollBody}
          columnCount={columns.length}
          columnWidth={getColumnWidth}
          cellRenderer={(cellRenderProps) => {
            const result = (c) => {
              return (
                <Cell
                  key={cellRenderProps.key}
                  parent={cellRenderProps.parent}
                  style={cellRenderProps.style}
                  backgroundColor={
                    cellRenderProps.rowIndex % 2 ? '#f8f8f8' : '#e8e8e8'
                  }
                >
                  {c}
                </Cell>
              );
            };
            const document = props.documents[cellRenderProps.rowIndex];
            if (cellRenderProps.columnIndex === 0) {
              return result(
                <Bracket
                  onClick={() => onClickBracket(document)}
                >{`{}`}</Bracket>
              );
            }

            const path = columns[cellRenderProps.columnIndex];
            const value = _.get(document, path);
            let content;
            if (isValidHttpUrl(value)) {
              content = <A href={value}>{value}</A>;
            } else {
              content = JSON.stringify(value);
            }
            return result(<JSONCell>{content}</JSONCell>);
          }}
        />
      </BodyContainer>
    </DataContainer>
  );
}

// annoyed w/ perf on reat-virtualized scrollsync hoc so diy with react-draggable and normal divs
// i want smooth scrolling when dragging the header ... keep body fixed & render it when done w/ header
// google sheets does something similar - just moves the line & waits for drag to stop before re-rending body
function DraggableHeaders(props) {
  const widthsRef = useRef(
    (() => {
      const widths = {};

      const preferences = _.assign({}, getColumnWidthPreferences(), { '': 30 });
      _.forEach(props.paths, (p) => {
        widths[p] = preferences[p] || Sizes.DefaultColumnWidth;
      });

      return widths;
    })()
  );

  const onChangeWidth = (p, w) => {
    widthsRef.current = _.cloneDeep(widthsRef.current);
    widthsRef.current[p] = w;
    props.onChangeWidths(widthsRef.current);
    later(() => setColumnWidthPreference(p, w));
  };

  useEffect(() => {
    props.onChangeWidths(widthsRef.current);
    // only on mount
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <div
      style={{
        overflow: 'hidden',
        width: props.width,
        display: 'flex',
        alignItems: 'center',
      }}
    >
      <div
        style={{
          display: 'inline-flex',
          transform: `translateX(${props.scrollOffset * -1}px)`,
        }}
      >
        {_.map(props.paths, (p) => {
          return (
            <DraggableHeader
              key={p}
              path={p}
              onChangeWidth={(width) => onChangeWidth(p, width)}
              initialWidth={widthsRef.current[p]}
              minWidth={props.minWidth || Sizes.DefaultMinColumnWidth}
              disableDrag={_.includes(props.dragDisabledPaths, p)}
            />
          );
        })}
      </div>
    </div>
  );
}

function DraggableHeader(props) {
  const widthRef = useRef(props.initialWidth);
  const path = props.path;
  const disableDrag = props.disableDrag;
  const elementRef = useRef(undefined);
  const minWidth = props.minWidth;

  return (
    <div
      key={path}
      ref={elementRef}
      style={{
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        width: `${widthRef.current}px`,
        overflow: 'visible',
        ':hover': {
          backgroundColor: '#eee',
        },
      }}
    >
      <HeaderLabel>{path || <span>&nbsp;</span>}</HeaderLabel>
      {disableDrag ? (
        <div
          style={{
            height: '100%',
            borderLeft: 'solid #bbb 1px',
            marginRight: '3px', // align the re-sizer with the column border
          }}
        />
      ) : (
        <Draggable
          axis="x"
          onStart={() => {
            props.onDragStart && props.onDragStart();
          }}
          onStop={() => {
            props.onDragStop && props.onDragStop();
            const element = elementRef.current;
            if (element) {
              const currentWidth = parseFloat(
                _.replace(element.style.width, 'px', '')
              );
              props.onChangeWidth(currentWidth);
            }
          }}
          onDrag={(event, data) => {
            const element = elementRef.current;
            if (element) {
              const currentWidth = parseFloat(
                _.replace(element.style.width, 'px', '')
              );
              const newWidth = Math.max(currentWidth + data.deltaX, minWidth);
              element.style.width = `${newWidth}px`;
            }
          }}
          position={{ x: 0 }}
        >
          <span
            style={{
              fontFamily: '"Courier New", Courier, monospace',
              cursor: 'col-resize',
              color: Colors.Blue,
              fontWeight: 'bold',
              fontSize: '20px',
              width: '20px',
              marginRight: '-10px', // align the re-sizer with the column border
              zIndex: 1,
              textAlign: 'center',
            }}
          >
            |
          </span>
        </Draggable>
      )}
    </div>
  );
}
