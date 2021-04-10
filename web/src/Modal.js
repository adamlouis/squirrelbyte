import React from "react";
import styled from "styled-components";

const Container = styled.div`
  position: fixed;
  z-index: 1;
  top: 0px;
  left: 0px;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
`;

const Overlay = styled(Container)`
  background-color: #000;
  opacity: 0.5;
`;

const Content = styled.div`
  position: absolute;
  display: flex;
  flex-direction: column;
  cursor: default;
  width: 80%;
  height: 80%;
  overflow: auto;
  background-color: #fff;
  border-radius: 3px;
  padding: 10px;
`;

export function Modal(props) {
  const onClickOverlay = () => {
    props.onExit();
  };

  const onClickContent = (e) => {
    e.stopPropagation();
  };

  return (
    <div>
      <Overlay />
      <Container onClick={onClickOverlay}>
        <Content onClick={onClickContent}>{props.children}</Content>
      </Container>
    </div>
  );
}
