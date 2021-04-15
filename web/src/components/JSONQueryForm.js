import React, { useEffect, useState } from 'react';
import styled from 'styled-components';
import { CodeEditor } from './standard/CodeEditor';
import { InputButton } from './standard/Button';
import { subscribeKeyDown } from '../utils/KeyPublisher';

const Container = styled.form`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const ButtonRow = styled.div`
  display: flex;
  align-items: center;
  justify-content: flex-end;
  margin: 10px 0px;
  width: 100%;
`;

const defaultQuery = JSON.stringify(
  {
    select: [],
    where: {},
    group_by: [],
    order_by: [],
    limit: 1000,
  },
  undefined,
  2
);

export function JSONQueryForm(props) {
  const initialValue = props.initialValue || defaultQuery;
  const [query, setQuery] = useState(initialValue);

  const propsOnSubmit = props.onSubmit;

  const onSubmitForm = (e) => {
    e.preventDefault();
    propsOnSubmit(query);
  };

  useEffect(() => {
    const unsubscribeCmdEnter = subscribeKeyDown('Enter', true, () => {
      propsOnSubmit(query);
    });
    return () => {
      unsubscribeCmdEnter();
    };
  }, [propsOnSubmit, query]);

  return (
    <Container onSubmit={onSubmitForm}>
      <CodeEditor
        initialValue={initialValue}
        height={'400px'}
        width={'100%'}
        onChange={setQuery}
      />
      <ButtonRow>
        <InputButton type="submit" value="submit" />
      </ButtonRow>
    </Container>
  );
}
