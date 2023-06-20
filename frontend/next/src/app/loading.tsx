"use client";
import CircularProgress from "@mui/material/CircularProgress";
import Card from "@mui/material/Card";
import Backdrop from "@mui/material/Backdrop";

export default function Loading() {
  return (
    <Card
      component={"div"}
      sx={{
        width: "100%",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        height: "100vh",
      }}
    >
      <Backdrop open>
        <CircularProgress />
      </Backdrop>
    </Card>
  );
}
