import React, {useContext} from 'react';
import {Button} from "@nextui-org/button";
import {NodeContext} from "../pages";
import {StartNode} from "../wailsjs/wailsjs/go/gui/App";

const Receive = () => {
  const node = useContext(NodeContext)!.node

  return (
    <div className="flex justify-center">
      <Button color={"success"} onClick={() => StartNode(node, "")}>
        Start Node
      </Button>
    </div>
  );
};

export default Receive;
