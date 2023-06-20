import useScrollTrigger from "@mui/material/useScrollTrigger";
import Slide from "@mui/material/Slide";
import { ReactNode } from "react";
import AppBar from "./AppBar";
import Grid from "@mui/material/Grid";

export default function Navigation({ children }: { children: ReactNode }) {
  const trigger = useScrollTrigger();

  return (
    <>
      <Slide appear={false} direction="down" in={!trigger}>
        <AppBar position="sticky" />
      </Slide>
      <Grid container component={"main"} height={"100vh"}>
        {children}
      </Grid>
      <Slide appear={true} direction="up" in={trigger}>
        <AppBar position="fixed" sx={{ top: "auto", bottom: 0 }} />
      </Slide>
    </>
  );
}
