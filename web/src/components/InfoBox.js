import React, { useEffect, useRef, useState } from 'react';
import styled from 'styled-components';
import ReactMarkdown from 'react-markdown';

import { Button } from './standard/Button';
import InfoBoxBodyMD from './InfoBoxBody.md';
import InfoBoxHeaderMD from './InfoBoxHeader.md';
import { Colors } from '../utils/Colors';

const Container = styled.div`
  position: relative;
  margin: 10px 0px;
  padding: 5px 20px;
  background-color: ${Colors.GrayD};
  font-size: 14px;

  h1 {
    font-size: 16px;
  }
  h2 {
    font-size: 14px;
  }
  h3 {
    font-size: 12px;
  }
`;

const Content = styled.div`
  display: block;
`;

const IndentButton = styled(Button)`
  margin: 10px 20px 10px 0px;
`;

const Row = styled.div`
  position: absolute;
  right: 0px;
  top: 0px;
  width: 100%;
  display: flex;
  justify-content: flex-end;
`;

const Clicker = styled.div`
  cursor: pointer;
  :hover {
    opacity: 0.5;
  }
`;

export function InfoBox() {
  const [headerMD, setHeaderMD] = useState('');
  const [bodyMD, setBodyMD] = useState('');
  const [showBody, setShowBody] = useState(false);

  // TODO: store info markdown server side & system table after I implement multi db / multi table
  useEffect(() => {
    (async () => {
      try {
        const headerRes = await fetch(InfoBoxHeaderMD);
        const headerTxt = await headerRes.text();
        setHeaderMD(headerTxt);
      } catch (e) {
        setHeaderMD(`erroring loading markdown: ${e}`);
      }
    })();
  }, []);

  useEffect(() => {
    (async () => {
      try {
        const bodyRes = await fetch(InfoBoxBodyMD);
        const bodyTxt = await bodyRes.text();
        setBodyMD(bodyTxt);
      } catch (e) {
        setBodyMD(`erroring loading markdown: ${e}`);
      }
    })();
  }, []);

  const onClickToggle = (e) => {
    e.stopPropagation();
    setShowBody(!showBody);
  };

  return (
    <Container>
      <Clicker onClick={onClickToggle}>
        <Content>
          <Row>
            <IndentButton onClick={onClickToggle}>
              see {showBody ? 'less' : 'more'}
            </IndentButton>
          </Row>
          <ReactMarkdown>{headerMD}</ReactMarkdown>
        </Content>
      </Clicker>
      {showBody && (
        <Content>
          <ReactMarkdown>{bodyMD}</ReactMarkdown>
        </Content>
      )}
    </Container>
  );
}
