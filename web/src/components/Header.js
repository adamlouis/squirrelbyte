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

const LinkContainer = styled.div`
  display: flex;
  align-items: center;
  justify-conent: center;
  padding: 6px;
  font-size: 12px;
  border-radius: 4px;

  span {
    padding-left: 10px;
  }

  :hover {
    background-color: ${Colors.GrayC};
  }
`;

export function Header(props) {
  return (
    <Container>
      <Row>
        <Title>squirrel byte</Title>
        <Chip>alpha</Chip>
      </Row>
      <A href="https://github.com/adamlouis/squirrelbyte">
        <LinkContainer>
          <GitHubOctocat size={15} />
          <span>github.com/adamlouis/squirrelbyte</span>
        </LinkContainer>
      </A>
    </Container>
  );
}

function GitHubOctocat(props) {
  return (
    <svg
      height={props.size || 32}
      width={props.size || 32}
      viewBox="0 0 16 16"
      version="1.1"
      aria-hidden="true"
    >
      <path
        fill-rule="evenodd"
        d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"
      ></path>
    </svg>
  );
}
