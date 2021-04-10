import React, { useEffect } from "react";
import Util from "./Util";
import { subscribeKeyDown } from "./KeyPublisher";

export function Ace(props) {
  useEffect(() => {
    var editor = window.ace.edit("editor");
    editor.session.setMode("ace/mode/json");
    editor.setTheme("ace/theme/eclipse");

    editor.setValue(props.initialValue || "");
    editor.clearSelection();

    editor.on("change", (e) => {
      props.onChange(editor.getValue());
    });

    const unsubscribeCmdB = subscribeKeyDown("KeyB", true, () => {
      editor.setValue(Util.safePretty(editor.getValue()));
      editor.clearSelection();
    });

    return () => {
      editor.destroy();
      unsubscribeCmdB();
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <div
      style={{
        position: "relative",
        height: `${props.height}`,
        width: `${props.width}`,
        border: "solid black 1px",
      }}
    >
      <div
        id="editor"
        style={{
          height: `100%`,
          width: `100%`,
        }}
      ></div>
    </div>
  );
}
