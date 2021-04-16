import React, { useState } from 'react';
import _ from 'lodash';
import styled from 'styled-components';

const ColumnPicker = styled.div`
  border: solid black 1px;
  justify-content: flex-start;
  align-items: flex-start;
  margin: 0px 10px 0px 0px;
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

export function JSONPathSelector(props) {
  const initialSelectedPaths = props.initialSelectedPaths || [];
  const [selectedPaths, setSelectedPaths] = useState(initialSelectedPaths);
  const selectedPathSet = new Set(selectedPaths);

  const onClickColumnCheck = (p) => {
    let newSelectedPaths = _.cloneDeep(selectedPaths);
    if (selectedPathSet.has(p)) {
      newSelectedPaths = _.filter(newSelectedPaths, (path) => path !== p);
    } else {
      newSelectedPaths.push(p);
    }
    setSelectedPaths(newSelectedPaths);
    props.onChange(newSelectedPaths);
  };

  return (
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
  );
}
