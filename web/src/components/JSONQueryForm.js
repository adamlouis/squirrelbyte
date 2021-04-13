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

export function JSONQueryForm(props) {
  const [query, setQuery] = useState('');

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
        initialValue={''}
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
