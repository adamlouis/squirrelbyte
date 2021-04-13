import React, { useEffect, useState, useRef } from 'react';
import _ from 'lodash';
import styled from 'styled-components';
import { Grid } from 'react-virtualized';
import Draggable from 'react-draggable';
import { Modal } from './Modal';
import { subscribeKeyDown } from './KeyPublisher';
import Util from './Util';

const Sizes = {
  DefaultRowHeight: 30,
  DefaultColumnWidth: 120,
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

const JSONCell = styled(Elip)`
  padding: 3px 6px;
`;

const ColLabel = styled(Elip)`
  font-weight: bold;
  padding: 0px 8px;
  font-size: 12px;
  text-align: center;
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
  background-color: ${Util.Colors.Blue};
  cursor: pointer;

  :hover {
    opacity: 0.8;
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

const ColumnPicker = styled.div`
  border: solid black 1px;
  justify-content: flex-start;
  align-items: flex-start;
  margin: 0px 10px 10px 0px;
  background-color: #fff;
`;

const ColumnListContainer = styled.div`
  max-width: 200px;
  max-height: 500px;
  padding: 5px;
  overflow: auto;
`;
const ColumnList = styled.div`
  width: 200px;
  font-size: 14px;
`;

const ColumnListHeader = styled.div`
  margin: 4px 0px;
  font-size: 14px;
  text-align: center;
  font-weight: bold;
`;
const ColumnListRow = styled.div`
  display: flex;
  cursor: pointer;
  :hover {
    background-color: #eee;
  }
`;
const ColumnListLabel = styled.span`
  font-family: 'Courier New', Courier, monospace;
`;
const ColumnListCheck = styled.input`
  cursor: pointer;
`;

const ModalButtonRow = styled.div`
  display: flex;
  justify-content: flex-end;
`;

const ModalButton = styled.button`
  padding: 5px 10px;
  cursor: pointer;
  margin-left: 5px;
  background-color: #ddd;
`;

export function DocumentList(props) {
  const [selectedDoc, setSelectedDoc] = useState(undefined);
  const [selectedPaths, setSelectedPaths] = useState(
    _.slice(
      _.uniq(
        _.filter(
          [
            // hand pick for demo
            'body.title',
            'body.url',
            'header.hn_url',
            'body.score',
            'body.by',
            'body',
            'header',
            'id',
            ...props.paths,
          ],
          (p) => _.includes(props.paths, p)
        )
      ),
      0,
      4
    )
  );
  const columns = ['', ...selectedPaths];
  const selectedPathSet = new Set(selectedPaths);

  const [gridWidth, setGridWidth] = useState(0);
  const [gridHeight, setGridHeight] = useState(
    Math.min(20, props.documents.length + 1) * Sizes.DefaultRowHeight
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

  const onClickColumnCheck = (p) => {
    let newSelectedPaths;
    if (selectedPathSet.has(p)) {
      newSelectedPaths = _.filter(selectedPaths, (c) => c !== p);
    } else {
      newSelectedPaths = _.cloneDeep(selectedPaths);
      newSelectedPaths.push(p);
    }
    setSelectedPaths(newSelectedPaths);
  };

  const onClickBracket = (doc) => setSelectedDoc(doc);
  const onExitModal = () => setSelectedDoc(undefined);

  useEffect(() => {
    const unsubscribeEscape = subscribeKeyDown('Escape', false, () => {
      setSelectedDoc(undefined);
    });
    return unsubscribeEscape;
  }, [setSelectedDoc]);

  const [widthsByColumn, setWidthsByColumn] = useState({});

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

  // todo: preserve column sizes on recompute
  const onHeadersInit = (widths) => {
    setWidthsByColumn(widths);
    if (bodyRef.current) {
      bodyRef.current.recomputeGridSize();
    }
  };

  const onDragHeaderStop = (idx, widths) => {
    setWidthsByColumn(widths);
    if (bodyRef.current) {
      bodyRef.current.recomputeGridSize(idx);
    }
  };

  const [scrollLeft, setScrollLeft] = useState(0);
  const onScrollBody = (data) => setScrollLeft(data.scrollLeft);

  return (
    <div style={{ display: 'flex' }}>
      {selectedDoc && (
        <Modal onExit={onExitModal}>
          <ModalButtonRow>
            <ModalButton
              onClick={() =>
                Util.copyTextToClipboard(JSON.stringify(selectedDoc))
              }
            >
              copy
            </ModalButton>
            <ModalButton onClick={onExitModal}>close</ModalButton>
          </ModalButtonRow>
          <JSONPre>{Util.JSONToReact(selectedDoc, 0, 2)}</JSONPre>
        </Modal>
      )}
      <div>
        <ColumnPicker>
          <ColumnListHeader>paths</ColumnListHeader>
          <ColumnListContainer>
            <ColumnList>
              {_.map(props.paths, (p) => (
                <ColumnListRow key={p} onClick={() => onClickColumnCheck(p)}>
                  <ColumnListCheck
                    type="checkbox"
                    value={p}
                    onChange={() => onClickColumnCheck(p)}
                    checked={selectedPathSet.has(p)}
                  />
                  <ColumnListLabel>{p}</ColumnListLabel>
                </ColumnListRow>
              ))}
            </ColumnList>
          </ColumnListContainer>
        </ColumnPicker>
      </div>
      <DataContainer>
        <DraggableHeaders
          key={_.join(columns, '')} // the key is cat of all columns, re-mount if any col changes
          onInit={onHeadersInit}
          onDragStart={() => {}}
          onDragStop={onDragHeaderStop}
          items={columns}
          initialWidths={[30]}
          disabled={[0]}
          scrollLeft={scrollLeft}
          width={gridWidth}
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
              const path = selectedPaths[cellRenderProps.columnIndex - 1];
              const value = _.get(document, path);
              let content;
              if (Util.isValidHttpUrl(value)) {
                content = (
                  <a target="_blank" rel="noreferrer" href={value}>
                    {value}
                  </a>
                );
              } else {
                content = JSON.stringify(value);
              }
              return result(<JSONCell>{content}</JSONCell>);
            }}
          />
        </BodyContainer>
      </DataContainer>
    </div>
  );
}

// annoyed w/ perf on scrollsync hoc so diy
// i want smooth scrolling when dragging the header ... keep body fixed & render it when done w/ header
// google sheets just moves the line & waits for drag to stop before computing body
class DraggableHeaders extends React.Component {
  constructor(props) {
    super(props);
    this.items = _.cloneDeep(props.items);
    this.elements = new Array(this.items.length);
    this.widths = new Array(this.items.length);
    this.minWidth = props.minWidth || 30;
    for (let i = 0; i < this.widths.length; i++) {
      this.widths[i] = (props.initialWidths || [])[i] || 240;
    }
    this.disabled = new Set(props.disabled);
    this.containerRef = React.createRef();
  }

  componentDidMount() {
    this.props.onInit && this.props.onInit(this.getWidthsByColumn(this.widths));
  }

  shouldComponentUpdate() {
    return !this.isDragging;
  }

  getWidthsByColumn() {
    const result = {};
    for (let i = 0; i < this.items.length; i++) {
      result[this.items[i]] = this.widths[i];
    }
    return result;
  }

  refCollector = (idx) => {
    return (e) => {
      if (e) {
        this.elements[idx] = e;
      }
    };
  };

  render() {
    return (
      <div
        style={{
          overflow: 'hidden',
          width: this.props.width,
          display: 'flex',
          alignItems: 'center',
        }}
      >
        <div
          ref={this.containerRef}
          style={{
            padding: '4px',
            display: 'inline-flex',
            transform: `translateX(${this.props.scrollLeft * -1}px)`,
          }}
        >
          {_.map(this.items, (i, idx) => {
            return (
              <div
                key={i}
                ref={this.refCollector(idx)}
                style={{
                  display: 'flex',
                  justifyContent: 'space-between',
                  alignItems: 'center',
                  width: `${this.widths[idx]}px`,
                  overflow: 'visible',
                  ':hover': {
                    backgroundColor: '#eee',
                  },
                }}
              >
                <div />
                <ColLabel>{i}</ColLabel>
                {this.disabled.has(idx) ? (
                  <div
                    style={{
                      height: '100%',
                      borderLeft: 'solid #bbb 1px',
                      marginRight: '3px', // seems about right
                    }}
                  />
                ) : (
                  <Draggable
                    axis="x"
                    onStart={() => {
                      this.props.onDragStart &&
                        this.props.onDragStart(
                          idx,
                          this.getWidthsByColumn(this.widths)
                        );
                    }}
                    onStop={() => {
                      this.props.onDragStop &&
                        this.props.onDragStop(
                          idx,
                          this.getWidthsByColumn(this.widths)
                        );
                    }}
                    onDrag={(event, data) => {
                      const el = this.elements[idx];
                      if (this.containerRef.current) {
                        // reset scroll position on drag so left side stays fixed & right side grows
                        this.containerRef.current.style.transform = `translateX(${
                          this.props.scrollLeft * -1
                        }px)`;
                      }
                      if (el) {
                        const sz = parseFloat(
                          _.replace(el.style.width, 'px', '')
                        );
                        const newSz = Math.max(sz + data.deltaX, this.minWidth);
                        el.style.width = `${newSz}px`;
                        this.widths[idx] = newSz;
                      }
                    }}
                    position={{ x: 0 }}
                  >
                    <span
                      style={{
                        fontFamily: '"Courier New", Courier, monospace',
                        cursor: 'col-resize',
                        color: Util.Colors.Blue,
                        fontWeight: 'bold',
                        fontSize: '20px',
                        width: '20px',
                        marginRight: '-5px', // seems about right
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
          })}
        </div>
      </div>
    );
  }
}
