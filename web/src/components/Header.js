import React, { useEffect, useRef } from 'react';
import styled from 'styled-components';
import { A } from './standard/Link';
import { Colors } from '../utils/Colors';

const Container = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 20px;
  background-color: ${Colors.GrayD};
`;

const Title = styled.div`
  font-size: 16px;
  font-weight: bold;
`;

const Chip = styled.span`
  background-color: gold;
  font-size: 12px;
  padding: 5px;
  margin-left: 8px;
  font-weight: bold;
`;

const Row = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
`;

export function Header(props) {
  return (
    <Container>
      <Row>
        <Title>squirrel byte</Title>
        <Chip>alpha</Chip>
      </Row>
      <div>
        <A href="https://github.com/adamlouis/squirrelbyte">
          github.com/adamlouis/squirrelbyte
        </A>
      </div>
    </Container>
  );
}
