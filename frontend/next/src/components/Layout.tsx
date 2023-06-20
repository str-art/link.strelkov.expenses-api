"use client";
import { ReactNode } from "react";
import Grid from "@mui/material/Grid";
import Container from "@mui/material/Container";
import Navigation from "./Navigation";
import Box from "@mui/material/Box";

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <Navigation>
      <Container
        component={"section"}
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        {children}
      </Container>
    </Navigation>
  );
}
