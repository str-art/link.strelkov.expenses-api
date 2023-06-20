import Toolbar from "@mui/material/Toolbar";
import Bar, { AppBarProps } from "@mui/material/AppBar";
import { forwardRef } from "react";

const AppBar = forwardRef<HTMLElement, AppBarProps>(function AppBar(
  props,
  ref
) {
  return (
    <Bar component={"nav"} ref={ref} {...props}>
      <Toolbar>{"Test"}</Toolbar>
    </Bar>
  );
});

export default AppBar;
