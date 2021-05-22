import { initial } from 'lodash';
import { useRef, useState } from 'react';
import styled from 'styled-components';
import { Colors } from '../../utils/Colors';
import { Button } from './Button';
import { ErrorBanner } from './Banner';
import { CodeEditor } from './CodeEditor';

export const JSONPre = styled.pre`
  background-color: #eee;
  padding: 20px;
  margin: 0px;
  white-space: pre-wrap; /* Since CSS 2.1 */
  white-space: -moz-pre-wrap; /* Mozilla, since 1999 */
  white-space: -pre-wrap; /* Opera 4-6 */
  white-space: -o-pre-wrap; /* Opera 7 */
  word-wrap: break-word; /* Internet Explorer 5.5+ */
`;

export const JSONPreLine = styled.pre`
  background-color: #eee;
  padding: 20px;
  margin: 0px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
`;

export function JSONForm(props) {
  const preRef = useRef(undefined);
  const initialValueRef = useRef(props.value);
  const [editable, setEditable] = useState(false);
  const [error, setError] = useState(undefined);

  const onClickEdit = () => {
    setEditable(true);
  };
  const onSubmit = async (e) => {
    setEditable(false);
    try {
      // let v = props.submit(preRef.current.innerText);
      // if (v instanceof Promise) {
      //   v = await v;
      // }
      // console.log(v);
      setError(`success!`);
    } catch (e) {
      setError(`${e}`);
      setEditable(true);
      return;
    }

    preRef.current.innerText = initialValueRef.current;
  };
  const onCancel = () => {
    setEditable(false);
    // preRef.current.innerText = initialValueRef.current;
  };

  return (
    <div style={{ border: 'solid black 1px' }}>
      <Button onClick={onClickEdit}>edit</Button>
      <Button onClick={onSubmit}>save</Button>
      <Button onClick={onCancel}>cancel</Button>
      {error && (
        <div style={{ padding: '10px' }}>
          <ErrorBanner>{error}</ErrorBanner>
        </div>
      )}
      <div
        style={{
          border: editable ? 'solid blue 1px' : 'solid 1px',
          borderRadius: '3px',
          height: '300px',
        }}
      >
        {editable && (
          <CodeEditor
            initialValue={props.value}
            height={'100%'}
            width={'100%'}
            onChange={() => {}}
          />
        )}

        {!editable && (
          <JSONPre
            ref={preRef}
            contentEditable={editable}
            suppressContentEditableWarning
          >
            {props.value}
          </JSONPre>
        )}
      </div>
    </div>
  );
}
