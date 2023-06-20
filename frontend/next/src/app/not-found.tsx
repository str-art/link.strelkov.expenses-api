"use client";

import Backdrop from "@mui/material/Backdrop";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import Typography from "@mui/material/Typography";
import { useRouter } from "next/navigation";

export default function NotFound() {
  const router = useRouter();

  setTimeout(() => {
    router.replace("/");
  }, 2000);

  return (
    <Card>
      <Backdrop open>
        <Box>
          <Typography>
            {"This page doesn't exist. Redirecting to main page."}
          </Typography>
        </Box>
      </Backdrop>
    </Card>
  );
}
