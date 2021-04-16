import React, { useEffect, useRef } from 'react';
import styled from 'styled-components';

const Container = styled.div`
  position: relative;
  height: ${(p) => p.height};
  width: ${(p) => p.width};
  border: 'solid black 1px';
`;

export function CodeEditor(props) {
  const editorRef = useRef();

  useEffect(() => {
    if (!editorRef.current) {
      return;
    }

    if (!window.ace) {
      return;
    }

    const editor = window.ace.edit(editorRef.current);
    editor.session.setMode('ace/mode/json');
    editor.setTheme('ace/theme/eclipse');
    editor.setValue(props.initialValue || '');
    editor.setShowPrintMargin(false);
    editor.clearSelection();
    editor.on('change', (e) => {
      props.onChange(editor.getValue());
    });
    // only on mount
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <Container height={props.height} width={props.width}>
      <div
        ref={editorRef}
        style={{
          height: `100%`,
          width: `100%`,
        }}
      ></div>
    </Container>
  );
}
